package datasource

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
)

type ItemsWrapper struct {
	Items []IdNameItem `json:"items"`
}

type Contents struct {
	Water     float64 `json:"water"`
	Sugar     float64 `json:"sugar"`
	Fat       float64 `json:"fat"`
	DryMatter float64 `json:"dry_matter"`
}

type IdNameItem struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Contents     Contents  `json:"contents"`
	Calories100g float64   `json:"calories_100g"`
}

type Category struct {
	Items []IdNameItem `json:"items"`
}

type FileData struct {
	Dairy      Category `json:"Dairy"`
	Sugar      Category `json:"Sugar"`
	Fruit      Category `json:"Fruit"`
	Syrup      Category `json:"Syrup"`
	Vegetables Category `json:"Vegetables"`
	Herbs      Category `json:"Herbs"`
	Alcohol    Category `json:"Alcohol"`
	Candy      Category `json:"Candy"`
}

type FileDS struct {
	data FileData
}

func FileDatasource(bytes []byte) FileDS {
	var data FileData
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("cannot unmarshal %v", err)
	}

	return FileDS{data: data}
}

func (f FileDS) Create(key string, value json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

func (f FileDS) Read(id uuid.UUID) (json.RawMessage, error) {
	// Loop through categories and items to find the item by ID
	categories := []Category{
		f.data.Dairy,
		f.data.Sugar,
		f.data.Fruit,
		f.data.Syrup,
		f.data.Vegetables,
		f.data.Herbs,
		f.data.Alcohol,
		f.data.Candy,
	}
	for _, category := range categories {
		for _, item := range category.Items {
			if item.ID == id {
				itemData, err := json.Marshal(item)
				if err != nil {
					return nil, err
				}
				return itemData, nil
			}
		}
	}
	return nil, nil
}

func (f FileDS) Update(key string, value json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

func (f FileDS) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (f FileDS) List() (json.RawMessage, error) {
	data, err := json.Marshal(f.data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
