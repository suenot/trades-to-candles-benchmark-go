package models

import (
	"fmt"
	"time"
)

// Trade represents a single trade record
type Trade struct {
	Timestamp time.Time
	Price     float64
	Amount    float64
	IsBuyer   bool
}

// Validate checks if the trade data is valid
func (t *Trade) Validate() error {
	if t.Timestamp.IsZero() {
		return fmt.Errorf("timestamp is required")
	}
	if t.Price <= 0 {
		return fmt.Errorf("price must be positive, got: %v", t.Price)
	}
	if t.Amount <= 0 {
		return fmt.Errorf("amount must be positive, got: %v", t.Amount)
	}
	return nil
}

// String returns a string representation of the trade
func (t *Trade) String() string {
	return fmt.Sprintf("Trade{Time: %v, Price: %v, Amount: %v, IsBuyer: %v}",
		t.Timestamp, t.Price, t.Amount, t.IsBuyer)
}
