package models

import "time"

// Trade represents a single trade record
type Trade struct {
	Timestamp time.Time
	Price     float64
	Amount    float64
	IsBuyer   bool
}
