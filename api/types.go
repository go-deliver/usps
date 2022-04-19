package api

import "net/http"

// Client is the main Client Interface
type Client struct {
	username   string
	httpClient *http.Client
}

type request interface {
	// Retrieves the API type for the request
	API() string

	// Sets the user for the request
	SetUser(USERID string)
}

type response interface {}