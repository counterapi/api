package controllers

import "fmt"

// QueryParamMissingError is a error for missing query for a controller.
type QueryParamMissingError struct {
	Parameter string
}

// Error fails for the error.
func (e QueryParamMissingError) Error() string {
	return fmt.Sprintf("Parameter %s is missing from query.", e.Parameter)
}
