package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Parse command line flags
	dataFile := flag.String("data", "", "Path to trades data file")
	flag.Parse()

	if *dataFile == "" {
		log.Fatal("Please provide path to data file using -data flag")
	}

	fmt.Println("Starting benchmark with data file:", *dataFile)
	// TODO: Implement benchmark runner
}
