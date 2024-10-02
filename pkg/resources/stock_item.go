package resources

import (
	"context"
	v1 "gitlab.com/eiseisbaby1/api/api/v1"
)

type stockItemHandler struct {
	getter StockItemGetter
}

type StockItemGetter interface {
	GetStockItem(id string) *v1.Item
}

func NewStockItemHandler(getter StockItemGetter) Handler[v1.Item] {
	return stockItemHandler{
		getter: getter,
	}
}

func (s stockItemHandler) List(ctx context.Context) []v1.Item {
	panic("implement me")
}

func (s stockItemHandler) Get(ctx context.Context, id string) (*v1.Item, error) {
	return s.getter.GetStockItem(id), nil
}

func (s stockItemHandler) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (s stockItemHandler) Create(ctx context.Context, item *v1.Item) (*v1.Item, error) {
	panic("implement me")
}

func (s stockItemHandler) Update(ctx context.Context, oldItem *v1.Item, newItem *v1.Item) (*v1.Item, error) {
	panic("implement me")
}

func (s stockItemHandler) Name() string {
	panic("implement me")
}
