package openapi

import (
	"fmt"
	"strings"
)

// codebeat:disable[TOO_MANY_IVARS]

// Schema Object
type Schema struct {
	Title            string   `json:"title,omitempty" yaml:"title,omitempty"`
	MultipleOf       int      `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Maximum          int      `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool     `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          int      `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool     `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        int      `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        int      `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string   `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         int      `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         int      `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	MaxProperties    int      `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties    int      `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required         []string `json:"required,omitempty" yaml:"required,omitempty"`
	Enum             []string `json:"enum,omitempty" yaml:"enum,omitempty"`

	Type                 string             `json:"type,omitempty" yaml:"type,omitempty"`
	AllOf                []*Schema          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []*Schema          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []*Schema          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not                  *Schema            `json:"not,omitempty" yaml:"not,omitempty"`
	Items                *Schema            `json:"items,omitempty" yaml:"items,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties *Schema            `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Description          string             `json:"description,omitempty" yaml:"description,omitempty"`
	Format               string             `json:"format,omitempty" yaml:"format,omitempty"`
	Default              string             `json:"default,omitempty" yaml:"default,omitempty"`

	Nullable      bool                   `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Discriminator *Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ReadOnly      bool                   `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly     bool                   `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	XML           *XML                   `json:"xml,omitempty" yaml:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example       interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Deprecated    bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	Extension map[string]interface{} `json:"extension,omitempty" yaml:",inline"`
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
