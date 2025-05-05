package cmd

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart/timeserieslinechart"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/ruegerj/stock-sight/internal/repository"
	"github.com/ruegerj/stock-sight/internal/terminal"
	"github.com/spf13/cobra"
)

const timespanFlag = "timespan"

func ShowPriceHistoryCmd(ctx context.Context,
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
				return strings.ToUpper(stock.Ticker) == strings.ToUpper(ticker)
			})

			if !isStockTracked {
				return errors.New("Price history can only be shown for tracked stocks")
			}

			width, height := terminalAccessor.ResolveDimensions()
			timeFormatter := timeserieslinechart.WithXLabelFormatter(func(i int, f float64) string {
				t := time.Unix(int64(f), 0)
				return t.Format(time.DateOnly)
			})
			tslc := timeserieslinechart.New(width, height, timeFormatter)
			for i, v := range []float64{0, 4, 8, 10, 8, 4, 0, -4, -8, -10, -8, -4, 0} {
				date := time.Now().Add(time.Hour * time.Duration(24*i))
				tslc.Push(timeserieslinechart.TimePoint{Time: date, Value: v})
			}
			tslc.DrawBrailleAll()

			fmt.Println(tslc.View())
			return nil
		},
	}

	showPriceHistCmd.Flags().StringP(timespanFlag, "t", "", "Timespan for which the history shall be shown. Valid options are: d(ay), w(eek), m(onth), y(ear) and y2d (01.01.xxxx to now)")
	cobra.MarkFlagRequired(showPriceHistCmd.Flags(), timespanFlag)

	return GenericCommand{
		cmd:  showPriceHistCmd,
		path: "root hist",
	}
}

func isValidTimespanFlag(value string) bool {
	validTimeSpans := []string{"d", "w", "m", "y", "y2d"}
	return slices.Contains(validTimeSpans, value)
}
