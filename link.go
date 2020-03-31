package openapi

import "errors"

// codebeat:disable[TOO_MANY_IVARS]

// Link Object
type Link struct {
	OperationRef string                 `yaml:"operationRef" json:"operationRef,omitempty"`
	OperationID  string                 `yaml:"operationId" json:"operationID,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	RequestBody  interface{}            `yaml:"requestBody" json:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Server       *Server                `json:"server,omitempty"`

	Ref string `yaml:"$ref" json:"$ref,omitempty"`
}

// Validate the values of Link object.
func (link Link) Validate() error {
	if link.OperationRef != "" && link.OperationID != "" {
		return errors.New("operationRef and operationId are mutually exclusive")
	}
	validaters := []validater{}
	for _, i := range link.Parameters {
		if v, ok := i.(validater); ok {
			validaters = append(validaters, v)
		}
	}
	if v, ok := link.RequestBody.(validater); ok {
		validaters = append(validaters, v)
	}
	if link.Server != nil {
		validaters = append(validaters, link.Server)
	}
	return validateAll(validaters)
}
