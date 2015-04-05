package apigenerator

type JSONSchema struct {
	name string
	content map[string]interface{}
}

func  NewJSONSchema(name string, content map[string]interface{}) *JSONSchema {
	c := &JSONSchema{name, content}
	return c
}

func (js *JSONSchema) Name() string {
	return js.name
}

func (js *JSONSchema) Content() map[string]interface{} {
	return js.content
}

type JSONSchemaExtended struct {
}
