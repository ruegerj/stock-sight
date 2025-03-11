package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/ruegerj/stock-sight/internal/tui"
	"github.com/spf13/cobra"
)

func ShowTuiCmd(ctx context.Context, transactionRepo repository.TransactionRepository) CobraCommand {
	showTuiCmd := &cobra.Command{
		Use:   "show-tui",
		Short: "show graphical tui",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := tui.Start(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return nil
		},
	}

	return GenericCommand{
		cmd:  showTuiCmd,
		path: "root show-tui",
	}
}
