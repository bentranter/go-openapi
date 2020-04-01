package openapi

import (
	"net/url"
)

// codebeat:disable[TOO_MANY_IVARS]

// Info Object
type Info struct {
	Title          string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version,omitempty" yaml:"version,omitempty"`
}

// Validate the values of Info object.
func (info Info) Validate() error {
	if err := info.validateRequiredFields(); err != nil {
		return err
	}
	return info.validateFields()
}

func (info Info) validateRequiredFields() error {
	if info.Title == "" {
		return ErrRequired{Target: "info.title"}
	}
	if info.Version == "" {
		return ErrRequired{Target: "info.version"}
	}
	return nil
}

func (info Info) validateFields() error {
	if info.TermsOfService != "" {
		if _, err := url.ParseRequestURI(info.TermsOfService); err != nil {
			return ErrFormatInvalid{Target: "info.termsOfService", Format: "URL"}
		}
	}

	var validaters []validater
	if info.Contact != nil {
		validaters = append(validaters, info.Contact)
	}
	if info.License != nil {
		validaters = append(validaters, info.License)
	}
	return validateAll(validaters)
}
