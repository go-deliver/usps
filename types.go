package usps

import (
	"encoding/xml"
	"net/http"
)

type Error struct {
	Number      string `xml:"Number"`
	Source      string `xml:"Source,omitempty"`
	Description string `xml:"Description,omitempty"`
	HelpFile    string `xml:"HelpFile,omitempty"`
	HelpContext string `xml:"HelpContext,omitempty"`
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type USPSClient interface {
	do(request USPSRequest, result interface{}) error
	call(requestURL string) ([]byte, error)
}

type API struct {
	USPSClient
	HttpClient HttpClient
	Production bool
}

type USPSRequest interface {
	toHTTPRequestStr(isProduction bool) (string, error)
}

type USPS struct {
	Username   string
	Password   string
	Production bool `default:"false"`
	Client     USPSClient
}

type CloseTimes struct {
	M  string `xml:"M,omitempty"`
	Tu string `xml:"Tu,omitempty"`
	W  string `xml:"W,omitempty"`
	Th string `xml:"Th,omitempty"`
	F  string `xml:"F,omitempty"`
	Sa string `xml:"Sa,omitempty"`
	Su string `xml:"Su,omitempty"`
	H  string `xml:"H,omitempty"`
}

type TransMsg struct {
	MsgCode string `xml:"MsgCode,omitempty"`
	Msg     string `xml:"Msg,omitempty"`
}

type NonExpeditedExceptions struct {
	NonExpeditedMsg      string `xml:"NonExpeditedMsg,omitempty"`
	NonExpeditedTransMsg []TransMsg `xml:"NonExpeditedTransMsg,omitempty"`
}

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

type SDCGetLocationsResponse struct {
	Release     string `xml:"Release"`
	CallerID    string `xml:"CallerID"`
	SourceID    string `xml:"SourceID"`
	MailClass   string `xml:"MailClass"`
	OriginZIP   string `xml:"OriginZIP"`
	OriginCity  string `xml:"OriginCity"`
	OriginState string `xml:"OriginState"`
	DestZIP     string `xml:"DestZIP"`
	DestCity    string `xml:"DestCity"`
	DestState   string `xml:"DestState"`
	AcceptDate  string `xml:"AcceptDate"`
	AcceptTime  string `xml:"AcceptTime"`
	Weight      string `xml:"Weight,omitempty"`
	Expedited   struct {
		EAD        string `xml:"EAD,omitempty"`
		MailClass  string `xml:"MailClass,omitempty"`
		Commitment []struct {
			CommitmentName string `xml:"CommitmentName,omitempty"`
			CommitmentTime string `xml:"CommitmentTime,omitempty"`
			CommitmentSeq  string `xml:"CommitmentSeq,omitempty"`
			Location       struct {
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
		ExpeditedTransMsg []TransMsg `xml:"ExpeditedTransMsg,omitempty"`
	} `xml:"Expedited,omitempty"`
	NonExpedited struct {
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
					CloseTimes      CloseTimes `xml:"CloseTimes,omitempty"`
					NonExpeditedExceptions NonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"`
					City  string `xml:"City,omitempty"`
					State string `xml:"State,omitempty"`
				} `xml:"Location,omitempty"`
			} `xml:"ServiceStandard,omitempty"`
			HFPUGlobalExcept struct {
				PostCOT              string `xml:"PostCOT,omitempty"`
				OverMaxResults       string `xml:"OverMaxResults,omitempty"`
				NoHFPULocInd         string `xml:"NoHFPULocInd,omitempty"`
				NonExpeditedWTMsg    string `xml:"NonExpeditedWTMsg,omitempty"`
				NonExpeditedTransMsg []TransMsg `xml:"NonExpeditedTransMsg,omitempty"`
			} `xml:"HFPUGlobalExcept,omitempty"`
		} `xml:"HFPU,omitempty"`
	} `xml:"NonExpedited,omitempty"`
	Commitment string `xml:"Commitment,omitempty"`
}
