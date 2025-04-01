package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	_ = cli.NewApp()
	app := &cli.App{
		Name:                  "example",
		EnableBashCompletion:  true,
		CustomAppHelpTemplate: "{{.Name}} - {{.Usage}}",
		ExitErrHandler: func(_ *cli.Context, err error) {
			if err != nil {
				os.Exit(1)
			}
		},
		Commands: []*cli.Command{
			{
				Name: "sub",
				Subcommands: []*cli.Command{
					{
						Name: "subsub",
						Action: func(c *cli.Context) error {
							name := c.String("name")
							age := c.Int("age")
							fmt.Fprintln(os.Stdout, name, age)
							return nil
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								EnvVars: []string{"NAME"},
							},
						},
					},
				},
			},
		},
	}
	app.RunContext(context.Background(), os.Args)
}
