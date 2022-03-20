package response

// Error is returned from USPS API
type Error struct {
	Number      string `xml:"Number"`
	Description string `xml:"Description,omitempty"`
	Source      string `xml:"Source,omitempty"`
	HelpFile    string `xml:"HelpFile,omitempty"`
	HelpContext string `xml:"HelpContext,omitempty"`
}

func (e *Error) Error() string {
	return e.Description
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