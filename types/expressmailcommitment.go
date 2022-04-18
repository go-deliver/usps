package types

func (r *ExpressMailCommitmentRequest) API() string {
	return "ExpressMailCommitment"
}

type ExpressMailCommitmentRequest struct {
	request
	OriginZIP      string `xml:"OriginZIP"`
	DestinationZIP string `xml:"DestinationZIP"`
	Date           string `xml:"Date"`
	DropOffTime    string `xml:"DropOffTime,omitempty"`
	ReturnDates    string `xml:"ReturnDates,omitempty"`
	PMGuarantee    string `xml:"PMGuarantee,omitempty"`
}

type ExpressMailCommitmentResponse struct {
	OriginZip               string                   `xml:"OriginZip"`
	OriginCity              string                   `xml:"OriginCity"`
	OriginState             string                   `xml:"OriginState"`
	DestinationZip          string                   `xml:"DestinationZip"`
	DestinationCity         string                   `xml:"DestinationCity"`
	DestinationState        string                   `xml:"DestinationState"`
	Date                    string                   `xml:"Date"`
	Time                    string                   `xml:"Time"`
	ExpeditedTransMessage   string                   `xml:"ExpeditedTransMessage,omitempty"`
	MsgCode                 string                   `xml:"MsgCode,omitempty"`
	Msg                     string                   `xml:"Msg,omitempty"`
	EffectiveAcceptanceDate string                   `xml:"EffectiveAcceptanceDate,omitempty"`
	Commitment              []*ExpressMailCommitment `xml:"Commitment,omitempty"`
	Message                 string                   `xml:"Message,omitempty"`
}

type ExpressMailCommitment struct {
	Name     string                         `xml:"Name,omitempty"`
	Time     string                         `xml:"Time,omitempty"`
	Sequence string                         `xml:"Sequence,omitempty"`
	Location *ExpressMailCommitmentLocation `xml:"Location,omitempty"`
}

type ExpressMailCommitmentLocation struct {
	ScheduledDeliveryDate string `xml:"ScheduledDeliveryDate,omitempty"`
	CutOff                string `xml:"CutOff,omitempty"`
	Facility              string `xml:"Facility,omitempty"`
	Street                string `xml:"Street,omitempty"`
	City                  string `xml:"City,omitempty"`
	State                 string `xml:"State,omitempty"`
	Zip                   string `xml:"Zip,omitempty"`
	IsGuaranteed          string `xml:"IsGuaranteed,omitempty"`
}
