package openapi

import "errors"

// codebeat:disable[TOO_MANY_IVARS]

// Link Object
type Link struct {
	OperationRef string                 `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationID  string                 `yaml:"operationId,omitempty" json:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Server       *Server                `json:"server,omitempty" yaml:"server,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
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
