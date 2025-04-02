package migrate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_fixCode(t *testing.T) { //nolint:funlen
	t.Parallel()
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "import",
			code: `import "github.com/urfave/cli/v2"`,
			want: `import "github.com/urfave/cli/v3"`,
		},
		{
			name: "import altsrc",
			code: `import "github.com/urfave/cli/v2/altsrc"`,
			want: `import "github.com/urfave/cli-altsrc/v3"`,
		},
		{
			name: "cli.App",
			code: `cli.App{}`,
			want: `cli.Command{}`,
		},
		{
			name: "*cli.Author",
			code: `*cli.Author`,
			want: `any`,
		},
		{
			name: "&cli.Author",
			code: `&cli.Author`,
			want: `any`,
		},
		{
			name: "cli.Author",
			code: `cli.Author`,
			want: `any`,
		},
		{
			name: "EnableBashCompletion",
			code: `EnableBashCompletion: true`,
			want: `EnableShellCompletion: true`,
		},
		{
			name: "RunContext",
			code: `app.RunContext`,
			want: `app.Run`,
		},
		{
			name: "Subcommands",
			code: `Subcommands`,
			want: `Commands`,
		},
		{
			name: "CustomAppHelpTemplate",
			code: `CustomAppHelpTemplate`,
			want: `CustomRootCommandHelpTemplate`,
		},
		{
			name: "cli.NewApp()",
			code: `cli.NewApp()`,
			want: `&cli.Command{}`,
		},
		{
			name: "(c *cli.Context) error",
			code: `(c *cli.Context) error`,
			want: `(ctx context.Context, c *cli.Command) error`,
		},
		{
			name: "(*cli.Context) error",
			code: `(*cli.Context) error`,
			want: `(context.Context, *cli.Command) error`,
		},
		{
			name: "ExitErrHandler ",
			code: `ExitErrHandler = func(*cli.Context, error)`,
			want: `ExitErrHandler = func(context.Context, *cli.Command, error)`,
		},
		{
			name: "*cli.Context",
			code: `*cli.Context`,
			want: `context.Context`,
		},
		{
			name: "App.ToFishCompletion",
			code: `c.App.ToFishCompletion`,
			want: `c.ToFishCompletion`,
		},
		{
			name: "EnvVars",
			code: `EnvVars: []string{"FOO", "BAR"}`,
			want: `Sources: cli.EnvVars("FOO", "BAR")`,
		},
		{
			name: "EnvVars 2",
			code: `EnvVars: []string{
	"FOO", "BAR",
}`,
			want: `Sources: cli.EnvVars(
	"FOO", "BAR",
)`,
		},
	}
	replacers, err := newReplacers()
	if err != nil {
		t.Fatalf("newReplacers() = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := fixCode([]byte(tt.code), replacers)
			gotS := string(got)
			if diff := cmp.Diff(gotS, tt.want); diff != "" {
				t.Error("(- got, + want)", diff)
			}
		})
	}
}
