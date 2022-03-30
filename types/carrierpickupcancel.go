package types

func (r *CarrierPickupCancelRequest) ToHTTP() (string, error) {
	return createRequest("CarrierPickupCancel", r)
}

type CarrierPickupCancelRequest struct {
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

type CarrierPickupCancelResponse struct {
	FirmName           string `xml:"FirmName"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
	Status             string `xml:"Status"`
}
