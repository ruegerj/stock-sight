package cmd_test

import (
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestParseBuyCmdFlags(t *testing.T) {
	tests := []struct {
		name       string
		setupFlags map[string]string
		wantParams cmd.AddTransactionCmdParams
		wantErrMsg string
	}{
		{
			name: "valid buy transaction",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "USD",
				"date":        "24.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantParams: cmd.AddTransactionCmdParams{
				Ticker:             "AAPL",
				PricePerUnit:       150.50,
				Amount:             10,
				Currency:           "USD",
				Date:               time.Date(2025, 3, 24, 15, 4, 5, 0, time.UTC),
				TransactionTypeBuy: true,
			},
		},
		{
			name: "valid sell transaction",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "USD",
				"date":        "24.03.2025 15:04:05",
				"transaction": "sell",
			},
			wantParams: cmd.AddTransactionCmdParams{
				Ticker:             "AAPL",
				PricePerUnit:       150.50,
				Amount:             10,
				Currency:           "USD",
				Date:               time.Date(2025, 3, 24, 15, 4, 5, 0, time.UTC),
				TransactionTypeBuy: false,
			},
		},
		{
			name: "invalid ticker length",
			setupFlags: map[string]string{
				"stock":       "AA",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "USD",
				"date":        "24.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantErrMsg: "please enter a valid --stock",
		},
		{
			name: "invalid price per unit",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "A",
				"amount":      "10",
				"currency":    "GBP",
				"date":        "24.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantErrMsg: "please enter a valid price per unit (--ppu)",
		},
		{
			name: "invalid amount",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "A",
				"currency":    "USD",
				"date":        "24.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantErrMsg: "please enter a valid --amount",
		},
		{
			name: "invalid currency",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "GBP",
				"date":        "24.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantErrMsg: "please enter a valid --currency (USD, EUR, CHF)",
		},
		{
			name: "invalid date",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "USD",
				"date":        "32.03.2025 15:04:05",
				"transaction": "buy",
			},
			wantErrMsg: "please enter a valid date (dont forget parenthesis): --date \"DD.MM.YYYY hh:mm:ss\"",
		},
		{
			name: "invalid transaction",
			setupFlags: map[string]string{
				"stock":       "AAPL",
				"ppu":         "150.50",
				"amount":      "10",
				"currency":    "USD",
				"date":        "24.03.2025 15:04:05",
				"transaction": "invalid",
			},
			wantErrMsg: "please enter a valid transaction direction: --transaction (buy / sell)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command := setupTestCommand()
			//setup command flags
			for flag, value := range tt.setupFlags {
				err := command.Flags().Set(flag, value)
				assert.NoError(t, err, "failed to set flag")
			}

			//validate flags
			gotParams, err := cmd.ParseBuyCmdFlags(command)

			//validate error message or params
			if tt.wantErrMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantParams, gotParams)
		})
	}
}

func setupTestCommand() *cobra.Command {
	command := &cobra.Command{}
	command.Flags().String("stock", "", "")
	command.Flags().String("ppu", "", "")
	command.Flags().String("amount", "", "")
	command.Flags().String("currency", "", "")
	command.Flags().String("date", "", "")
	command.Flags().String("transaction", "", "")
	return command
}
