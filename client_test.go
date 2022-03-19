package usps

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestClient (t *testing.T){
	var (
		userID = os.Getenv("USERID")
		pass = "PASS"
	)

	USPS := Init(userID, pass, true)

	request := SDCGetLocationsRequest{
		USERID: userID,
		MailClass: "0",
		OriginZIP: "00777",
		DestinationZIP: "01337",
	}
	resp, err := USPS.ServiceDeliveryCalculatorGetLocations(&request)
	if err != nil {
		t.Fatal(err)
	}

	str, _ := xml.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}