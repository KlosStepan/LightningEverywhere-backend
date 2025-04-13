package eshop

import "context"

type Store interface {
	GetAll(ctx context.Context) ([]Eshop, error)
	Create(ctx context.Context, e *Eshop) error
	Update(ctx context.Context, id string, e *Eshop) error
	Delete(ctx context.Context, id string) error
}
