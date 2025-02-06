package benchmark

import (
	"time"
)

// BenchmarkResult stores the results of a single benchmark run
type BenchmarkResult struct {
	LibraryName     string
	ExecutionTime   time.Duration
	MemoryUsage     uint64
	ProcessedTrades int64
	Error           error
}

// Runner interface defines the benchmark runner behavior
type Runner interface {
	Run() (*BenchmarkResult, error)
}
