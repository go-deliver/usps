package usps

import (
	"encoding/json"
	"testing"
)

func TestClient (t *testing.T){
	const (
		userID = "USERID"
		pass = "PASSWORD"
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
		t.Fatal("Can't get a response", err)
	}

	str, _ := json.MarshalIndent(resp, "", " ")

	t.Log(string(str))
}