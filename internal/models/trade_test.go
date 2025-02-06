package models

import (
	"testing"
	"time"
)

func TestTradeValidation(t *testing.T) {
	tests := []struct {
		name    string
		trade   Trade
		wantErr bool
	}{
		{
			name: "valid trade",
			trade: Trade{
				Timestamp: time.Now(),
				Price:     100.0,
				Amount:    1.0,
				IsBuyer:   true,
			},
			wantErr: false,
		},
		{
			name: "zero timestamp",
			trade: Trade{
				Timestamp: time.Time{},
				Price:     100.0,
				Amount:    1.0,
				IsBuyer:   true,
			},
			wantErr: true,
		},
		{
			name: "zero price",
			trade: Trade{
				Timestamp: time.Now(),
				Price:     0.0,
				Amount:    1.0,
				IsBuyer:   true,
			},
			wantErr: true,
		},
		{
			name: "negative price",
			trade: Trade{
				Timestamp: time.Now(),
				Price:     -100.0,
				Amount:    1.0,
				IsBuyer:   true,
			},
			wantErr: true,
		},
		{
			name: "zero amount",
			trade: Trade{
				Timestamp: time.Now(),
				Price:     100.0,
				Amount:    0.0,
				IsBuyer:   true,
			},
			wantErr: true,
		},
		{
			name: "negative amount",
			trade: Trade{
				Timestamp: time.Now(),
				Price:     100.0,
				Amount:    -1.0,
				IsBuyer:   true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.trade.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Trade.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
