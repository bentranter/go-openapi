package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Parameter Object
type Parameter struct {
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	In              InType `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      string `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	Style         string              `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema        *Schema             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example       interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	Examples      map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`

	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Validate the values of Parameter object.
// This function DOES NOT check whether the name field correspond to the associated path or not.
func (parameter Parameter) Validate() error {
	if err := parameter.validateRequiredObjects(); err != nil {
		return err
	}
	switch parameter.In {
	case InQuery, InHeader, InPath, InCookie:
	default:
		return ErrMustOneOf{Object: "parameter.in", ValidValues: ParameterInList}
	}
	if parameter.In == InPath && !parameter.Required {
		return ErrRequiredMustTrue
	}
	if parameter.In != InQuery && parameter.AllowEmptyValue {
		return ErrAllowEmptyValueNotValid
	}
	if len(parameter.Content) > 1 {
		return ErrTooManyParameterContent
	}

	return validateAll(parameter.reduceValidaters())
}

func (parameter Parameter) validateRequiredObjects() error {
	if parameter.Name == "" {
		return ErrRequired{Target: "parameter.name"}
	}
	if parameter.In == "" {
		return ErrRequired{Target: "parameter.in"}
	}
	return nil
}

func (parameter Parameter) reduceValidaters() []validater {
	validaters := []validater{}
	if parameter.Schema != nil {
		validaters = append(validaters, parameter.Schema)
	}
	if v, ok := parameter.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	for _, mediaType := range parameter.Content {
		validaters = append(validaters, mediaType)
	}
	return validaters
}
