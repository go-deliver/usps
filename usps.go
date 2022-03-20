package usps

import (
	"github.com/p-lau/usps/pkg/api"
)

// Init starts up the USPS client and returns a pointer for the client
func Init(username, password string, production bool) *api.API {
	return &api.API{
		Username: username,
		Password: password,
		Production: production,
	}
}