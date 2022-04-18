package types

func (r *PriorityMailRequest) API() string {
	return "PriorityMail"
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
	response
	OriginZIP               string `xml:"OriginZIP"`
	DestinationZIP          string `xml:"DestinationZIP,omitempty"`
	Days                    string `xml:"Days,omitempty"`
	IsGuaranteed            string `xml:"IsGuaranteed,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
}
