package go_trade_aggregation

import (
	"fmt"
	"time"

	"github.com/MathisWellmann/go_trade_aggregation/aggregator"
	"github.com/yourusername/trades-to-candles-benchmark-go/internal/models"
)

// Adapter implements TradeAggregator interface for go_trade_aggregation library
type Adapter struct {
	aggregator *aggregator.TimeAggregator
	timeframe  time.Duration
}

// NewAdapter creates a new adapter instance
func NewAdapter(timeframe time.Duration) (*Adapter, error) {
	if timeframe < time.Minute {
		return nil, fmt.Errorf("timeframe must be at least 1 minute")
	}

	return &Adapter{
		aggregator: aggregator.NewTimeAggregator(timeframe),
		timeframe:  timeframe,
	}, nil
}

// AddTrade adds a single trade to the aggregator
func (a *Adapter) AddTrade(trade *models.Trade) error {
	if err := trade.Validate(); err != nil {
		return fmt.Errorf("invalid trade: %w", err)
	}

	// Convert our trade model to library's trade model
	libTrade := aggregator.Trade{
		Timestamp: trade.Timestamp,
		Price:     trade.Price,
		Volume:    trade.Amount,
	}

	a.aggregator.AddTrade(libTrade)
	return nil
}

// GetCandles returns the current candles
func (a *Adapter) GetCandles() ([]*models.Candle, error) {
	// Get candles from the library
	libCandles := a.aggregator.GetCandles()

	// Convert library candles to our model
	candles := make([]*models.Candle, len(libCandles))
	for i, lc := range libCandles {
		candle := &models.Candle{
			Timestamp: lc.Timestamp,
			Open:      lc.Open,
			High:      lc.High,
			Low:       lc.Low,
			Close:     lc.Close,
			Volume:    lc.Volume,
		}

		if err := candle.Validate(); err != nil {
			return nil, fmt.Errorf("invalid candle generated: %w", err)
		}

		candles[i] = candle
	}

	return candles, nil
}
