package utils

import (
	"strings"
	"testing"
	"time"
)

func TestTradeParser(t *testing.T) {
	// Sample trade data
	testData := `{"Timestamp":"2025-01-30T12:00:00Z","Price":50000.0,"Amount":1.5,"IsBuyer":true}
{"Timestamp":"2025-01-30T12:00:01Z","Price":50001.0,"Amount":0.5,"IsBuyer":false}
`
	parser := NewTradeParser(strings.NewReader(testData))

	// Test parsing single trade
	trade, err := parser.ParseTrade()
	if err != nil {
		t.Fatalf("Failed to parse trade: %v", err)
	}

	expectedTime := time.Date(2025, 1, 30, 12, 0, 0, 0, time.UTC)
	if !trade.Timestamp.Equal(expectedTime) {
		t.Errorf("Expected timestamp %v, got %v", expectedTime, trade.Timestamp)
	}

	if trade.Price != 50000.0 {
		t.Errorf("Expected price 50000.0, got %v", trade.Price)
	}

	// Test parsing all trades
	parser = NewTradeParser(strings.NewReader(testData))
	tradeChan, err := parser.ParseAll()
	if err != nil {
		t.Fatalf("Failed to start parsing all trades: %v", err)
	}

	tradeCount := 0
	for range tradeChan {
		tradeCount++
	}

	if tradeCount != 2 {
		t.Errorf("Expected 2 trades, got %d", tradeCount)
	}
}
