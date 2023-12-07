/*
 * AppRepo - Get apps
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AppCreatePostRequest struct {
	Name string `json:"name"`

	Description string `json:"description"`
}

// AssertAppCreatePostRequestRequired checks if the required fields are not zero-ed
func AssertAppCreatePostRequestRequired(obj AppCreatePostRequest) error {
	elements := map[string]interface{}{
		"name":        obj.Name,
		"description": obj.Description,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseAppCreatePostRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AppCreatePostRequest (e.g. [][]AppCreatePostRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAppCreatePostRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAppCreatePostRequest, ok := obj.(AppCreatePostRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAppCreatePostRequestRequired(aAppCreatePostRequest)
	})
}
