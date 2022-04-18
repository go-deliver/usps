package usps

import (
	"github.com/p-lau/usps/api"
)

// Init starts up the USPS client and returns a pointer for the client
func Init(username string) *api.Client {
	client := api.New()
	client.SetUser(username)
	return client
}