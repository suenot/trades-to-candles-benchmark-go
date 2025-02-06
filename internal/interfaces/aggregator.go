package interfaces

import "github.com/yourusername/trades-to-candles-benchmark-go/internal/models"

// TradeAggregator defines interface for trade to candle conversion
type TradeAggregator interface {
	// AddTrade adds a single trade to the aggregator
	AddTrade(trade *models.Trade) error

	// GetCandles returns the current candles
	GetCandles() ([]*models.Candle, error)
}

// CandleAggregator defines interface for candle to candle conversion
type CandleAggregator interface {
	// AddCandle adds a single candle to the aggregator
	AddCandle(candle *models.Candle) error

	// GetCandles returns the aggregated candles
	GetCandles() ([]*models.Candle, error)
}
