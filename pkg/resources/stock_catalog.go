package resources

import (
	"context"
	v1 "gitlab.com/eiseisbaby1/api/api/v1"
)

type stockCatalogHandler struct {
	getter StockCatalogGetter
}

type StockCatalogGetter interface {
	GetStockCatalog() *v1.StockCatalog
}

func NewStockCatalogHandler(getter StockCatalogGetter) Handler[v1.StockCatalog] {
	return stockCatalogHandler{
		getter: getter,
	}
}

func (s stockCatalogHandler) List(ctx context.Context) []v1.StockCatalog {
	//TODO implement me
	panic("implement me")
}

func (s stockCatalogHandler) Get(ctx context.Context, id string) (*v1.StockCatalog, error) {
	return s.getter.GetStockCatalog(), nil
}

func (s stockCatalogHandler) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s stockCatalogHandler) Create(ctx context.Context, item *v1.StockCatalog) (*v1.StockCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (s stockCatalogHandler) Update(ctx context.Context, oldItem *v1.StockCatalog, newItem *v1.StockCatalog) (*v1.StockCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (s stockCatalogHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
