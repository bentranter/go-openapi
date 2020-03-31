package openapi_test

import (
	"testing"

	openapi "github.com/bentranter/go-openapi"
)

func TestRequestBody_Validate(t *testing.T) {
	contentRequiredError := openapi.ErrRequired{Target: "requestBody.content"}
	candidates := []candidate{
		{"empty", openapi.RequestBody{}, contentRequiredError},
		{"emptyContent", openapi.RequestBody{Content: map[string]*openapi.MediaType{}}, contentRequiredError},
		{"valid", openapi.RequestBody{Content: map[string]*openapi.MediaType{"application/json": &openapi.MediaType{}}}, nil},
	}
	testValidater(t, candidates)
}
