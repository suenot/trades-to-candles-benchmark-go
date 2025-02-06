package models

import (
	"fmt"
	"time"
)

// Candle represents OHLCV data for a time period
type Candle struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// Validate checks if the candle data is valid
func (c *Candle) Validate() error {
	if c.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}
	if c.Open <= 0 {
		return fmt.Errorf("open price must be positive, got: %v", c.Open)
	}
	if c.High < c.Open || c.High < c.Low || c.High < c.Close {
		return fmt.Errorf("high price must be highest value")
	}
	if c.Low > c.Open || c.Low > c.High || c.Low > c.Close {
		return fmt.Errorf("low price must be lowest value")
	}
	if c.Close <= 0 {
		return fmt.Errorf("close price must be positive, got: %v", c.Close)
	}
	if c.Volume < 0 {
		return fmt.Errorf("volume cannot be negative, got: %v", c.Volume)
	}
	return nil
}

// String returns a string representation of the candle
func (c *Candle) String() string {
	return fmt.Sprintf("Candle{Time: %v, O: %v, H: %v, L: %v, C: %v, V: %v}",
		c.Timestamp, c.Open, c.High, c.Low, c.Close, c.Volume)
}
