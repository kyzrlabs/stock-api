package main

import (
	_ "embed"
	v1 "gitlab.com/eiseisbaby1/api/api/v1"
	"gitlab.com/eiseisbaby1/api/internal/data"
	resthttp "gitlab.com/eiseisbaby1/api/internal/http"
	"gitlab.com/eiseisbaby1/api/internal/rest"
	"gitlab.com/eiseisbaby1/api/pkg/resources"
	"log"
)

const ()

//go:embed data.json
var testfile []byte

func main() {

	reader, err := data.NewCatalogReader(testfile)
	if err != nil {
		log.Fatalf("Failed to create catalog reader: %v", err)
	}

	stockCatalogHandler := rest.NewHandler[v1.StockCatalog](resources.NewStockCatalogHandler(reader))
	categoryHandler := rest.NewHandler[v1.Category](resources.NewCategoryHandler(reader))
	stockItemHandler := rest.NewHandler[v1.Item](resources.NewStockItemHandler(reader))

	apiServer := resthttp.NewApiServer(8001)

	apiServer.Use(resthttp.MiddlewareRecovery)
	apiServer.Use(resthttp.MiddlewareCORS)

	apiServer.AddHandler("/stock-catalog", stockCatalogHandler.Get)
	apiServer.AddHandler("/stock-catalog/{id}", categoryHandler.Get)
	apiServer.AddHandler("/items/{id}", stockItemHandler.Get)

	apiServer.AddStaticHandler("/static/", "./static")

	apiServer.ListenAndServe()
}
