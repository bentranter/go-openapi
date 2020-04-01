package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Discriminator Object
type Discriminator struct {
	PropertyName string            `json:"propertyName,omitempty" yaml:"propertyName,omitempty"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}

// Validate the values of Descriminator object.
func (discriminator Discriminator) Validate() error {
	if discriminator.PropertyName == "" {
		return ErrRequired{Target: "discriminator.propertyName"}
	}
	return nil
}
