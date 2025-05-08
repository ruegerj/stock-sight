package cmd_test

import (
	"context"
	"testing"
	"time"

	"github.com/ruegerj/stock-sight/cmd"
	"github.com/ruegerj/stock-sight/internal/queries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTrackCommand(t *testing.T) {
	mockRepo := new(MockStockRepository)

	trackCmd := cmd.TrackCmd(context.Background(), mockRepo).Command()

	mockRepo.On("GetTrackedStocks", mock.Anything).Return([]queries.TrackedStock{
		{Ticker: "AAPL", DateAdded: time.Now()},
		{Ticker: "GOOGL", DateAdded: time.Now()},
	}, nil).Once()

	t.Run("Track META", func(t *testing.T) {
		mockRepo.On("AddTrackedStock", mock.Anything, "META", mock.Anything).Return(
			queries.TrackedStock{Ticker: "META", DateAdded: time.Now()},
			nil,
		).Once()

		err := trackCmd.RunE(trackCmd, []string{"META"})
		assert.NoError(t, err)

		mockRepo.AssertCalled(t, "AddTrackedStock", mock.Anything, "META", mock.Anything)
	})

	t.Run("Track Duplicate AAPL", func(t *testing.T) {
		mockRepo.On("GetTrackedStocks", mock.Anything).Return([]queries.TrackedStock{
			{Ticker: "AAPL", DateAdded: time.Now()},
		}, nil).Once()

		mockRepo.On("AddTrackedStock", mock.Anything, "AAPL", mock.Anything).Return(
			queries.TrackedStock{Ticker: "AAPL", DateAdded: time.Now()},
			nil,
		).Once()

		err := trackCmd.RunE(trackCmd, []string{"AAPL"})
		assert.NoError(t, err)

		mockRepo.AssertNotCalled(t, "AddTrackedStock", mock.Anything, "AAPL", mock.Anything)
	})

	t.Run("Track with no ticker", func(t *testing.T) {
		err := trackCmd.RunE(trackCmd, []string{})
		assert.Error(t, err)
		assert.Equal(t, "usage: track <stock>. You must provide the ticker symbol of the stock you want to track. For example, 'track AAPL' to track Apple stock", err.Error())
	})

	t.Run("Track GOOGL", func(t *testing.T) {
		mockRepo.On("AddTrackedStock", mock.Anything, "GOOGL", mock.Anything).Return(
			queries.TrackedStock{Ticker: "GOOGL", DateAdded: time.Now()},
			nil,
		)

		mockRepo.On("GetTrackedStocks", mock.Anything).Return([]queries.TrackedStock{
			{Ticker: "AAPL", DateAdded: time.Now()},
		}, nil)

		err := trackCmd.RunE(trackCmd, []string{"GOOGL"})
		assert.NoError(t, err)

		mockRepo.AssertCalled(t, "AddTrackedStock", mock.Anything, "GOOGL", mock.Anything)
	})
}
