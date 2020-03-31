package openapi_test

import (
	"testing"

	openapi "github.com/bentranter/go-openapi"
)

func TestDiscriminator_Validate(t *testing.T) {
	candidates := []candidate{
		{"empty", openapi.Discriminator{}, openapi.ErrRequired{Target: "discriminator.propertyName"}},
		{"withPropertyName", openapi.Discriminator{PropertyName: "foobar"}, nil},
	}
	testValidater(t, candidates)
}
