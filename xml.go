package openapi

// codebeat:disable[TOO_MANY_IVARS]

// XML Object
type XML struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

// Validate the values of XML object.
func (xml XML) Validate() error {
	return mustURL("xml.namespace", xml.Namespace)
}
