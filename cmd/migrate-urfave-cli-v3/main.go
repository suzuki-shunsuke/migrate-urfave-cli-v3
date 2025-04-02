package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/suzuki-shunsuke/migrate-urfave-cli-v3/pkg/migrate"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	runner := &migrate.Runner{
		LDFlags: &migrate.LDFlags{
			Version: version,
			Commit:  commit,
		},
	}
	return runner.Run(ctx, args) //nolint:wrapcheck
}
