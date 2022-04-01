package types

type CarrierPickupChangeRequest struct {
	request
	FirstName           string                  `xml:"FirstName"`
	LastName            string                  `xml:"LastName"`
	FirmName            string                  `xml:"FirmName,omitempty"`
	SuiteOrApt          string                  `xml:"SuiteOrApt,omitempty"`
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
	ConfirmationNumber  string                  `xml:"ConfirmationNumber,omitempty"`
	EmailAddress        string                  `xml:"EmailAddress,omitempty"`
}
