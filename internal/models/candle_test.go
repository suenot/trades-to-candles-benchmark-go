package models

import (
	"testing"
	"time"
)

func TestCandleValidation(t *testing.T) {
	tests := []struct {
		name    string
		candle  Candle
		wantErr bool
	}{
		{
			name: "valid candle",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      100.0,
				High:      110.0,
				Low:       90.0,
				Close:     105.0,
				Volume:    10.0,
			},
			wantErr: false,
		},
		{
			name: "zero timestamp",
			candle: Candle{
				Timestamp: time.Time{},
				Open:      100.0,
				High:      110.0,
				Low:       90.0,
				Close:     105.0,
				Volume:    10.0,
			},
			wantErr: true,
		},
		{
			name: "high not highest",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      100.0,
				High:      90.0,
				Low:       80.0,
				Close:     105.0,
				Volume:    10.0,
			},
			wantErr: true,
		},
		{
			name: "low not lowest",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      100.0,
				High:      110.0,
				Low:       105.0,
				Close:     105.0,
				Volume:    10.0,
			},
			wantErr: true,
		},
		{
			name: "negative volume",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      100.0,
				High:      110.0,
				Low:       90.0,
				Close:     105.0,
				Volume:    -1.0,
			},
			wantErr: true,
		},
		{
			name: "zero open price",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      0.0,
				High:      110.0,
				Low:       90.0,
				Close:     105.0,
				Volume:    10.0,
			},
			wantErr: true,
		},
		{
			name: "zero close price",
			candle: Candle{
				Timestamp: time.Now(),
				Open:      100.0,
				High:      110.0,
				Low:       90.0,
				Close:     0.0,
				Volume:    10.0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.candle.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Candle.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
