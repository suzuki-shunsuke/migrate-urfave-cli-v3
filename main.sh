#!/usr/bin/env bash

set -euxo pipefail

patterns=(
	's|"github\.com/urfave/cli/v2"|"github.com/urfave/cli/v3"|'
	's|"github\.com/urfave/cli/v2/altsrc"|"github.com/urfave/cli-altsrc/v3"|'
	's|cli\.App|cli.Command|'
	's|EnableBashCompletion|EnableShellCompletion|'
	's|RunContext|Run|'
	's|Subcommands|Commands|'
	's|CustomAppHelpTemplate|CustomRootCommandHelpTemplate|'
	's|cli\.NewApp\(\)|\&cli.Command{}|'
	's|\(([^ ]+?) \*cli\.Context\) error|(ctx context.Context, \1 *cli.Command) error|'
	's|\(\*cli\.Context\) error|(context.Context, *cli.Command) error|'
	's|ExitErrHandler = func\(\*cli\.Context, error\)|ExitErrHandler = func(context\.Context, *cli.Command, error)|'
	's|\*cli\.Context|context.Context|'
	's|EnvVars: \[\]string\{([^}]+)\}|Sources: cli.EnvVars(\1)|'
	's|EnvVars: \[\]string|Sources: cli.EnvVars|'
)

while read -r file; do
	for p in "${patterns[@]}"; do
		sed -Ei "$p" "$file"
	done
	sed -Ei "/^	*Compiled: /d" "$file"
done < <(git ls-files | grep -E '\.go$')

go mod tidy
