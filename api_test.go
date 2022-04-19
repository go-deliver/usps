package usps_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/p-lau/usps"
	"github.com/p-lau/usps/types"
)

var (
	client = usps.Init(os.Getenv("USERID"))
)

func TestAPI(t *testing.T) {
	address1 := new(types.AddressRequest)
	address1.ID = "Error"
	address1.Address2 = "123 Main St"
	address1.City = "Flushing"
	address1.State = "NY"
	address1.Zip5 = "11354"
	address1.Zip4 = "1234"

	address2 := new(types.AddressRequest)
	address2.ID = "Actual"
	address2.Address1 = "SUITE K"
	address2.Address2 = "29851 Aventura"
	address2.City = "RANCHO SANTA MARGARITA"
	address2.State = "CA"

	req := new(types.AddressValidateRequest)
	req.Revision = "1"
	req.Address = []*types.AddressRequest{address1, address2}

	resp := new(types.AddressValidateResponse)

	if err := client.Call(req, resp); err != nil {
		t.Error(err)
	}
	body, err := xml.MarshalIndent(resp, "", " ")
	if err != nil {
		t.Error(err)
	}

	t.Log("Returned body:\n",string(body))
}