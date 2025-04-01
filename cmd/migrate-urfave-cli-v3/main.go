package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/suzuki-shunsuke/migrate-urfave-cli-v3/pkg/migrate"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	return migrate.Run(ctx, args) //nolint:wrapcheck
}
