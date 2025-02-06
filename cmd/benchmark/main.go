package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/adapters/go_trade_aggregation"
	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
	"github.com/suenot/trades-to-candles-benchmark-go/internal/utils"
	"github.com/suenot/trades-to-candles-benchmark-go/pkg/benchmark"
)

func main() {
	// Parse command line flags
	dataFile := flag.String("data", "", "Path to trades data file")
	timeframe := flag.Duration("timeframe", time.Minute, "Candle timeframe")
	batchSize := flag.Int("batch", 1000, "Batch size for processing")
	outputDir := flag.String("output", "benchmarks/results", "Output directory for results")
	flag.Parse()

	if *dataFile == "" {
		log.Fatal("Please provide path to data file using -data flag")
	}

	// Initialize ZIP reader
	zipReader, err := utils.NewZipReader(*dataFile)
	if err != nil {
		log.Fatalf("Failed to open ZIP file: %v", err)
	}
	defer zipReader.Close()

	// Get trades data
	tradesFile, err := zipReader.GetFirstFile()
	if err != nil {
		log.Fatalf("Failed to get trades file: %v", err)
	}
	defer tradesFile.Close()

	// Create trade parser
	parser := utils.NewTradeParser(tradesFile)
	trades, err := parser.ParseAll()
	if err != nil {
		log.Fatalf("Failed to parse trades: %v", err)
	}

	// Collect all trades from channel
	var allTrades []*models.Trade
	for trade := range trades {
		allTrades = append(allTrades, trade)
	}

	fmt.Printf("Loaded %d trades from file\n", len(allTrades))

	// Create go_trade_aggregation adapter
	adapter, err := go_trade_aggregation.NewAdapter(*timeframe)
	if err != nil {
		log.Fatalf("Failed to create adapter: %v", err)
	}

	// Create and run benchmark
	config := benchmark.Config{
		LibraryName: "go_trade_aggregation",
		Timeframe:   *timeframe,
		BatchSize:   *batchSize,
	}

	bench := benchmark.NewTradesBenchmark(config, allTrades, adapter)
	result, err := bench.Run()
	if err != nil {
		log.Fatalf("Benchmark failed: %v", err)
	}

	// Save and display results
	if err := benchmark.SaveResults([]*benchmark.BenchmarkResult{result}, *outputDir); err != nil {
		log.Printf("Warning: Failed to save results: %v", err)
	}

	benchmark.CompareResults([]*benchmark.BenchmarkResult{result})
}
