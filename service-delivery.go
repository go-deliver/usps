package usps

import "encoding/xml"

/*
ServiceDeliveryCalculatorGetLocations is a USPS API

Allows customers to get estimates on delivery standards between 3 or 5 digit
Zip Codes for Priority Mail Express, Priority Mail, First Class Mail,
Marketing Mail, Periodicals, and Package Services.

The data returned by the API is intended
for display only. The content or sequence of the string data returned in
the API response may change.

	Source: https://www.usps.com/business/web-tools-apis/service-delivery-calculator-get-locations-api.htm

*/
func (usps *USPS) ServiceDeliveryCalculatorGetLocations(request *SDCGetLocationsRequest) (SDCGetLocationsResponse, error) {
	request.USERID = usps.Username

	// Initilize result
	result := new(SDCGetLocationsResponse)

	// Perform API call to USPS through the Client interace
	err := usps.Client.do(request, result)

	return *result, err
}

func (r *SDCGetLocationsRequest) toHTTP(bool) (string, error) {
	return createRequest("SDCGetLocations", r)
}

/* Types */

// SDCGetLocationsRequest is the request sent to the USPS ServiceDeliveryGetLocations API
type SDCGetLocationsRequest struct {
	XMLName         xml.Name `xml:"SDCGetLocationsRequest"`
	USERID          string   `xml:"USERID,attr"`
	MailClass       string   `xml:"MailClass"`
	OriginZIP       string   `xml:"OriginZIP"`
	DestinationZIP  string   `xml:"DestinationZIP"`
	AcceptDate      string   `xml:"AcceptDate,omitempty"`
	AcceptTime      string   `xml:"AcceptTime,omitempty"`
	NonEMDetail     string   `xml:"NonEMDetail,omitempty"`
	NonEMOriginType string   `xml:"NonEMOriginType,omitempty"`
	NonEMDestType   string   `xml:"NonEMDestType,omitempty"`
	Weight          string   `xml:"Weight,omitempty"`
}

// SDCGetLocationsResponse is the expected response returned when sending SDCGetLocations request to USPS
type SDCGetLocationsResponse struct {
	// Release version
	Release     string `xml:"Release"`
	// Caller ID
	CallerID    string `xml:"CallerID"`
	SourceID    string `xml:"SourceID"`
	// MailClass from request
	MailClass   string `xml:"MailClass"`
	// OriginZIP from request
	OriginZIP   string `xml:"OriginZIP"`
	// Origin City
	OriginCity  string `xml:"OriginCity"`
	// Origin State
	OriginState string `xml:"OriginState"`
	DestZIP     string `xml:"DestZIP"`
	DestCity    string `xml:"DestCity"`
	DestState   string `xml:"DestState"`
	AcceptDate  string `xml:"AcceptDate"`
	AcceptTime  string `xml:"AcceptTime"`
	Weight      string `xml:"Weight,omitempty"`
	Expedited   struct {
		EAD        string `xml:"EAD,omitempty"`
		Commitment []struct {
			MailClass  string `xml:"MailClass,omitempty"`
			CommitmentName string `xml:"CommitmentName,omitempty"`
			CommitmentTime string `xml:"CommitmentTime,omitempty"`
			CommitmentSeq  string `xml:"CommitmentSeq,omitempty"`
			Location       []struct {
				SDD          string `xml:"SDD,omitempty"`
				COT          string `xml:"COT,omitempty"`
				FacType      string `xml:"FacType,omitempty"`
				Street       string `xml:"Street,omitempty"`
				City         string `xml:"City,omitempty"`
				State        string `xml:"State,omitempty"`
				ZIP          string `xml:"ZIP,omitempty"`
				IsGuaranteed string `xml:"IsGuaranteed,omitempty"`
			} `xml:"Location,omitempty"`
		} `xml:"Commitment,omitempty"`
		ExpeditedMessage  string `xml:"ExpeditedMessage,omitempty"`
		ExpeditedTransMsg []struct {
			MsgCode string `xml:"MsgCode,omitempty"`
			Msg     string `xml:"Msg,omitempty"`
		} `xml:"ExpeditedTransMsg,omitempty"`
	} `xml:"Expedited,omitempty"`
	NonExpedited []struct {
		MailClass            string `xml:"MailClass,omitempty"`
		NonExpeditedDestType string `xml:"NonExpeditedDestType,omitempty"`
		EAD                  string `xml:"EAD,omitempty"`
		SvcStdDays           string `xml:"SvcStdDays,omitempty"`
		SchedDlvryDate       string `xml:"SchedDlvryDate,omitempty"`
		HFPU                 struct {
			EAD             string `xml:"EAD,omitempty"`
			COT             string `xml:"COT,omitempty"`
			ServiceStandard struct {
				SvcStdMsg  string `xml:"SvcStdMsg,omitempty"`
				SvcStdDays string `xml:"SvcStdDays,omitempty"`
				Location   struct {
					SchedDlvryDate  string `xml:"SchedDlvryDate,omitempty"`
					NonDeliveryDays string `xml:"NonDeliveryDays,omitempty"`
					RAUName         string `xml:"RAUName,omitempty"`
					Street          string `xml:"Street,omitempty"`
					ZIP             string `xml:"ZIP,omitempty"`
					CloseTimes      struct {
						M  string `xml:"M,omitempty"`
						Tu string `xml:"Tu,omitempty"`
						W  string `xml:"W,omitempty"`
						Th string `xml:"Th,omitempty"`
						F  string `xml:"F,omitempty"`
						Sa string `xml:"Sa,omitempty"`
						Su string `xml:"Su,omitempty"`
						// Close time for Holiday
						H  string `xml:"H,omitempty"`
					} `xml:"CloseTimes,omitempty"`
					NonExpeditedExceptions struct {
						NonExpeditedMsg      []string `xml:"NonExpeditedMsg,omitempty"`
						NonExpeditedTransMsg []struct {
							MsgCode string `xml:"MsgCode,omitempty"`
							Msg     string `xml:"Msg,omitempty"`
						} `xml:"NonExpeditedTransMsg,omitempty"`
					} `xml:"NonExpeditedExceptions,omitempty"`
					City  string `xml:"City,omitempty"`
					State string `xml:"State,omitempty"`
				} `xml:"Location,omitempty"`
			} `xml:"ServiceStandard,omitempty"`
			HFPUGlobalExcept struct {
				PostCOT              string `xml:"PostCOT,omitempty"`
				OverMaxResults       string `xml:"OverMaxResults,omitempty"`
				NoHFPULocInd         string `xml:"NoHFPULocInd,omitempty"`
				NonExpeditedWTMsg    []string `xml:"NonExpeditedWTMsg,omitempty"`
				NonExpeditedTransMsg struct {
					MsgCode string `xml:"MsgCode,omitempty"`
					Msg     string `xml:"Msg,omitempty"`
				} `xml:"NonExpeditedTransMsg,omitempty"`
			} `xml:"HFPUGlobalExcept,omitempty"`
		} `xml:"HFPU,omitempty"`
	} `xml:"NonExpedited,omitempty"`
}