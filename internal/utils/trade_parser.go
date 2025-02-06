package utils

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/yourusername/trades-to-candles-benchmark-go/internal/models"
)

// TradeParser handles parsing of trade data from various formats
type TradeParser struct {
	reader *bufio.Reader
}

// NewTradeParser creates a new TradeParser instance
func NewTradeParser(r io.Reader) *TradeParser {
	return &TradeParser{
		reader: bufio.NewReader(r),
	}
}

// ParseTrade reads and parses a single trade from the input
func (p *TradeParser) ParseTrade() (*models.Trade, error) {
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var trade models.Trade
	if err := json.Unmarshal([]byte(line), &trade); err != nil {
		return nil, err
	}

	return &trade, nil
}

// ParseAll reads all trades from the input
func (p *TradeParser) ParseAll() (chan *models.Trade, error) {
	trades := make(chan *models.Trade, 1000) // Buffer size can be adjusted

	go func() {
		defer close(trades)

		for {
			trade, err := p.ParseTrade()
			if err == io.EOF {
				break
			}
			if err != nil {
				// Log error but continue processing
				continue
			}
			trades <- trade
		}
	}()

	return trades, nil
}
