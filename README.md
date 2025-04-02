# migrate-urfave-cli-v3

Migrate github.com/urfave/cli/v2 to v3.
This tool doesn't aim to the complete migration.
Probably you need to fix code manually after running this tool, but this tool makes the migration easy

## How To Use

1. [Install migrate-urfave-cli-v3](INSTALL.md)
2. Run migrate-urfave-cli-v3

```sh
migrate-urfave-cli-v3
```

By default, this tool finds files by `git ls-files` and filters files by file extension `.go`.
You can also pass migrated files via command line arguments:

```sh
migrate-urfave-cli-v3 cmd/foo/main.go cmd/bar/main.go
```

### go run

Run by `go run`:

```sh
go run github.com/suzuki-shunsuke/migrate-urfave-cli-v3/cmd/migrate-urfave-cli-v3@latest
```

## Note

- https://cli.urfave.org/migrate-v2-to-v3/
- https://github.com/urfave/cli/discussions/2084
