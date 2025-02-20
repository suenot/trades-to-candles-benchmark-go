package go_trade_aggregation

import (
	"fmt"
	"time"

	gta "github.com/MathisWellmann/go_trade_aggregation"
	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

// Adapter implements TradeAggregator interface for go_trade_aggregation library
type Adapter struct {
	timeframe time.Duration
	trades    []*gta.Trade
}

// NewAdapter creates a new adapter instance
func NewAdapter(timeframe time.Duration) (*Adapter, error) {
	if timeframe < time.Minute {
		return nil, fmt.Errorf("timeframe must be at least 1 minute")
	}

	return &Adapter{
		timeframe: timeframe,
		trades:    make([]*gta.Trade, 0),
	}, nil
}

// AddTrade adds a single trade to the aggregator
func (a *Adapter) AddTrade(trade *models.Trade) error {
	if err := trade.Validate(); err != nil {
		return fmt.Errorf("invalid trade: %w", err)
	}

	// Convert our trade model to library's trade model
	libTrade := &gta.Trade{
		Timestamp: trade.Timestamp.UnixMilli(), // Convert to milliseconds
		Price:     trade.Price,
		Size:      trade.Amount * (map[bool]float64{true: 1, false: -1})[trade.IsBuyer], // Negative if sell
	}

	a.trades = append(a.trades, libTrade)
	return nil
}

// GetCandles returns the current candles
func (a *Adapter) GetCandles() ([]*models.Candle, error) {
	if len(a.trades) == 0 {
		return []*models.Candle{}, nil
	}

	// Convert timeframe to milliseconds
	timeframeMillis := int64(a.timeframe.Milliseconds())

	// Get candles from the library
	libCandles := gta.AggTime(a.trades, timeframeMillis)

	// Convert library candles to our model
	candles := make([]*models.Candle, len(libCandles))
	for i, lc := range libCandles {
		candle := &models.Candle{
			Timestamp: time.UnixMilli(lc.Timestamp),
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
