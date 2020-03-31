package openapi

import (
	"fmt"
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Schema Object
type Schema struct {
	Title            string   `json:"title,omitempty"`
	MultipleOf       int      `yaml:"multipleOf" json:"multipleOf,omitempty"`
	Maximum          int      `json:"maximum,omitempty"`
	ExclusiveMaximum bool     `yaml:"exclusiveMaximum" json:"exclusiveMaximum,omitempty"`
	Minimum          int      `json:"minimum,omitempty"`
	ExclusiveMinimum bool     `yaml:"exclusiveMinimum" json:"exclusiveMinimum,omitempty"`
	MaxLength        int      `yaml:"maxLength" json:"maxLength,omitempty"`
	MinLength        int      `yaml:"minLength" json:"minLength,omitempty"`
	Pattern          string   `json:"pattern,omitempty"`
	MaxItems         int      `yaml:"maxItems" json:"maxItems,omitempty"`
	MinItems         int      `yaml:"minItems" json:"minItems,omitempty"`
	MaxProperties    int      `yaml:"maxProperties" json:"maxProperties,omitempty"`
	MinProperties    int      `yaml:"minProperties" json:"minProperties,omitempty"`
	Required         []string `json:"required,omitempty"`
	Enum             []string `json:"enum,omitempty"`

	Type                 string             `json:"type,omitempty"`
	AllOf                []*Schema          `yaml:"allOf" json:"allOf,omitempty"`
	OneOf                []*Schema          `yaml:"oneOf" json:"oneOf,omitempty"`
	AnyOf                []*Schema          `yaml:"anyOf" json:"anyOf,omitempty"`
	Not                  *Schema            `json:"not,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema            `yaml:"additionalProperties" json:"additionalProperties,omitempty"`
	Description          string             `json:"description,omitempty"`
	Format               string             `json:"format,omitempty"`
	Default              string             `json:"default,omitempty"`

	Nullable      bool                   `json:"nullable,omitempty"`
	Discriminator *Discriminator         `json:"discriminator,omitempty"`
	ReadOnly      bool                   `yaml:"readOnly" json:"readOnly,omitempty"`
	WriteOnly     bool                   `yaml:"writeOnly" json:"writeOnly,omitempty"`
	XML           *XML                   `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `yaml:"externalDocs" json:"externalDocs,omitempty"`
	Example       interface{}            `json:"example,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty"`

	Ref string `yaml:"$ref" json:"$ref,omitempty"`

	Extension map[string]interface{} `yaml:",inline" json:"extension,omitempty"`
}

// Validate the values of Schema object.
func (schema Schema) Validate() error {
	validaters := []validater{}
	for _, s := range schema.AllOf {
		validaters = append(validaters, s)
	}
	for _, s := range schema.OneOf {
		validaters = append(validaters, s)
	}
	for _, s := range schema.AnyOf {
		validaters = append(validaters, s)
	}
	if schema.Not != nil {
		validaters = append(validaters, schema.Not)
	}
	if schema.Items != nil {
		validaters = append(validaters, schema.Items)
	}
	if schema.Discriminator != nil {
		validaters = append(validaters, schema.Discriminator)
	}
	if schema.XML != nil {
		validaters = append(validaters, schema.XML)
	}
	if schema.ExternalDocs != nil {
		validaters = append(validaters, schema.ExternalDocs)
	}
	for _, property := range schema.Properties {
		validaters = append(validaters, property)
	}
	if e, ok := schema.Example.(validater); ok {
		validaters = append(validaters, e)
	}
	for k := range schema.Extension {
		if !strings.HasPrefix(k, "x-") {
			return fmt.Errorf("unknown field: %s", k)
		}
	}
	return validateAll(validaters)
}
