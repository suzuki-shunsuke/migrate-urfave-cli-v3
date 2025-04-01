# migrate-urfave-cli-v3

Migrate github.com/urfave/cli/v2 to v3.
This tool doesn't aim to the complete migration.
Probably you need to fix code manually after running this tool, but this tool makes the migration easy

## How To Use

### Go

```sh
go run github.com/suzuki-shunsuke/migrate-urfave-cli-v3/cmd/migrate-urfave-cli-v3@latest
```

By default, this tool finds files by `git ls-files` and filters files by file extension `.go`.

You can also pass migrated files via command line arguments:

```sh
go install github.com/suzuki-shunsuke/migrate-urfave-cli-v3/cmd/migrate-urfave-cli-v3@latest
migrate-urfave-cli-v3 cmd/foo/main.go cmd/bar/main.go
```

### Shell Script

> [!CAUTION]
> Please use Go version.
> We'll remove the shell script when Go version is released.

Download and run the script [main.sh](main.sh).

```sh
curl -Lq https://raw.githubusercontent.com/suzuki-shunsuke/migrate-urfave-cli-v3/refs/heads/main/main.sh | bash
```

## Note

- https://cli.urfave.org/migrate-v2-to-v3/
- https://github.com/urfave/cli/discussions/2084
