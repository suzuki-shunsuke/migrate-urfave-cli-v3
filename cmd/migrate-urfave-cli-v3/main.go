package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type Replacer struct {
	pattern     string
	regexp      *regexp.Regexp
	replacement []byte
}

func (r *Replacer) Init() error {
	re, err := regexp.Compile(r.pattern)
	if err != nil {
		return fmt.Errorf("compile a replacer pattern (regular expression): %w", err)
	}
	r.regexp = re
	return nil
}

func newReplacer(pattern, replacement string) *Replacer {
	return &Replacer{
		pattern:     pattern,
		replacement: []byte(replacement),
	}
}

func (r *Replacer) Replace(input []byte) []byte {
	return r.regexp.ReplaceAll(input, r.replacement)
}

func newReplacers() ([]*Replacer, error) {
	replacers := []*Replacer{
		newReplacer(`"github\.com/urfave/cli/v2"`, `"github.com/urfave/cli/v3"`),
		newReplacer(`"github\.com/urfave/cli/v2/altsrc"`, `"github.com/urfave/cli-altsrc/v3"`),
		newReplacer(`cli\.App`, `cli.Command`),
		newReplacer(`EnableBashCompletion`, `EnableShellCompletion`),
		newReplacer(`RunContext`, `Run`),
		newReplacer(`Subcommands`, `Commands`),
		newReplacer(`CustomAppHelpTemplate`, `CustomRootCommandHelpTemplate`),
		newReplacer(`cli\.NewApp\(\)`, `&cli.Command{}`),
		newReplacer(`\(([^ ]+?) \*cli\.Context\) error`, `(ctx context.Context, $1 *cli.Command) error`),
		newReplacer(`\(\*cli\.Context\) error`, `(context.Context, *cli.Command) error`),
		newReplacer(`ExitErrHandler = func\(\*cli\.Context, error\)`, `ExitErrHandler = func(context\.Context, *cli.Command, error)`),
		newReplacer(`\*cli\.Context`, `context.Context`),
		newReplacer(`EnvVars: \[\]string\{([^}]+)\}`, `Sources: cli.EnvVars($1)`),
		newReplacer(`EnvVars: \[\]string`, `Sources: cli.EnvVars`),
	}
	for _, r := range replacers {
		if err := r.Init(); err != nil {
			return nil, fmt.Errorf("init replacer: %w", err)
		}
	}
	return replacers, nil
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	replacers, err := newReplacers()
	if err != nil {
		return fmt.Errorf("new replacers: %w", err)
	}
	files, err := listFiles(ctx)
	if err != nil {
		return fmt.Errorf("list files: %w", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file, ".go") {
			// Filter *.go files
			continue
		}
		if err := fixFile(file, replacers); err != nil {
			return fmt.Errorf("fix a file %s: %w", file, err)
		}
	}
	if err := goModTidy(ctx); err != nil {
		return err
	}
	return nil
}

func fixFile(file string, replacers []*Replacer) error {
	code, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("read a file %s: %w", file, err)
	}
	code = fixCode(code, replacers)
	stat, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("stat a file %s: %w", file, err)
	}
	if err := os.WriteFile(file, code, stat.Mode()); err != nil {
		return fmt.Errorf("write a file %s: %w", file, err)
	}
	return nil
}

func fixCode(code []byte, replacers []*Replacer) []byte {
	for _, r := range replacers {
		code = r.Replace(code)
	}
	return code
}

func goModTidy(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "go", "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy: %w", err)
	}
	return nil
}

func listFiles(ctx context.Context) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "ls-files")
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("git ls-files: %w", err)
	}
	return strings.Split(strings.TrimSpace(buf.String()), "\n"), nil
}
