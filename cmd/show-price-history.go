package cmd

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/NimbleMarkets/ntcharts/linechart/timeserieslinechart"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/ruegerj/stock-sight/internal/stocks"
	"github.com/ruegerj/stock-sight/internal/terminal"
	"github.com/spf13/cobra"
)

const timespanFlag = "timespan"

func ShowPriceHistoryCmd(
	ctx context.Context,
	stockRepository repository.StockRepository,
	terminalAccessor terminal.TerminalAccessor) CobraCommand {
	showPriceHistCmd := &cobra.Command{
		Use:   "hist <ticker>",
		Short: "Shows the price histogram of the given stock",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			timespan := cmd.Flag(timespanFlag).Value
			if timespan == nil || timespan.String() == "" {
				return errors.New("Must supply value for timespan flag")
			}

			if !isValidTimespanFlag(timespan.String()) {
				return errors.New("Must supply a valid value for the timespan flag")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ticker := args[0]

			trackedStocks, err := stockRepository.GetTrackedStocks(ctx)
			if err != nil {
				return err
			}

			isStockTracked := slices.ContainsFunc(trackedStocks, func(stock queries.TrackedStock) bool {
				return strings.EqualFold(stock.Ticker, ticker)
			})

			if !isStockTracked {
				return errors.New("Price history can only be shown for tracked stocks")
			}

			timespan := cmd.Flag(timespanFlag).Value.String()

			width, height := terminalAccessor.ResolveDimensions()
			var timestampFormatter linechart.LabelFormatter = func(i int, unix float64) string {
				t := time.Unix(int64(unix), 0)
				return t.Format(time.DateOnly)
			}
			if timespan == "d" {
				timestampFormatter = func(i int, unix float64) string {
					t := time.Unix(int64(unix), 0)
					return t.Format(time.TimeOnly)
				}
			}

			tslc := timeserieslinechart.New(
				width,
				height,
				timeserieslinechart.WithXLabelFormatter(timestampFormatter))

			generator := stocks.NewFakeDataGenerator(150.56, 0.05, time.Now().UnixNano())
			stockData := generator.GenerateForTimeSpan(timespan)

			for _, point := range stockData {
				tslc.Push(timeserieslinechart.TimePoint{Time: point.Timestamp, Value: point.Price})
			}

			tslc.DrawBrailleAll()
			fmt.Printf("\nPrice history of %q (%s)", ticker, timespan)
			fmt.Println(tslc.View())
			return nil
		},
	}

	showPriceHistCmd.Flags().StringP(timespanFlag, "t", "d", "Timespan for which the history shall be shown (default=d). Valid options are: d(ay), w(eek), m(onth), y(ear) and y2d (01.01.xxxx to now)")

	return GenericCommand{
		cmd:  showPriceHistCmd,
		path: "root hist",
	}
}

func isValidTimespanFlag(value string) bool {
	validTimeSpans := []string{"d", "w", "m", "y", "y2d"}
	return slices.Contains(validTimeSpans, value)
}
