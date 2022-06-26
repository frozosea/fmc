package city

import (
	"context"
)

type IController interface {
	AddCity(ctx context.Context, city BaseCity) error
	GetAll(ctx context.Context) ([]*City, error)
}
