package migrate

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func listFiles(ctx context.Context) ([]string, error) {
	cmd := exec.CommandContext(ctx, "git", "ls-files")
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("git ls-files: %w", err)
	}
	return strings.Split(strings.TrimSpace(buf.String()), "\n"), nil
}
