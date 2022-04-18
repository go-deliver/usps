package types

type CarrierPickupAvailabilityRequest struct {
	request
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	Date         string `xml:"Date,omitempty"`
}

type CarrierPickupAvailabilityResponse struct {
	FirmName     string `xml:"FirmName,omitempty"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	DayOfWeek    string `xml:"DayOfWeek"`
	Date         string `xml:"Date"`
	CarrierRoute string `xml:"CarrierRoute,omitempty"`
}