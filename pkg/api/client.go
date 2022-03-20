package api

import (
	"github.com/p-lau/usps/pkg/request"
)

// Client is the interface that calls the USPS API
type Client interface {
	do(request request.Request, result interface{}) error
	call(requestURL string) ([]byte, error)
}