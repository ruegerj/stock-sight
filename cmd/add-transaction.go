package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/spf13/cobra"
)

var FLAGS = map[string]string{
	"stock":       "Stock ticker (name)",
	"ppu":         "Price price per unit",
	"amount":      "How much stock you bought",
	"currency":    "The currency you payed (USD, EUR, CHF)",
	"date":        "When you bought the stock 'DD.MM.YYYY hh:mm:ss' or 'now'",
	"transaction": "'buy' or 'sell'",
}

type AddTransactionCmdParams struct {
	Ticker             string
	PricePerUnit       float64
	Currency           string
	Amount             float64
	Date               time.Time
	TransactionTypeBuy bool
}

func AddTransactionCmd(ctx context.Context, transactionRepo repository.TransactionRepository) CobraCommand {
	addBuyCmd := &cobra.Command{
		Use:   "add-trx",
		Short: "Adds a transaction for a stock to your portfolio",
		RunE: func(cmd *cobra.Command, args []string) error {
			params, err := ParseBuyCmdFlags(cmd)
			if err != nil {
				return err
			}

			_, err = transactionRepo.Create(ctx,
				params.Ticker,
				params.Amount,
				params.Currency,
				params.PricePerUnit,
				params.Date,
				params.TransactionTypeBuy)

			if err != nil {
				return err
			}

			fmt.Printf("Added the following transaction to the Portfolio\n")
			fmt.Printf("Ticker: {%s} |Ppu: {%f} {%s} |Amount: {%f} | Date: {%s} | Buy: {%t} \n",
				params.Ticker,
				params.PricePerUnit,
				params.Currency,
				params.Amount,
				params.Date.Format("02.01.2006 15:04:05"),
				params.TransactionTypeBuy)

			return nil
		},
	}

	registerCmdFlags(addBuyCmd)

	return GenericCommand{
		cmd:  addBuyCmd,
		path: "root add-trx",
	}
}

func registerCmdFlags(addBuyCmd *cobra.Command) {
	for flag, description := range FLAGS {
		addBuyCmd.PersistentFlags().String(flag, "", description)
	}
}

func ParseBuyCmdFlags(cmd *cobra.Command) (AddTransactionCmdParams, error) {
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
		return AddTransactionCmdParams{}, fmt.Errorf("error: \n%s", errorText)
	}

	return AddTransactionCmdParams{
		Ticker:             ticker,
		PricePerUnit:       ppu,
		Currency:           currency,
		Amount:             amount,
		Date:               date,
		TransactionTypeBuy: isBuy,
	}, nil

}
