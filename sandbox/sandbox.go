package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"github.com/j3rrywan9/apigenerator"
)

func main() {
	//data, err := ioutil.ReadFile("schemas.txt")
	data, err := ioutil.ReadFile("schemas2.txt")
	if err != nil {
		panic(err)
	}

	var f map[string]apigenerator.JSONSchemaExtended

	if err := json.Unmarshal(data, &f); err != nil {
		panic(err)
	}
	
	fmt.Println("The type of f is:", reflect.TypeOf(f))

	var schemafiles []apigenerator.JSONSchemaFile

	for file_name, file_content := range f {
		schemafile := new(apigenerator.JSONSchemaFile)
		schemafile.SetName(file_name)
		schemafile.SetContent(file_content)
		schemafiles = append(schemafiles, *schemafile)
	}

	for i := 0; i < len(schemafiles); i++ {
		fmt.Println(schemafiles[i].Name())
		schema := new(apigenerator.JSONSchemaExtended)
		*schema = schemafiles[i].Content()
		fmt.Println(schema.Name())
		fmt.Println(schema.Title())
		fmt.Println(schema.Description())
		fmt.Println(reflect.TypeOf(schema.Title()))
	}
}
