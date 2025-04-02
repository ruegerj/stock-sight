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
		Short: "Terminal based stock tracker",
		Long:  "Enables one to track the price of their favorite stocks, while persisting all performed stock transactions.",
	}

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
