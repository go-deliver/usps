package types

func (r *CarrierPickupScheduleRequest) ToHTTP() (string, error) {
	return createRequest("CarrierPickupSchedule", r)
}

type CarrierPickupScheduleRequest struct {
	request
	FirstName           string                  `xml:"FirstName"`
	LastName            string                  `xml:"LastName"`
	FirmName            string                  `xml:"FirmName,omitempty"`
	SuiteOrApt          string                  `xml:"SuiteOrApt"`
	Address2            string                  `xml:"Address2"`
	Urbanization        string                  `xml:"Urbanization"`
	City                string                  `xml:"City"`
	State               string                  `xml:"State"`
	ZIP5                string                  `xml:"ZIP5"`
	ZIP4                string                  `xml:"ZIP4"`
	Phone               string                  `xml:"Phone"`
	Extension           string                  `xml:"Extension,omitempty"`
	Package             []*CarrierPickupPackage `xml:"Package"`
	EstimatedWeight     string                  `xml:"EstimatedWeight"`
	PackageLocation     string                  `xml:"PackageLocation"`
	SpecialInstructions string                  `xml:"SpecialInstructions,omitempty"`
	EmailAddress        string                  `xml:"EmailAddress,omitempty"`
}

type CarrierPickupScheduleResponse struct {
	FirstName           string                  `xml:"FirstName"`
	LastName            string                  `xml:"LastName"`
	FirmName            string                  `xml:"FirmName,omitempty"`
	SuiteOrApt          string                  `xml:"SuiteOrApt"`
	Address2            string                  `xml:"Address2"`
	Urbanization        string                  `xml:"Urbanization"`
	City                string                  `xml:"City"`
	State               string                  `xml:"State"`
	ZIP5                string                  `xml:"ZIP5"`
	ZIP4                string                  `xml:"ZIP4"`
	Phone               string                  `xml:"Phone"`
	Extension           string                  `xml:"Extension,omitempty"`
	Package             []*CarrierPickupPackage `xml:"Package"`
	EstimatedWeight     string                  `xml:"EstimatedWeight"`
	PackageLocation     string                  `xml:"PackageLocation"`
	SpecialInstructions string                  `xml:"SpecialInstructions"`
	ConfirmationNumber  string                  `xml:"ConfirmationNumber"`
	DayOfWeek           string                  `xml:"DayOfWeek"`
	Date                string                  `xml:"Date"`
	CarrierRoute        string                  `xml:"CarrierRoute,omitempty"`
	EmailAddress        string                  `xml:"EmailAddress,omitempty"`
}
