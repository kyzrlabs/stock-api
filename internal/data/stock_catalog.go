package data

import (
	"encoding/json"
	v1 "gitlab.com/eiseisbaby1/api/api/v1"
	"gitlab.com/eiseisbaby1/api/pkg/util"
)

type StockCatalogReader struct {
	data *v1.StockCatalog
}

func NewCatalogReader(data []byte) (*StockCatalogReader, error) {

	var stockCatalog v1.StockCatalog

	err := json.Unmarshal(data, &stockCatalog.Categories)
	if err != nil {
		return nil, err // Return the error if parsing fails
	}

	return &StockCatalogReader{
		data: &stockCatalog,
	}, nil
}

func (s *StockCatalogReader) GetStockCatalog() *v1.StockCatalog {
	return s.data
}

func (s *StockCatalogReader) GetCategory(name string) *v1.Category {
	name = util.UpperFirst(name)
	return util.ToPtr(s.data.Categories[name])
}

func (s *StockCatalogReader) GetCategories() []v1.Category {
	var categories []v1.Category
	for _, category := range s.data.Categories {
		categories = append(categories, category)
	}
	return categories
}

func (s *StockCatalogReader) GetStockItem(id string) *v1.Item {
	for _, category := range s.data.Categories {
		for _, item := range category.Items {
			if item.ID == id {
				return &item
			}
		}
	}
	return nil
}
