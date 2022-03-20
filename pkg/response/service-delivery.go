package response

// SDCGetLocationsResponse is the expected response returned when sending SDCGetLocations request to USPS
type SDCGetLocationsResponse struct {
	Release      string          `xml:"Release"`     // Release version
	CallerID     string          `xml:"CallerID"`    // Caller ID
	SourceID     string          `xml:"SourceID"`    // Source ID
	MailClass    string          `xml:"MailClass"`   // MailClass from request
	OriginZIP    string          `xml:"OriginZIP"`   // OriginZIP from request
	OriginCity   string          `xml:"OriginCity"`  // Origin City
	OriginState  string          `xml:"OriginState"` // Origin State
	DestZIP      string          `xml:"DestZIP"`     // Destination ZIP Code
	DestCity     string          `xml:"DestCity"`    // Destination City
	DestState    string          `xml:"DestState"`   // Destination State
	AcceptDate   string          `xml:"AcceptDate"`  //
	AcceptTime   string          `xml:"AcceptTime"`
	Weight       string          `xml:"Weight,omitempty"`
	Expedited    *expedited      `xml:"Expedited,omitempty"`
	NonExpedited []*nonExpedited `xml:"NonExpedited,omitempty"`
}

type expedited struct {
	EAD               string                 `xml:"EAD,omitempty"`
	Commitment        []*expeditedCommitment `xml:"Commitment,omitempty"`
	ExpeditedMessage  string                 `xml:"ExpeditedMessage,omitempty"`
	ExpeditedTransMsg []*transMsg            `xml:"ExpeditedTransMsg,omitempty"`
}

type expeditedCommitment struct {
	MailClass      string               `xml:"MailClass,omitempty"`
	CommitmentName string               `xml:"CommitmentName,omitempty"`
	CommitmentTime string               `xml:"CommitmentTime,omitempty"`
	CommitmentSeq  string               `xml:"CommitmentSeq,omitempty"`
	Locations      []*expeditedLocation `xml:"Location,omitempty"`
}

type expeditedLocation struct {
	SDD          string `xml:"SDD,omitempty"`
	COT          string `xml:"COT,omitempty"`
	FacType      string `xml:"FacType,omitempty"`
	Street       string `xml:"Street,omitempty"`
	City         string `xml:"City,omitempty"`
	State        string `xml:"State,omitempty"`
	ZIP          string `xml:"ZIP,omitempty"`
	IsGuaranteed string `xml:"IsGuaranteed,omitempty"`
}

type nonExpedited struct {
	MailClass            string `xml:"MailClass,omitempty"`
	NonExpeditedDestType string `xml:"NonExpeditedDestType,omitempty"`
	EAD                  string `xml:"EAD,omitempty"`
	SvcStdDays           string `xml:"SvcStdDays,omitempty"`
	SchedDlvryDate       string `xml:"SchedDlvryDate,omitempty"`
	HFPU                 *hfpu  `xml:"HFPU,omitempty"`
	NonExpeditedExceptions *nonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"`
}

type nonExpeditedExceptions struct {
	NonExpeditedTransMsg *transMsg `xml:"NonExpeditedTransMsg,omitempty"`
}

type hfpu struct {
	EAD             string           `xml:"EAD,omitempty"`
	COT             string           `xml:"COT,omitempty"`
	ServiceStandard *serviceStandard `xml:"ServiceStandard,omitempty"`
	GlobalExcept    *globalExcept    `xml:"HFPUGlobalExcept,omitempty"`
}

type globalExcept struct {
	PostCOT                    string    `xml:"PostCOT,omitempty"`
	OverMaxResults             string    `xml:"OverMaxResults,omitempty"`
	NoHFPULocInd               string    `xml:"NoHFPULocInd,omitempty"`
	NonExpeditedWTMsg          []string  `xml:"NonExpeditedWTMsg,omitempty"`
	GlobalNonExpeditedTransMsg *transMsg `xml:"NonExpeditedTransMsg,omitempty"`
}

type serviceStandard struct {
	ServiceStanardMessage string                `xml:"SvcStdMsg,omitempty"`
	ServiceStandardDays   string                `xml:"SvcStdDays,omitempty"`
	Location              *nonExpeditedLocation `xml:"Location,omitempty"`
}

type nonExpeditedLocation struct {
	ScheduledDeliveryDate    string      `xml:"SchedDlvryDate,omitempty"`
	NonDeliveryDays          string      `xml:"NonDeliveryDays,omitempty"`
	RAUName                  string      `xml:"RAUName,omitempty"`
	Street                   string      `xml:"Street,omitempty"`
	ZIP                      string      `xml:"ZIP,omitempty"`
	City                     string      `xml:"City,omitempty"`
	State                    string      `xml:"State,omitempty"`
	CloseTimes               *closeTimes `xml:"CloseTimes,omitempty"`
	NonExpeditedMsg          string      `xml:"NonExpeditedExceptions>NonExpeditedMsg,omitempty"`
	HFPUNonExpeditedTransMsg transMsg    `xml:"NonExpeditedExceptions>NonExpeditedTransMsg,omitempty"`
}

type closeTimes struct {
	M  string `xml:"M,omitempty"`
	Tu string `xml:"Tu,omitempty"`
	W  string `xml:"W,omitempty"`
	Th string `xml:"Th,omitempty"`
	F  string `xml:"F,omitempty"`
	Sa string `xml:"Sa,omitempty"`
	Su string `xml:"Su,omitempty"`
	H  string `xml:"H,omitempty"`
}

type transMsg struct {
	MsgCode string `xml:"MsgCode,omitempty"`
	Msg     string `xml:"Msg,omitempty"`
}
