package test

import (
	"os"

	"github.com/p-lau/usps"
)

func main() {
	var (
		userID = os.Getenv("USERID")
		password = os.Getenv("PASS")
		production = false
	)

	usps.Init(userID, password, production)
}