package openapi

import (
	"strconv"
)

// codebeat:disable[TOO_MANY_IVARS]

// Operation Object
type Operation struct {
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationID,omitempty" yaml:"operationID,omitempty"`
	Parameters   []*Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    Responses              `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    map[string]*Callback   `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// SuccessResponse returns a success response object.
// If there are 2 or more success responses (like created and ok),
// it's not sure which is returned.
// If only match the default response or 2XX response, returned status code will be 0.
func (operation *Operation) SuccessResponse() (*Response, int, bool) {
	if operation == nil || operation.Responses == nil {
		return nil, -1, false
	}
	var defaultResponse *Response
	for statusStr, resp := range operation.Responses {
		switch statusStr {
		case "default", "2XX":
			defaultResponse = resp
		case "1XX", "3XX", "4XX", "5XX":
			continue
		}
		statusInt, err := strconv.Atoi(statusStr)
		if err != nil {
			continue
		}
		if statusInt/100 == 2 {
			if resp == nil {
				continue
			}
			return resp, statusInt, true
		}
	}
	return defaultResponse, 0, (defaultResponse != nil)
}

// Validate the values of Operation object.
func (operation Operation) Validate() error {
	validaters := []validater{}
	if operation.ExternalDocs != nil {
		validaters = append(validaters, operation.ExternalDocs)
	}
	if hasDuplicatedParameter(operation.Parameters) {
		return ErrParameterDuplicated
	}
	if operation.RequestBody != nil {
		validaters = append(validaters, operation.RequestBody)
	}
	if operation.Responses == nil {
		return ErrRequired{Target: "operation.responses"}
	}
	validaters = append(validaters, operation.Responses)
	for _, callback := range operation.Callbacks {
		validaters = append(validaters, callback)
	}
	for _, security := range operation.Security {
		validaters = append(validaters, security)
	}
	for _, server := range operation.Servers {
		validaters = append(validaters, server)
	}
	return validateAll(validaters)
}
