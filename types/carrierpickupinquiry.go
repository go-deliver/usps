package types

type CarrierPickupInquiryRequest struct {
	request
	FirmName           string `xml:"FirmName,omitempty"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
}

type CarrierPickupInquiryResponse struct {
	FirstName           string                  `xml:"FirstName"`
	LastName            string                  `xml:"LastName"`
	FirmName            string                  `xml:"FirmName"`
	SuiteOrApt          string                  `xml:"SuiteOrApt"`
	Address2            string                  `xml:"Address2"`
	Urbanization        string                  `xml:"Urbanization"`
	City                string                  `xml:"City"`
	State               string                  `xml:"State"`
	ZIP5                string                  `xml:"ZIP5"`
	ZIP4                string                  `xml:"ZIP4"`
	Phone               string                  `xml:"Phone"`
	Extension           string                  `xml:"Extension"`
	Package             []*CarrierPickupPackage `xml:"Package"`
	EstimatedWeight     string                  `xml:"EstimatedWeight"`
	PackageLocation     string                  `xml:"PackageLocation"`
	SpecialInstructions string                  `xml:"SpecialInstructions"`
	ConfirmationNumber  string                  `xml:"ConfirmationNumber"`
	DayOfWeek           string                  `xml:"DayOfWeek"`
	Date                string                  `xml:"Date"`
	EmailAddress        string                  `xml:"EmailAddress,omitempty"`
}
