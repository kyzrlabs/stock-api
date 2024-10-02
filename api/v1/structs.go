package v1

type Category struct {
	Items []Item `json:"items" bson:"items"`
}

type Contents struct {
	Water     float64 `json:"water" bson:"water"`
	Sugar     float64 `json:"sugar" bson:"sugar"`
	Fat       float64 `json:"fat" bson:"fat"`
	DryMatter float64 `json:"dry_matter" bson:"dry_matter"`
}

type Item struct {
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Contents     Contents `json:"contents" bson:"contents"`
	Calories100g float64  `json:"calories_100g" bson:"calories_100g"`
}

type StockCatalog struct {
	Categories map[string]Category `json:"categories" bson:"categories"`
}
