package usps_test

import (
	"encoding/xml"
	"testing"

	"github.com/p-lau/usps"
	"github.com/p-lau/usps/types"
)


var (
	client = usps.Init("179STARL7513")
)

func TestAPI(t *testing.T) {
	t.Log("Beginning " + t.Name())

	address1 := new(types.AddressRequest)
	address1.ID = "1"
	address1.Address1 = "123 Main St"
	address1.Address2 = "Apt 1"
	address1.City = "New York"
	address1.State = "NY"
	address1.Zip5 = "10001"
	address1.Zip4 = "1234"

	t.Log(address1)

	address2 := new(types.AddressRequest)
	address2.ID = "Error"
	address2.Address1 = "SUITE K"
	address2.Address2 = "29851 Aventura"
	address2.City = "RANCHO SANTA MARGARITA"
	address2.State = "CA"

	t.Log(address2)

	req := new(types.AddressValidateRequest)
	req.Revision = "1"
	req.Address = []*types.AddressRequest{address1, address2}

	t.Log(req)

	resp := new(types.AddressValidateResponse)

	if err := client.Call(req, resp); err != nil {
		t.Error(err)
	}
	body, err := xml.MarshalIndent(resp, "", " ")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}