package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/spf13/cobra"
)

func TrackCmd(ctx context.Context, repo repository.StockRepository) CobraCommand {
	trackCmd := &cobra.Command{
		Use:   "track",
		Short: "Track a stock by its ticker symbol",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("usage: track <stock>. You must provide the ticker symbol of the stock you want to track. For example, 'track AAPL' to track Apple stock")
			}
			ticker := args[0]

			trackedStocks, err := repo.GetTrackedStocks(ctx)
			if err != nil {
				return err
			}

			for _, tracked := range trackedStocks {
				if tracked.Ticker == ticker {
					fmt.Printf("Stock %s is already being tracked.\n", ticker)
					return nil
				}
			}

			_, err = repo.AddTrackedStock(ctx, ticker, time.Now())
			if err != nil {
				return err
			}

			fmt.Printf("Stock %s is now being tracked.\n", ticker)
			return nil
		},
	}
	trackCmd.PersistentFlags().String("stock", "", "Stock ticker (name)")

	return GenericCommand{
		cmd:  trackCmd,
		path: "root track",
	}
}
