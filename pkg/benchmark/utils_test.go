package benchmark

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSaveResults(t *testing.T) {
	// Create temporary directory for test
	tmpDir, err := os.MkdirTemp("", "benchmark_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test results
	results := []*BenchmarkResult{
		{
			LibraryName:     "test_lib",
			ExecutionTime:   time.Second,
			MemoryUsage:     1024 * 1024, // 1MB
			ProcessedTrades: 1000,
			CandlesCreated:  10,
		},
	}

	// Save results
	if err := SaveResults(results, tmpDir); err != nil {
		t.Fatalf("Failed to save results: %v", err)
	}

	// Verify file was created
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read temp dir: %v", err)
	}

	if len(files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(files))
	}

	// Verify file content
	filePath := filepath.Join(tmpDir, files[0].Name())
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read result file: %v", err)
	}

	if len(content) == 0 {
		t.Error("Result file is empty")
	}
}
