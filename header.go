package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Header Object
type Header struct {
	Description     string `json:"description" yaml:"description"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	Style         string              `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema        *Schema             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example       interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	Examples      map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Validate the values of Header object.
func (header Header) Validate() error {
	validaters := []validater{}
	if header.Schema != nil {
		validaters = append(validaters, header.Schema)
	}
	if v, ok := header.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	if len(header.Content) > 1 {
		return ErrTooManyHeaderContent
	}
	for _, mediaType := range header.Content {
		validaters = append(validaters, mediaType)
	}
	return validateAll(validaters)
}
