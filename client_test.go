package usps

import (
	"encoding/json"
	"os"
	"testing"
)

func TestClient (t *testing.T){
	var (
		userID = os.Getenv("USERID")
		pass = os.Getenv("PASS")
	)

	USPS := InitUSPS(userID, pass, true)

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

	str, _ := json.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}