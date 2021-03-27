package middlewares

// TooManyRequestError for too many requests.
type TooManyRequestError struct{}

// Error fails for too many requests.
func (e TooManyRequestError) Error() string {
	return "too many requests"
}
