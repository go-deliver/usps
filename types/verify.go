package types

func (r *AddressValidateRequest) ToHTTP() (string, error) {
	return createRequest("Verify", r)
}

type AddressValidateRequest struct {
	USERID   string     `xml:"USERID,attr"`        // Your Web Tools ID.
	Revision string     `xml:"Revision,omitempty"` // Integer value used to return of all available response fields. Set this value to 1 to return all currently documented response fields
	Address  []*Address `xml:"Address"`            // Up to 5 address verifications can be included per transaction.
}

type Address struct {
	ID           string `xml:"ID,attr"`      // ID for the individual address lookup
	FirmName     string `xml:"FirmName"`     // Firm Name
	Address1     string `xml:"Address1"`     // Delivery Address in the destination address
	Address2     string `xml:"Address2"`     // Delivery Address in the destination address. Required for all mail and packages, however 11-digit Destination Delivery Point ZIP+4 Code can be provided as an alternative in the Detail 1 Record
	City         string `xml:"City"`         // City name of the destination address
	State        string `xml:"State"`        // Two-character state code of the destination address
	Urbanization string `xml:"Urbanization"` // Urbanization. For Puerto Rico addresses only
	Zip5         string `xml:"Zip5"`         // Destination 5-digit ZIP Code. Numeric values (0-9) only. If International, all zeroes.
	Zip4         string `xml:"Zip4"`         // Destination ZIP+4 Numeric values (0-9) only. If International, all zeroes. Default to spaces if not available.
}

type AddressValidateResponse struct {
	Address []*struct {
		Text                 string `xml:",chardata"`
		ID                   string `xml:"ID,attr"`
		Address1             string `xml:"Address1"`
		Address2             string `xml:"Address2"`
		City                 string `xml:"City"`
		CityAbbreviation     string `xml:"CityAbbreviation"`
		State                string `xml:"State"`
		Zip5                 string `xml:"Zip5"`
		Zip4                 string `xml:"Zip4"`
		DeliveryPoint        string `xml:"DeliveryPoint"`
		CarrierRoute         string `xml:"CarrierRoute"`
		Footnotes            string `xml:"Footnotes"`
		DPVConfirmation      string `xml:"DPVConfirmation"`
		DPVCMRA              string `xml:"DPVCMRA"`
		DPVFootnotes         string `xml:"DPVFootnotes"`
		Business             string `xml:"Business"`
		CentralDeliveryPoint string `xml:"CentralDeliveryPoint"`
		Vacant               string `xml:"Vacant"`
		Error                Error  `xml:"Error,omitempty"`
	} `xml:"Address"`
}
