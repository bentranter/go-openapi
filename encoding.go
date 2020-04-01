package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Encoding Object
type Encoding struct {
	ContentType   string             `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string             `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool               `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool               `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

// Validate the values of Encoding object.
func (encoding Encoding) Validate() error {
	for _, header := range encoding.Headers {
		if err := header.Validate(); err != nil {
			return err
		}
	}
	return nil
}
