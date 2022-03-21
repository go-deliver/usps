package api

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps/request"
)

func TestResponse(t *testing.T) {
	var userID = os.Getenv("USERID")

	api := API{Username: userID}

	request := request.SDCGetLocationsRequest{
		USERID:         userID,
		MailClass:      "0",
		OriginZIP:      "01852",
		DestinationZIP: "90011",
	}

	resp, err := api.SDCGetLocations(&request)
	if err != nil {
		t.Fatal(err)
	}

	str, _ := xml.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}
