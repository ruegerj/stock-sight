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

	// register cmd's in their correct hierarchy
	for _, cmd := range commands {
		parentalPath := cmd.Path()[:strings.LastIndex(cmd.Path(), " ")]
		if parentalPath == "root" {
			rootCmd.AddCommand(cmd.Command())
			continue
		}

		matchesParentPath := func(c CobraCommand) bool {
			return c.Path() == parentalPath
		}
		c2 := commands[slices.IndexFunc(commands, matchesParentPath)]
		c2.Command().AddCommand(cmd.Command())
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := rootCmd.Execute(); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR while executing the CLI: '%s'", err)
				shutdowner.Shutdown(fx.ExitCode(1))
			}
			shutdowner.Shutdown()
			return nil
		},
	})

	return rootCmd
}
