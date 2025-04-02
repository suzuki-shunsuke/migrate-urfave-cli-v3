package migrate

import (
	"fmt"
	"regexp"
)

func newReplacers() ([]*Replacer, error) {
	replacers := []*Replacer{
		newReplacer(`"github\.com/urfave/cli/v2"`, `"github.com/urfave/cli/v3"`),
		newReplacer(`"github\.com/urfave/cli/v2/altsrc"`, `"github.com/urfave/cli-altsrc/v3"`),
		newReplacer(`"github\.com/suzuki-shunsuke/urfave-cli-help-all/helpall"`, `"github.com/suzuki-shunsuke/urfave-cli-v3-help-all/helpall"`),
		newReplacer(`cli\.App\b`, `cli.Command`),
		newReplacer(`[*&]?cli\.Author\b`, `any`),
		newReplacer(`\bEnableBashCompletion\b`, `EnableShellCompletion`),
		newReplacer(`\bRunContext\b`, `Run`),
		newReplacer(`\bSubcommands\b`, `Commands`),
		newReplacer(`\bCustomAppHelpTemplate\b`, `CustomRootCommandHelpTemplate`),
		newReplacer(`\.App\.ToFishCompletion\b`, `.ToFishCompletion`),
		newReplacer(`\bcli\.NewApp\(\)`, `&cli.Command{}`),
		newReplacer(`\(([^ ]+?) \*cli\.Context\) error\b`, `(ctx context.Context, $1 *cli.Command) error`),
		newReplacer(`\(\*cli\.Context\) error\b`, `(context.Context, *cli.Command) error`),
		newReplacer(`\bExitErrHandler = func\(\*cli\.Context, error\)`, `ExitErrHandler = func(context.Context, *cli.Command, error)`),
		newReplacer(`\*cli\.Context\b`, `context.Context`),
		newReplacer(`\bEnvVars: \[\]string\{([^}]+)\}`, `Sources: cli.EnvVars($1)`),
		newReplacer(`\bEnvVars: \[\]string\b`, `Sources: cli.EnvVars`),
		newReplacer(`\bEnvVars: `, `Sources: `),
	}
	for _, r := range replacers {
		if err := r.Init(); err != nil {
			return nil, fmt.Errorf("init replacer: %w", err)
		}
	}
	return replacers, nil
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
