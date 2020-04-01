package openapi

// codebeat:disable[TOO_MANY_IVARS]

// Contact Object
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

// Validate the values of Contact object.
func (contact Contact) Validate() error {
	if err := mustURL("contact.url", contact.URL); err != nil {
		return err
	}
	if contact.Email != "" {
		if !emailRegexp.MatchString(contact.Email) {
			return ErrFormatInvalid{Target: "contact.email", Format: "email"}
		}
	}
	return nil
}
