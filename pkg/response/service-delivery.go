package response

// SDCGetLocationsResponse is the expected response returned when sending SDCGetLocations request to USPS
type SDCGetLocationsResponse struct {
	Release      string         `xml:"Release"`     // Release version
	CallerID     string         `xml:"CallerID"`    // Caller ID
	SourceID     string         `xml:"SourceID"`    // Source ID
	MailClass    string         `xml:"MailClass"`   // MailClass from request
	OriginZIP    string         `xml:"OriginZIP"`   // OriginZIP from request
	OriginCity   string         `xml:"OriginCity"`  // Origin City
	OriginState  string         `xml:"OriginState"` // Origin State
	DestZIP      string         `xml:"DestZIP"`     // Destination ZIP Code
	DestCity     string         `xml:"DestCity"`    // Destination City
	DestState    string         `xml:"DestState"`   // Destination State
	AcceptDate   string         `xml:"AcceptDate"`  //
	AcceptTime   string         `xml:"AcceptTime"`
	Weight       string         `xml:"Weight,omitempty"`
	Expedited    Expedited      `xml:"Expedited,omitempty"`
	NonExpedited []NonExpedited `xml:"NonExpedited,omitempty"`
}

type Expedited struct {
	EAD               string                `xml:"EAD,omitempty"`
	Commitment        []ExpeditedCommitment `xml:"Commitment,omitempty"`
	ExpeditedMessage  string                `xml:"ExpeditedMessage,omitempty"`
	ExpeditedTransMsg []TransMsg            `xml:"ExpeditedTransMsg,omitempty"`
}

type ExpeditedCommitment struct {
	MailClass      string              `xml:"MailClass,omitempty"`
	CommitmentName string              `xml:"CommitmentName,omitempty"`
	CommitmentTime string              `xml:"CommitmentTime,omitempty"`
	CommitmentSeq  string              `xml:"CommitmentSeq,omitempty"`
	Locations      []ExpeditedLocation `xml:"Location,omitempty"`
}

type ExpeditedLocation struct {
	SDD          string `xml:"SDD,omitempty"`
	COT          string `xml:"COT,omitempty"`
	FacType      string `xml:"FacType,omitempty"`
	Street       string `xml:"Street,omitempty"`
	City         string `xml:"City,omitempty"`
	State        string `xml:"State,omitempty"`
	ZIP          string `xml:"ZIP,omitempty"`
	IsGuaranteed string `xml:"IsGuaranteed,omitempty"`
}

type TransMsg struct {
	MsgCode string `xml:"MsgCode,omitempty"`
	Msg     string `xml:"Msg,omitempty"`
}

type NonExpedited struct {
	MailClass                  string   `xml:"MailClass,omitempty"`
	NonExpeditedDestType       string   `xml:"NonExpeditedDestType,omitempty"`
	EAD                        string   `xml:"EAD,omitempty"`
	SvcStdDays                 string   `xml:"SvcStdDays,omitempty"`
	SchedDlvryDate             string   `xml:"SchedDlvryDate,omitempty"`
	HFPUEAD                    string   `xml:"HFPU>EAD,omitempty"`
	COT                        string   `xml:"HFPU>COT,omitempty"`
	ServiceStanardMessage      string   `xml:"HFPU>ServiceStandard>SvcStdMsg,omitempty"`
	ServiceStandardDays        string   `xml:"HFPU>ServiceStandard>SvcStdDays,omitempty"`
	ScheduledDeliveryDate      string   `xml:"HFPU>ServiceStandard>Location>SchedDlvryDate,omitempty"`
	NonDeliveryDays            string   `xml:"HFPU>ServiceStandard>Location>NonDeliveryDays,omitempty"`
	RAUName                    string   `xml:"HFPU>ServiceStandard>Location>RAUName,omitempty"`
	Street                     string   `xml:"HFPU>ServiceStandard>Location>Street,omitempty"`
	ZIP                        string   `xml:"HFPU>ServiceStandard>Location>ZIP,omitempty"`
	City                       string   `xml:"HFPU>ServiceStandard>Location>City,omitempty"`
	State                      string   `xml:"HFPU>ServiceStandard>Location>State,omitempty"`
	MondayCloseTime            string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>M,omitempty"`
	TuesdayCloseTime           string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>Tu,omitempty"`
	WednesdayCloseTime         string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>W,omitempty"`
	ThursdayCloseTime          string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>Th,omitempty"`
	FridayCloseTime            string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>F,omitempty"`
	SaturdayCloseTime          string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>Sa,omitempty"`
	SundayCloseTime            string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>Su,omitempty"`
	HolidayCloseTime           string   `xml:"HFPU>ServiceStandard>Location>CloseTimes>H,omitempty"`
	NonExpeditedMsg            string   `xml:"HFPU>ServiceStandard>Location>NonExpeditedExceptions>NonExpeditedMsg,omitempty"`
	HFPUNonExpeditedTransMsg   TransMsg `xml:"HFPU>ServiceStandard>Location>NonExpeditedExceptions>NonExpeditedTransMsg,omitempty"`
	PostCOT                    string   `xml:"HFPU>HFPUGlobalExcept>PostCOT,omitempty"`
	OverMaxResults             string   `xml:"HFPU>HFPUGlobalExcept>OverMaxResults,omitempty"`
	NoHFPULocInd               string   `xml:"HFPU>HFPUGlobalExcept>NoHFPULocInd,omitempty"`
	NonExpeditedWTMsg          []string `xml:"HFPU>HFPUGlobalExcept>NonExpeditedWTMsg,omitempty"`
	GlobalNonExpeditedTransMsg TransMsg `xml:"HFPU>HFPUGlobalExcept>NonExpeditedTransMsg,omitempty"`
}