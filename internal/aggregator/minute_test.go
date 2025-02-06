package aggregator

import (
	"testing"
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

func TestMinuteAggregator(t *testing.T) {
	// Создаем агрегатор
	aggregator := NewMinuteAggregator()

	// Тестовые данные: 3 трейда в одной минуте
	now := time.Now().Truncate(time.Minute)
	trades := []*models.Trade{
		{
			Timestamp: now,
			Price:     100.0,
			Amount:    1.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(20 * time.Second),
			Price:     102.0,
			Amount:    2.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(40 * time.Second),
			Price:     98.0,
			Amount:    1.5,
			IsBuyer:   false,
		},
	}

	// Добавляем трейды
	for _, trade := range trades {
		if err := aggregator.AddTrade(trade); err != nil {
			t.Errorf("Failed to add trade: %v", err)
		}
	}

	// Получаем свечи
	candles, err := aggregator.GetCandles()
	if err != nil {
		t.Fatalf("Failed to get candles: %v", err)
	}

	// Проверяем количество свечей
	if len(candles) != 1 {
		t.Fatalf("Expected 1 candle, got %d", len(candles))
	}

	candle := candles[0]

	// Проверяем данные свечи
	if !candle.Timestamp.Equal(now) {
		t.Errorf("Expected timestamp %v, got %v", now, candle.Timestamp)
	}
	if candle.Open != 100.0 {
		t.Errorf("Expected open price 100.0, got %v", candle.Open)
	}
	if candle.High != 102.0 {
		t.Errorf("Expected high price 102.0, got %v", candle.High)
	}
	if candle.Low != 98.0 {
		t.Errorf("Expected low price 98.0, got %v", candle.Low)
	}
	if candle.Close != 98.0 {
		t.Errorf("Expected close price 98.0, got %v", candle.Close)
	}
	if candle.Volume != 4.5 {
		t.Errorf("Expected volume 4.5, got %v", candle.Volume)
	}
}

func TestMinuteAggregatorMultipleMinutes(t *testing.T) {
	aggregator := NewMinuteAggregator()

	// Тестовые данные: трейды в разных минутах
	now := time.Now().Truncate(time.Minute)
	trades := []*models.Trade{
		{
			Timestamp: now,
			Price:     100.0,
			Amount:    1.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(time.Minute),
			Price:     101.0,
			Amount:    2.0,
			IsBuyer:   true,
		},
		{
			Timestamp: now.Add(2 * time.Minute),
			Price:     102.0,
			Amount:    1.5,
			IsBuyer:   false,
		},
	}

	// Добавляем трейды
	for _, trade := range trades {
		if err := aggregator.AddTrade(trade); err != nil {
			t.Errorf("Failed to add trade: %v", err)
		}
	}

	// Получаем свечи
	candles, err := aggregator.GetCandles()
	if err != nil {
		t.Fatalf("Failed to get candles: %v", err)
	}

	// Проверяем количество свечей
	if len(candles) != 3 {
		t.Fatalf("Expected 3 candles, got %d", len(candles))
	}

	// Проверяем каждую свечу
	expectedPrices := []float64{100.0, 101.0, 102.0}
	expectedVolumes := []float64{1.0, 2.0, 1.5}

	for i, candle := range candles {
		if candle.Open != expectedPrices[i] {
			t.Errorf("Candle %d: Expected open price %v, got %v", i, expectedPrices[i], candle.Open)
		}
		if candle.Close != expectedPrices[i] {
			t.Errorf("Candle %d: Expected close price %v, got %v", i, expectedPrices[i], candle.Close)
		}
		if candle.Volume != expectedVolumes[i] {
			t.Errorf("Candle %d: Expected volume %v, got %v", i, expectedVolumes[i], candle.Volume)
		}
	}
}
