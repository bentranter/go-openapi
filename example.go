package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Example Object
type Example struct {
	Summary       string      `json:"summary,omitempty"`
	Description   string      `json:"description,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ExternalValue interface{} `yaml:"externalValue" json:"externalValue,omitempty"`

	Ref string `yaml:"$ref" json:"$ref,omitempty"`
}
