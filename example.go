package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Example Object
type Example struct {
	Summary       string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{} `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue interface{} `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
