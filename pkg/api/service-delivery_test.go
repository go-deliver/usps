package api

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps/pkg/request"
)

func TestResponse(t *testing.T) {
	var userID = os.Getenv("USERID")

	api := API{Username: userID}

	request := request.SDCGetLocationsRequest{
		USERID:         userID,
		MailClass:      "6",
		OriginZIP:      "00777",
		DestinationZIP: "01337",
	}

	resp, err := api.ServiceDeliveryCalculatorGetLocations(&request)
	if err != nil {
		t.Fatal(err)
	}

	str, _ := xml.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}
