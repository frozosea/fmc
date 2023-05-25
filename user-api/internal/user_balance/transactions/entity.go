package transactions

import "user-api/internal/domain"

type Transaction struct {
	*domain.Transaction
	Number string
}
