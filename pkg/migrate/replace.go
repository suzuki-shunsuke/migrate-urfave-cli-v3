package migrate

import (
	"fmt"
	"regexp"
)

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
