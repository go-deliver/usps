package api

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps/pkg/request"
)

func TestRequest (t *testing.T){
	var userID = os.Getenv("USERID")

	api := API{
		Username: userID,
		Production: true,
	}

	request := request.SDCGetLocationsRequest{
		USERID: userID,
		MailClass: "0",
		OriginZIP: "00777",
		DestinationZIP: "01337",
	}

	resp, err := api.ServiceDeliveryCalculatorGetLocations(&request)
	if err != nil {
		t.Fatal(err)
	}

	str, _ := xml.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}