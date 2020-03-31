package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Encoding Object
type Encoding struct {
	ContentType   string             `yaml:"contentType" json:"contentType,omitempty"`
	Headers       map[string]*Header `json:"headers,omitempty"`
	Style         string             `json:"style,omitempty"`
	Explode       bool               `json:"explode,omitempty"`
	AllowReserved bool               `yaml:"allowReserved" json:"allowReserved,omitempty"`
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
