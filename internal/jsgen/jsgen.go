// file: internal/jsgen/jsgen.go
package jsgen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

// GenerateJS generates JavaScript classes based on Go structs.
func GenerateJS(goStructFile string, outputFile string) error {
	// Open a file to write the generated JS classes
	jsFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create JS file: %v", err)
	}
	defer jsFile.Close()

	// Write a header to the JS file
	jsFile.WriteString("// Auto-generated JavaScript classes from Go structs\n\n")

	// Parse the Go file to extract struct info
	fset := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fset, goStructFile, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("failed to parse Go file: %v", err)
	}

	// Loop through the Go file and generate JavaScript classes for each struct
	for _, decl := range parsedFile.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Only process structs
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Write the JS class declaration
			className := typeSpec.Name.Name
			jsFile.WriteString(fmt.Sprintf("export class %s {\n", className))
			jsFile.WriteString("  constructor(")

			// Generate the constructor parameters and body using JSON tag names
			var params []string
			var constructorBody []string

			for _, field := range structType.Fields.List {
				// Check if there's a json tag and use it; otherwise, use the field name
				jsonTag := getJSONTag(field)
				if jsonTag == "" {
					continue // Skip unexported or non-JSON tagged fields
				}

				params = append(params, jsonTag)
				constructorBody = append(constructorBody, fmt.Sprintf("    this.%s = %s;", jsonTag, jsonTag))
			}

			// Write constructor parameters and body
			jsFile.WriteString(strings.Join(params, ", "))
			jsFile.WriteString(") {\n")
			jsFile.WriteString(strings.Join(constructorBody, "\n"))
			jsFile.WriteString("\n  }\n}\n\n")
		}
	}

	return nil
}

// getJSONTag returns the json tag for a field if it exists, otherwise returns the field name in lowercase.
func getJSONTag(field *ast.Field) string {
	if field.Tag != nil {
		tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`")).Get("json")
		if tag != "" && tag != "-" {
			return strings.Split(tag, ",")[0]
		}
	}
	if len(field.Names) > 0 {
		// Use lowercase field name if no JSON tag is provided
		return strings.ToLower(field.Names[0].Name)
	}
	return ""
}
