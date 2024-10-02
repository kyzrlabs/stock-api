// file: /cmd/generate_entities.go
package main

import (
	"flag"
	"fmt"
	"gitlab.com/eiseisbaby1/api/internal/jsgen"
	"log"
)

func main() {
	// Define flags for input and output files
	inputFile := flag.String("input", "", "Input Go file with structs")
	outputFile := flag.String("output", "", "Output JS file")
	flag.Parse()

	// Check if both input and output flags are set
	if *inputFile == "" || *outputFile == "" {
		log.Fatalf("Error: input and output files must be specified.")
	}

	// Call the GenerateJS function to generate the JS entities
	err := jsgen.GenerateJS(*inputFile, *outputFile)
	if err != nil {
		log.Fatalf("Error generating JS: %v", err)
	}

	fmt.Println("JavaScript entities successfully generated.")
}
