package cmd

import (
	"context"
	"fmt"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/spf13/cobra"
)

func BuysCmd(ctx context.Context, transactionRepo repository.TransactionRepository) CobraCommand {
	buysCmd := &cobra.Command{
		Use:   "buys",
		Short: "See all your buys",
		RunE: func(cmd *cobra.Command, args []string) error {

			transactions, err := transactionRepo.GetAll(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("You registered the following transactions to your Portfolio:\n")
			for _, transaction := range transactions {

				transactionType := "sell"
				if transaction.IsBuy {
					transactionType = "buy"
				}
				fmt.Printf("Ticker: {%s} |Ppu: {%f} {%s} |Amount: {%f} | Date: {%s} | Buy: {%s} \n", transaction.Ticker, transaction.PricePerUnit, transaction.Currency, transaction.Amount, transaction.Date.Format("02.01.2006 15:04:05"), transactionType)
			}

			return nil
		},
	}

	return GenericCommand{
		cmd:  buysCmd,
		path: "root buys",
	}
}
