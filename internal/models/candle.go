package models

import "time"

// Candle represents OHLCV data for a time period
type Candle struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}
