package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Response Object
type Response struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*Link      `json:"links,omitempty" yaml:"links,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Validate the value of Response object.
func (response Response) Validate() error {
	if response.Ref != "" {
		return nil // validated in doc.Components
	}
	if response.Description == "" {
		return ErrRequired{Target: "response.description"}
	}
	validaters := []validater{}
	for _, header := range response.Headers {
		validaters = append(validaters, header)
	}
	for _, mediaType := range response.Content {
		validaters = append(validaters, mediaType)
	}
	for _, link := range response.Links {
		validaters = append(validaters, link)
	}
	return validateAll(validaters)
}
