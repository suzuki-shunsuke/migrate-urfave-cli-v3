package migrate

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

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
