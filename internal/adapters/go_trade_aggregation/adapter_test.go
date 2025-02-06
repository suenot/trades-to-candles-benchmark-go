package go_trade_aggregation

import (
	"testing"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

func TestAdapter(t *testing.T) {
	// Create adapter with 1-minute timeframe
	adapter, err := NewAdapter(time.Minute)
	if err != nil {
		t.Fatalf("Failed to create adapter: %v", err)
	}

	// Test data
	now := time.Now().Truncate(time.Minute)
	trades := []*models.Trade{
		{
			Timestamp: now,
			Price:     100.0,
			Amount:    1.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(10 * time.Second),
			Price:     101.0,
			Amount:    2.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(20 * time.Second),
			Price:     99.0,
			Amount:    1.5,
			IsBuyer:   false,
		},
	}

	// Add trades
	for _, trade := range trades {
		if err := adapter.AddTrade(trade); err != nil {
			t.Errorf("Failed to add trade: %v", err)
		}
	}

	// Get candles
	candles, err := adapter.GetCandles()
	if err != nil {
		t.Fatalf("Failed to get candles: %v", err)
	}

	if len(candles) != 1 {
		t.Fatalf("Expected 1 candle, got %d", len(candles))
	}

	candle := candles[0]

	// Verify candle data
	if candle.Open != 100.0 {
		t.Errorf("Expected open price 100.0, got %v", candle.Open)
	}
	if candle.High != 101.0 {
		t.Errorf("Expected high price 101.0, got %v", candle.High)
	}
	if candle.Low != 99.0 {
		t.Errorf("Expected low price 99.0, got %v", candle.Low)
	}
	if candle.Volume != 4.5 {
		t.Errorf("Expected volume 4.5, got %v", candle.Volume)
	}
}
