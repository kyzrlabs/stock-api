//go:build js && wasm

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall/js"
)

const apiURL = "http://localhost:8001"

func fetchStockCatalog(this js.Value, p []js.Value) interface{} {
	return answer("stock-catalog")
}

func fetchCategory(this js.Value, p []js.Value) interface{} {
	return answer(fmt.Sprintf("stock-catalog/%s", p[0].String()))
}

func fetchItem(this js.Value, p []js.Value) interface{} {
	return answer(fmt.Sprintf("items/%s", p[0].String()))
}

func answer(path string) interface{} {
	// Create a channel to pass data back to JavaScript asynchronously
	resultChan := make(chan string)

	go func() {
		// Fetch data from the API
		resp, err := http.Get(fmt.Sprintf("%s/%s", apiURL, path))
		if err != nil {
			resultChan <- fmt.Sprintf("Error: %v", err)
			close(resultChan)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resultChan <- fmt.Sprintf("Error reading response: %v", err)
			close(resultChan)
			return
		}

		// Pass the result back to JavaScript
		resultChan <- string(body)
		close(resultChan)
	}()

	// Return the result via the channel
	promise := js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		thenCallback := args[0]
		go func() {
			result := <-resultChan
			thenCallback.Invoke(js.ValueOf(result))
		}()
		return nil
	}))

	return promise
}

func main() {
	// Expose Go function to JavaScript
	js.Global().Set("fetchStockCatalog", js.FuncOf(fetchStockCatalog))
	js.Global().Set("fetchCategory", js.FuncOf(fetchCategory))
	js.Global().Set("fetchItem", js.FuncOf(fetchItem))

	// Prevent the Go program from exiting immediately
	select {}
}
