package usps

// Request is the interface that utilizes toHTTP method to create a request body
type Request interface {
	// toHTTP is the request method to transcribe
	toHTTP(isProduction bool) (string, error)
}