package transactions

import "time"

type Transaction struct {
	ID        int       `json:"ID,omitempty"`
	UserID    int       `json:"UserID,omitempty"`
	Value     float64   `json:"Value,omitempty"`
	Type      int       `json:"Type"`
	TimeStamp time.Time `json:"TimeStamp"`
}
