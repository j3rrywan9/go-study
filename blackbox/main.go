package main

import (
	"encoding/json"
	"fmt"
	"github.com/j3rrywan9/apigenerator"
	"github.com/j3rrywan9/engine"
)

const path string = "/api/json/delphix.json"

func main() {
	de := new(engine.Engine)
	de.SetAddress("10.43.3.159")
	generator := new(apigenerator.ApiGenerator)
	schemas_string := generator.LoadSchemasFromEngine(de, path)

	var f map[string]map[string]interface{}
	err := json.Unmarshal([]byte(schemas_string), &f)
	if err != nil {
		panic(err)
	}
	
	var jsonschemas []apigenerator.JSONSchema
	for json_schema_name, json_schema_content := range f {
		jsonschema := apigenerator.NewJSONSchema(json_schema_name, json_schema_content)
		jsonschemas = append(jsonschemas, *jsonschema)
	}
	for i := 0; i < len(jsonschemas); i++ {
		fmt.Println(jsonschemas[i].Name())
		fmt.Println(jsonschemas[i].Content())
	}
}
