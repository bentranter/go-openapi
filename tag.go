package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Tag Object
type Tag struct {
	Name         string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// Validate the values of Tag object.
func (tag Tag) Validate() error {
	if tag.Name == "" {
		return ErrRequired{Target: "tag.name"}
	}
	if tag.ExternalDocs != nil {
		return tag.ExternalDocs.Validate()
	}
	return nil
}
