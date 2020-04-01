package openapi

// codebeat:disable[TOO_MANY_IVARS]

// ServerVariable Object
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

// Validate the values of Server Variable object.
func (sv ServerVariable) Validate() error {
	if sv.Default == "" {
		return ErrRequired{Target: "serverVariable.default"}
	}
	return nil
}
