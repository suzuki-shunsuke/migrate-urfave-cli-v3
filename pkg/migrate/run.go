package migrate

import (
	"context"
	"fmt"
	"strings"
)

func Run(ctx context.Context, args []string) error {
	replacers, err := newReplacers()
	if err != nil {
		return fmt.Errorf("new replacers: %w", err)
	}
	files := args[1:]
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
