package apigenerator

type JSONSchemaFile struct {
	name string
	content map[string]interface{}
}

type JSONSchema struct {
	name string
	title string
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
}

func (this *JSONSchemaFile) Name() string {
	return this.name
}

func (this *JSONSchemaFile) SetName(name string) {
	this.name = name
}

func (this *JSONSchemaFile) Content() map[string]interface{} {
	return this.content
}

func (this *JSONSchemaFile) SetContent(content map[string]interface{}) {
	this.content = content
}

type JSONSchemaExtended struct {
}
