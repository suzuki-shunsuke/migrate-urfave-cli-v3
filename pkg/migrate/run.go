package migrate

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

type Runner struct {
	LDFlags *LDFlags
}

type LDFlags struct {
	Version string
	Commit  string
}

func (r *Runner) Run(ctx context.Context, args []string) error {
	cmd := &cli.Command{
		Name:        "migrate-urfave-cli-v3",
		Usage:       "Migrate urfave/cli v2 to v3",
		Description: "Migrate urfave/cli v2 to v3",
		Version:     r.LDFlags.Version + " (" + r.LDFlags.Commit + ")",
		Action:      r.action,
	}
	return cmd.Run(ctx, args) //nolint:wrapcheck
}

func (r *Runner) action(ctx context.Context, cmd *cli.Command) error {
	replacers, err := newReplacers()
	if err != nil {
		return fmt.Errorf("new replacers: %w", err)
	}
	files := cmd.Args().Slice()
	if len(files) == 0 {
		a, err := listFiles(ctx)
		if err != nil {
			return fmt.Errorf("list files: %w", err)
		}
		files = a
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
