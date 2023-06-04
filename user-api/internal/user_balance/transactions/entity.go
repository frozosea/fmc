package transactions

import "user-api/internal/domain"

type Transaction struct {
	*domain.Transaction `json:"transactions"`
	Number              string  `json:"number"`
	IsContainer         bool    `json:"isContainer"`
	UserId              int64   `json:"userId"`
	MoneySpent          float64 `json:"moneySpent"`
	DaysOnTracking      int     `json:"daysOnTracking"`
}
