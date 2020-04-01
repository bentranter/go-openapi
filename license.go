package openapi

import (
	"net/url"
)

// codebeat:disable[TOO_MANY_IVARS]

// License Object
type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Validate the values of License object.
func (license License) Validate() error {
	if license.Name == "" {
		return ErrRequired{Target: "license.name"}
	}
	if license.URL != "" {
		_, err := url.ParseRequestURI(license.URL)
		if err != nil {
			return ErrFormatInvalid{Target: "license.url", Format: "URL"}
		}
		return nil
	}
	return nil
}
