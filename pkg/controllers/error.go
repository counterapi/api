package controllers

import "fmt"

type QueryParamMissing struct {
	Parameter string
}

func (e QueryParamMissing) Error() string {
	return fmt.Sprintf("Parameter %s is missing from query.", e.Parameter)
}
