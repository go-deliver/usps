package types

func (r *PriorityMailRequest) ToHTTP() (string, error) {
	return createRequest("PriorityMail", r)
}

type PriorityMailRequest struct {
	request
	OriginZIP       string `xml:"OriginZIP"`
	DestinationZIP  string `xml:"DestinationZIP"`
	DestinationType string `xml:"DestinationType,omitempty"`
	PMGuarantee     string `xml:"PMGuarantee,omitempty"`
	MailClass       string `xml:"MailClass"`
}

type PriorityMailResponse struct {
	OriginZIP               string `xml:"OriginZIP"`
	DestinationZIP          string `xml:"DestinationZIP,omitempty"`
	Days                    string `xml:"Days,omitempty"`
	IsGuaranteed            string `xml:"IsGuaranteed,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
}
