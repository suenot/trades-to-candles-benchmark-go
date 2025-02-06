package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/utils"
)

func main() {
	dataFile := flag.String("data", "", "Path to trades data file")
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

	// Get file info
	fileName, fileSize, err := zipReader.GetFileInfo()
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}
	fmt.Printf("Processing file: %s (%.2f MB)\n", fileName, float64(fileSize)/float64(1024*1024))

	// Get trades data
	tradesFile, err := zipReader.GetFirstFile()
	if err != nil {
		log.Fatalf("Failed to get trades file: %v", err)
	}
	defer tradesFile.Close()

	// Create trade parser
	parser := utils.NewTradeParser(tradesFile)
	count, err := parser.ParseAll()
	if err != nil {
		log.Fatalf("Failed to parse trades: %v", err)
	}

	fmt.Printf("Total trades in file: %d\n", count)
}
