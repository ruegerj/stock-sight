package cmd

import (
	"context"
	"fmt"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

func AddBuyCmd(ctx context.Context, transactionRepo repository.TransactionRepository) CobraCommand {
	addBuyCmd := &cobra.Command{
		Use:   "add-buy",
		Short: "Adds a buy for a stock to your portfolio",
		RunE: func(cmd *cobra.Command, args []string) error {

			errorText := ""
			ticker := cmd.Flag("stock").Value.String()
			if len(ticker) < 3 {
				errorText += "please enter a valid --stock \n"
			}

			ppu, err := strconv.ParseFloat(cmd.Flag("ppu").Value.String(), 64)
			if err != nil || ppu < 0 {
				errorText += "please enter a valid price per unit (--ppu) \n"
			}

			amount, err := strconv.ParseFloat(cmd.Flag("amount").Value.String(), 64)
			if err != nil || amount < 0 {
				errorText += "please enter a valid --amount \n"
			}

			currency := cmd.Flag("currency").Value.String()
			if currency != "USD" && currency != "EUR" && currency != "CHF" {
				errorText += "please enter a valid --currency (USD, EUR, CHF) \n"
			}

			dateStr := cmd.Flag("date").Value.String()
			date, err := time.Parse("02.01.2006 15:04:05", dateStr)
			if nil != err {
				if dateStr == "now" {
					date = time.Now()
				} else {
					errorText += "please enter a valid date (dont forget parenthesis): --date \"DD.MM.YYYY hh:mm:ss\" \n" + err.Error()
				}
			}

			transactionStr := cmd.Flag("transaction").Value.String()
			if transactionStr != "buy" && transactionStr != "sell" {
				errorText += "please enter a valid transaction direction: --transaction (buy / sell) \n"
			}
			isBuy := transactionStr == "buy"

			if errorText != "" {
				fmt.Printf("Error: \n%s", errorText)
				return nil
			}

			_, err = transactionRepo.Create(ctx, ticker, amount, currency, ppu, date, isBuy)
			if err != nil {
				return err
			}

			fmt.Printf("Added the following transaction to the Portfolio\n")
			fmt.Printf("Ticker: {%s} |Ppu: {%f} {%s} |Amount: {%f} | Date: {%s} | Buy: {%s} \n", ticker, ppu, currency, amount, date.Format("02.01.2006 15:04:05"), transactionStr)

			return nil
		},
	}

	addBuyCmd.PersistentFlags().String("stock", "", "Stock ticker (name)")
	addBuyCmd.PersistentFlags().String("ppu", "", "Price price per unit")
	addBuyCmd.PersistentFlags().String("currency", "", "The currency you payed (USD, EUR, CHF)")
	addBuyCmd.PersistentFlags().String("amount", "", "How much stock you bought")
	addBuyCmd.PersistentFlags().String("date", "", "when you bought the stock 'DD.MM.YYYY hh:mm:ss' or 'now'")
	addBuyCmd.PersistentFlags().String("transaction", "", "'buy' or 'sell'")

	return GenericCommand{
		cmd:  addBuyCmd,
		path: "root add-buy",
	}
}
