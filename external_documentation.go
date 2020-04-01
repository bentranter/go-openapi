package openapi

// codebeat:disable[TOO_MANY_IVARS]

// ExternalDocumentation Object
type ExternalDocumentation struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Validate the values of ExternalDocumentaion object.
func (externalDocumentation ExternalDocumentation) Validate() error {
	return mustURL("externalDocumentation.url", externalDocumentation.URL)
}
