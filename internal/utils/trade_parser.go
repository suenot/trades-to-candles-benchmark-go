package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"

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
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if err := trade.Validate(); err != nil {
		return nil, fmt.Errorf("invalid trade data: %w", err)
	}

	return &trade, nil
}

// ParseAll reads all trades from the input
func (p *TradeParser) ParseAll() (chan *models.Trade, error) {
	trades := make(chan *models.Trade, 1000)

	go func() {
		defer close(trades)

		for {
			trade, err := p.ParseTrade()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error parsing trade: %v", err)
				continue
			}
			trades <- trade
		}
	}()

	return trades, nil
}
