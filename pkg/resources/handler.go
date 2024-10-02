package resources

import (
	"context"
)

type Handler[T any] interface {
	List(ctx context.Context) []T
	Get(ctx context.Context, id string) (*T, error)
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, item *T) (*T, error)
	Update(ctx context.Context, oldItem *T, newItem *T) (*T, error)
	Name() string
}
