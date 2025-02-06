package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/aggregator"
	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
	"github.com/suenot/trades-to-candles-benchmark-go/pkg/benchmark"
)

func main() {
	// Parse command line flags
	dataPath := flag.String("data", "", "Path to CSV file with trade data")
	flag.Parse()

	if *dataPath == "" {
		log.Fatal("Data path is required")
	}

	// Open CSV file
	file, err := os.Open(*dataPath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Read all trades
	trades := make([]*models.Trade, 0)
	var firstTimestamp, lastTimestamp time.Time

	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		// Parse CSV record
		// Format: id,price,amount,quoteAmount,timestamp,isBuyer,isMaker
		if len(record) != 7 {
			log.Printf("Invalid record format: %v", record)
			continue
		}

		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Failed to parse price: %v", err)
			continue
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Printf("Failed to parse amount: %v", err)
			continue
		}

		timestamp, err := strconv.ParseInt(record[4], 10, 64)
		if err != nil {
			log.Printf("Failed to parse timestamp: %v", err)
			continue
		}

		isBuyer, err := strconv.ParseBool(record[5])
		if err != nil {
			log.Printf("Failed to parse isBuyer: %v", err)
			continue
		}

		tradeTime := time.UnixMicro(timestamp)
		trade := &models.Trade{
			Timestamp: tradeTime,
			Price:     price,
			Amount:    amount,
			IsBuyer:   isBuyer,
		}

		if err := trade.Validate(); err != nil {
			log.Printf("Invalid trade: %v", err)
			continue
		}

		trades = append(trades, trade)

		// Update first and last timestamp
		if firstTimestamp.IsZero() || tradeTime.Before(firstTimestamp) {
			firstTimestamp = tradeTime
		}
		if tradeTime.After(lastTimestamp) {
			lastTimestamp = tradeTime
		}
	}

	fmt.Printf("Loaded %d trades\n", len(trades))
	fmt.Printf("Time range: from %v to %v\n", firstTimestamp, lastTimestamp)
	fmt.Printf("Total duration: %v\n", lastTimestamp.Sub(firstTimestamp))

	// Run benchmark with our MinuteAggregator
	results := make([]*benchmark.BenchmarkResult, 0)

	// Create and test MinuteAggregator
	minuteAggregator := aggregator.NewMinuteAggregator()

	config := benchmark.Config{
		LibraryName: "minute_aggregator",
		Timeframe:   time.Minute,
		BatchSize:   1000,
	}

	bench := benchmark.NewTradesBenchmark(config, trades, minuteAggregator)
	result, err := bench.Run()
	if err != nil {
		log.Printf("Benchmark failed for minute_aggregator: %v", err)
	} else {
		results = append(results, result)

		// Get candles and analyze them
		candles, err := minuteAggregator.GetCandles()
		if err != nil {
			log.Printf("Failed to get candles: %v", err)
		} else {
			fmt.Printf("\nCandle Analysis:\n")
			fmt.Printf("Number of candles: %d\n", len(candles))
			if len(candles) > 0 {
				fmt.Printf("First candle: %v\n", candles[0])
				fmt.Printf("Last candle: %v\n", candles[len(candles)-1])

				// Check for gaps in candles
				expectedCandles := int(lastTimestamp.Sub(firstTimestamp).Minutes())
				if len(candles) != expectedCandles {
					fmt.Printf("Warning: Expected %d candles, but got %d\n", expectedCandles, len(candles))
				}

				// Print first 5 and last 5 candles
				fmt.Printf("\nFirst 5 candles:\n")
				for i := 0; i < 5 && i < len(candles); i++ {
					fmt.Printf("%v\n", candles[i])
				}

				fmt.Printf("\nLast 5 candles:\n")
				for i := len(candles) - 5; i < len(candles); i++ {
					if i >= 0 {
						fmt.Printf("%v\n", candles[i])
					}
				}
			}
		}
	}

	// Print results
	benchmark.CompareResults(results)

	// Save results
	if err := benchmark.SaveResults(results, "benchmark_results"); err != nil {
		log.Printf("Failed to save results: %v", err)
	}
}
