package benchmark

import (
	"testing"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

type mockAggregator struct {
	trades  []*models.Trade
	candles []*models.Candle
}

func (m *mockAggregator) AddTrade(trade *models.Trade) error {
	m.trades = append(m.trades, trade)
	return nil
}

func (m *mockAggregator) GetCandles() ([]*models.Candle, error) {
	if len(m.candles) == 0 {
		// Create a sample candle from trades
		if len(m.trades) > 0 {
			m.candles = append(m.candles, &models.Candle{
				Timestamp: m.trades[0].Timestamp,
				Open:      m.trades[0].Price,
				High:      m.trades[0].Price,
				Low:       m.trades[0].Price,
				Close:     m.trades[len(m.trades)-1].Price,
				Volume:    1.0,
			})
		}
	}
	return m.candles, nil
}

func TestTradesBenchmark(t *testing.T) {
	// Prepare test data
	now := time.Now().Truncate(time.Minute)
	trades := []*models.Trade{
		{
			Timestamp: now,
			Price:     100.0,
			Amount:    1.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(time.Second),
			Price:     101.0,
			Amount:    2.0,
			IsBuyer:   true,
		},
	}

	config := Config{
		LibraryName: "test_lib",
		Timeframe:   time.Minute,
		BatchSize:   10,
	}

	aggregator := &mockAggregator{}
	benchmark := NewTradesBenchmark(config, trades, aggregator)

	// Run benchmark
	result, err := benchmark.Run()
	if err != nil {
		t.Fatalf("Benchmark failed: %v", err)
	}

	// Verify results
	if result.LibraryName != config.LibraryName {
		t.Errorf("Expected library name %s, got %s", config.LibraryName, result.LibraryName)
	}

	if result.ProcessedTrades != int64(len(trades)) {
		t.Errorf("Expected %d processed trades, got %d", len(trades), result.ProcessedTrades)
	}

	if result.CandlesCreated != 1 {
		t.Errorf("Expected 1 candle created, got %d", result.CandlesCreated)
	}

	if result.ExecutionTime == 0 {
		t.Error("Execution time should not be zero")
	}
}
