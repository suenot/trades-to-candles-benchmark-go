package benchmark

import (
	"fmt"
	"runtime"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/interfaces"
	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

// Config holds benchmark configuration
type Config struct {
	LibraryName string
	Timeframe   time.Duration
	BatchSize   int // Number of trades to process in one batch
}

// BenchmarkResult stores the results of a single benchmark run
type BenchmarkResult struct {
	LibraryName     string
	ExecutionTime   time.Duration
	MemoryUsage     uint64
	ProcessedTrades int64
	Error           error
	// Additional metrics
	MaxMemoryUsage uint64
	CPUUsage       float64
	CandlesCreated int
}

// TradesBenchmark implements benchmark for trades to candles conversion
type TradesBenchmark struct {
	config     Config
	trades     []*models.Trade
	aggregator interfaces.TradeAggregator
}

// NewTradesBenchmark creates a new benchmark instance
func NewTradesBenchmark(config Config, trades []*models.Trade, aggregator interfaces.TradeAggregator) *TradesBenchmark {
	return &TradesBenchmark{
		config:     config,
		trades:     trades,
		aggregator: aggregator,
	}
}

// Run executes the benchmark
func (b *TradesBenchmark) Run() (*BenchmarkResult, error) {
	result := &BenchmarkResult{
		LibraryName: b.config.LibraryName,
	}

	// Get initial memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	initialMem := memStats.Alloc

	startTime := time.Now()

	// Process trades in batches
	for i := 0; i < len(b.trades); i += b.config.BatchSize {
		end := i + b.config.BatchSize
		if end > len(b.trades) {
			end = len(b.trades)
		}

		// Process batch
		for _, trade := range b.trades[i:end] {
			if err := b.aggregator.AddTrade(trade); err != nil {
				return nil, fmt.Errorf("failed to process trade: %w", err)
			}
		}

		// Update memory stats
		runtime.ReadMemStats(&memStats)
		currentMem := memStats.Alloc
		if currentMem-initialMem > result.MaxMemoryUsage {
			result.MaxMemoryUsage = currentMem - initialMem
		}
	}

	// Get final candles
	candles, err := b.aggregator.GetCandles()
	if err != nil {
		return nil, fmt.Errorf("failed to get candles: %w", err)
	}

	// Calculate final metrics
	result.ExecutionTime = time.Since(startTime)
	result.ProcessedTrades = int64(len(b.trades))
	result.CandlesCreated = len(candles)

	// Get final memory stats
	runtime.ReadMemStats(&memStats)
	result.MemoryUsage = memStats.Alloc - initialMem

	return result, nil
}
