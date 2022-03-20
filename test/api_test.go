package test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps"
	"github.com/p-lau/usps/request"
)

func Test(t *testing.T) {
	var (
		userID = os.Getenv("USERID")
		password = os.Getenv("PASS")
	)

	usps := usps.Init(userID, password)

	req := request.SDCGetLocationsRequest{
		MailClass: "0",
		OriginZIP: "00732",
		DestinationZIP: "01852",
	}

	res, err := usps.(&req)
	if err != nil{
		t.Fatal(err)
	}
	marshaled, err := xml.MarshalIndent(res, "", " ")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(string(marshaled))
}