package benchmark

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// SaveResults saves benchmark results to a JSON file
func SaveResults(results []*BenchmarkResult, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	filename := fmt.Sprintf("benchmark_results_%s.json",
		time.Now().Format("2006-01-02_15-04-05"))

	filepath := filepath.Join(outputDir, filename)

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create results file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results: %w", err)
	}

	return nil
}

// CompareResults prints a comparison of benchmark results
func CompareResults(results []*BenchmarkResult) {
	fmt.Println("\nBenchmark Results Comparison:")
	fmt.Println("=============================")

	for _, r := range results {
		fmt.Printf("\nLibrary: %s\n", r.LibraryName)
		fmt.Printf("Execution Time: %v\n", r.ExecutionTime)
		fmt.Printf("Max Memory Usage: %v MB\n", r.MaxMemoryUsage/1024/1024)
		fmt.Printf("Processed Trades: %d\n", r.ProcessedTrades)
		fmt.Printf("Candles Created: %d\n", r.CandlesCreated)
		if r.Error != nil {
			fmt.Printf("Error: %v\n", r.Error)
		}
	}
}
