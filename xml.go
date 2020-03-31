package openapi

// codebeat:disable[TOO_MANY_IVARS]

// XML Object
type XML struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

// Validate the values of XML object.
func (xml XML) Validate() error {
	return mustURL("xml.namespace", xml.Namespace)
}
