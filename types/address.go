package types

// These are shared between ZipCodeLookUp and Verify

type AddressRequest struct {
	ID           string `xml:"ID,attr,omitempty"`      // ID for the individual address lookup
	FirmName     string `xml:"FirmName,omitempty"`     // Firm Name
	Address1     string `xml:"Address1"`               // Delivery Address in the destination address
	Address2     string `xml:"Address2"`               // Delivery Address in the destination address. Required for all mail and packages, however 11-digit Destination Delivery Point ZIP+4 Code can be provided as an alternative in the Detail 1 Record
	City         string `xml:"City"`                   // City name of the destination address
	State        string `xml:"State"`                  // Two-character state code of the destination address
	Urbanization string `xml:"Urbanization,omitempty"` // Urbanization. For Puerto Rico addresses only
	Zip5         string `xml:"Zip5"`                   // Destination 5-digit ZIP Code. Numeric values (0-9) only. If International, all zeroes.
	Zip4         string `xml:"Zip4"`                   // Destination ZIP+4 Numeric values (0-9) only. If International, all zeroes. Default to spaces if not available.
}

type AddressResponse struct {
	*AddressRequest
	CityAbbreviation     string `xml:"CityAbbreviation,omitempty"`
	DeliveryPoint        string `xml:"DeliveryPoint,omitempty"`
	CarrierRoute         string `xml:"CarrierRoute,omitempty"`
	Footnotes            string `xml:"Footnotes,omitempty"`
	DPVConfirmation      string `xml:"DPVConfirmation,omitempty"`
	DPVCMRA              string `xml:"DPVCMRA,omitempty"`
	DPVFootnotes         string `xml:"DPVFootnotes,omitempty"`
	Business             string `xml:"Business,omitempty"`
	CentralDeliveryPoint string `xml:"CentralDeliveryPoint,omitempty"`
	Vacant               string `xml:"Vacant,omitempty"`
	Error                *Error  `xml:"Error,omitempty"`
}
