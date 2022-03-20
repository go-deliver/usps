package api

import (
	"github.com/p-lau/usps/pkg/request"
)

// Client is the interface that calls the USPS API
type Client interface {
	Do(request request.Request, result interface{}) error
	Call(requestURL string) ([]byte, error)
}