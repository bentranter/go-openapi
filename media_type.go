package openapi

// codebeat:disable[TOO_MANY_IVARS]

// MediaType Object
type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  interface{}          `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*Example  `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Validate the values of MediaType object.
// This function DOES NOT check whether the encoding object is in schema or not.
func (mediaType MediaType) Validate() error {
	validaters := []validater{}
	if mediaType.Schema != nil {
		validaters = append(validaters, mediaType.Schema)
	}
	if v, ok := mediaType.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	for _, e := range mediaType.Encoding {
		validaters = append(validaters, e)
	}
	return validateAll(validaters)
}
