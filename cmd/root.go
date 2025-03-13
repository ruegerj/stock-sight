package cmd

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewRootCmd(commands []CobraCommand, lc fx.Lifecycle, shutdowner fx.Shutdowner) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "stock-sight",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	registerChildCommands(rootCmd, commands)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := rootCmd.Execute(); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR while executing the CLI: '%s'", err)
				_ = shutdowner.Shutdown(fx.ExitCode(1))
			}
			_ = shutdowner.Shutdown()
			return nil
		},
	})

	return rootCmd
}

// registers all child cmd's in their correct hierarchy
func registerChildCommands(rootCmd *cobra.Command, children []CobraCommand) {
	for _, cmd := range children {
		parentalPath := cmd.Path()[:strings.LastIndex(cmd.Path(), " ")]
		if parentalPath == "root" {
			rootCmd.AddCommand(cmd.Command())
			continue
		}

		matchesParentPath := func(c CobraCommand) bool {
			return c.Path() == parentalPath
		}
		c2 := children[slices.IndexFunc(children, matchesParentPath)]
		c2.Command().AddCommand(cmd.Command())
	}
}
