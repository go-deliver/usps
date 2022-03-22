package test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps"
	"github.com/p-lau/usps/types"
)

var(
	userID = os.Getenv("USERID")
	password = os.Getenv("PASS")
	api = usps.Init(userID, password)
)

func TestSDCGETLocationsAPI(t *testing.T) {
	req := types.SDCGetLocationsRequest{
		MailClass: "0",
		OriginZIP: "00732",
		DestinationZIP: "01852",
	}

	res, err := api.SDCGetLocations(&req)
	if err != nil{
		t.Fatal(err)
	}
	marshaled, err := xml.MarshalIndent(res, "", " ")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(string(marshaled))
}

func TestVerifyAPI(t *testing.T){
	address := types.AddressRequest{
		ID: "TestAddress",
		Address1: "SUITE K",
		Address2: "29851 Aventura",
		City: "RANCHO SANTA MARGARITA",
		State: "CA",
	}

	address2 := types.AddressRequest{
		ID: "ERROR",
		Address2: "BAD ADDRESS",
		City: "CAMBRIDGE",
		State: "MA",
	}

	req := types.AddressValidateRequest{
		Revision: "1",
		Address: []*types.AddressRequest{&address, &address2},
	}

	res, err := api.Verify(&req)
	if err != nil{
		t.Fatal(err)
	}
	marshaled, err := xml.MarshalIndent(res, "", " ")
	if err != nil{
		t.Fatal(err)
	}

	t.Log(string(marshaled))
}