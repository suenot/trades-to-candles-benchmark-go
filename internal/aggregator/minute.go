package aggregator

import (
	"time"

	"github.com/suenot/trades-to-candles-benchmark-go/internal/models"
)

// MinuteAggregator агрегирует трейды в минутные свечи
type MinuteAggregator struct {
	trades []*models.Trade
}

// NewMinuteAggregator создает новый агрегатор
func NewMinuteAggregator() *MinuteAggregator {
	return &MinuteAggregator{
		trades: make([]*models.Trade, 0),
	}
}

// AddTrade добавляет трейд
func (a *MinuteAggregator) AddTrade(trade *models.Trade) error {
	if err := trade.Validate(); err != nil {
		return err
	}
	a.trades = append(a.trades, trade)
	return nil
}

// GetCandles возвращает минутные свечи
func (a *MinuteAggregator) GetCandles() ([]*models.Candle, error) {
	if len(a.trades) == 0 {
		return []*models.Candle{}, nil
	}

	// Группируем трейды по минутам
	candleMap := make(map[time.Time][]*models.Trade)

	for _, trade := range a.trades {
		// Округляем время до минуты
		candleTime := trade.Timestamp.Truncate(time.Minute)
		candleMap[candleTime] = append(candleMap[candleTime], trade)
	}

	// Создаем свечи из сгруппированных трейдов
	candles := make([]*models.Candle, 0, len(candleMap))

	for candleTime, trades := range candleMap {
		if len(trades) == 0 {
			continue
		}

		candle := &models.Candle{
			Timestamp: candleTime,
			Open:      trades[0].Price,
			High:      trades[0].Price,
			Low:       trades[0].Price,
			Close:     trades[len(trades)-1].Price,
			Volume:    0,
		}

		// Вычисляем High, Low и Volume
		for _, trade := range trades {
			if trade.Price > candle.High {
				candle.High = trade.Price
			}
			if trade.Price < candle.Low {
				candle.Low = trade.Price
			}
			candle.Volume += trade.Amount
		}

		candles = append(candles, candle)
	}

	return candles, nil
}
