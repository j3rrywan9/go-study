package main

import (
	"fmt"
	"github.com/j3rrywan9/apigenerator"
	"github.com/j3rrywan9/engine"
)

const path string = "/api/json/delphix.json"

func main() {
	de := new(engine.Engine)
	de.SetAddress("10.43.3.159")
	generator := new(apigenerator.ApiGenerator)
	schemas := generator.LoadSchemasFromEngine(de, path)
	fmt.Println(schemas)
}
