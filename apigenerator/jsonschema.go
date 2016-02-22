package apigenerator

import (
	"encoding/json"
	//"fmt"
	//"reflect"
)

type JSONSchemaFile struct {
	name string
	content JSONSchemaExtended
}

type JSONSchemaExtended struct {
	name string
	title map[string]interface{}
	description string
	required bool
	enum []string 
	minimum uint16
	maximum uint16
	exclusive_minimum uint16
	exclusive_maximum uint16
	divisible_by uint16
	default_value string
	ref string
	pattern string
	max_length uint16
	min_length uint16
	format string
	items map[string]string
	value_type string
	properties map[string]map[string]string
	root string
	extends map[string]string
	create map[string]interface{}
	read map[string]interface{}
	list map[string]interface{}
	update map[string]interface{}
	delete map[string]interface{}
	singleton bool
	referenceTo string
	operations map[string]interface{}
	rootOperations map[string]interface{}
}

type JSONSchema struct {
	name string
}

func (this *JSONSchemaFile) Name() string {
	return this.name
}

func (this *JSONSchemaFile) SetName(name string) {
	this.name = name
}

func (this *JSONSchemaFile) Content() JSONSchemaExtended {
	return this.content
}

func (this *JSONSchemaFile) SetContent(content JSONSchemaExtended) {
	this.content = content
}

func (this *JSONSchemaExtended) Name() string {
	return this.name
}

func (this *JSONSchemaExtended) Title() map[string]interface{} {
	return this.title
}

func (this *JSONSchemaExtended) Description() string {
	return this.description
}

func (this *JSONSchemaExtended) UnmarshalJSON(b []byte) error {
	var f map[string]interface{}
	type TitleType map[string]string
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	this.name = f["name"].(string)
	this.title = f["title"].(map[string]interface{})
	this.description = f["description"].(string)
	return nil
}
