package types

func (r *FirstClassMailRequest) API() string {
	return "FirstClassMail"
}

type FirstClassMailRequest struct {
	request
	OriginZIP       string `xml:"OriginZIP"`
	DestinationZIP  string `xml:"DestinationZIP"`
	DestinationType string `xml:"DestinationType,omitempty"`
	MailClass       string `xml:"MailClass"`
}

type FirstClassMailResponse struct {
	OriginZIP               string `xml:"OriginZIP"`
	DestinationZIP          string `xml:"DestinationZIP,omitempty"`
	Days                    string `xml:"Days,omitempty"`
	Message                 string `xml:"Message,omitempty"`
	EffectiveAcceptanceDate string `xml:"EffectiveAcceptanceDate,omitempty"`
	ScheduledDeliveryDate   string `xml:"ScheduledDeliveryDate,omitempty"`
}
