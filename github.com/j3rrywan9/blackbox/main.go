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
	fmt.Println(schemas_string)
	var f interface{}
	err := json.Unmarshal([]byte(schemas_string), &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}	
}
