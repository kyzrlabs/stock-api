package resources

import (
	"context"
	v1 "gitlab.com/eiseisbaby1/api/api/v1"
)

type CategoryGetter interface {
	GetCategory(name string) *v1.Category
}

type categoryHandler struct {
	getter CategoryGetter
}

func NewCategoryHandler(getter CategoryGetter) Handler[v1.Category] {
	return categoryHandler{
		getter: getter,
	}
}

func (c categoryHandler) List(ctx context.Context) []v1.Category {
	//TODO implement me
	panic("implement me")
}

func (c categoryHandler) Get(ctx context.Context, name string) (*v1.Category, error) {
	return c.getter.GetCategory(name), nil
}

func (c categoryHandler) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (c categoryHandler) Create(ctx context.Context, item *v1.Category) (*v1.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryHandler) Update(ctx context.Context, oldItem *v1.Category, newItem *v1.Category) (*v1.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryHandler) Name() string {
	//TODO implement me
	panic("implement me")
}
