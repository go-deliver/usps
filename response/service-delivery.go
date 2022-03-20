package response

// SDCGetLocationsResponse is the expected response returned when sending SDCGetLocations request to USPS
type SDCGetLocationsResponse struct {
	Release                string          `xml:"Release"`                          // Release version
	CallerID               string          `xml:"CallerID"`                         // Caller ID
	SourceID               string          `xml:"SourceID"`                         // Source ID
	MailClass              string          `xml:"MailClass"`                        // MailClass from request
	OriginZIP              string          `xml:"OriginZIP"`                        // OriginZIP from request
	OriginCity             string          `xml:"OriginCity"`                       // Origin City
	OriginState            string          `xml:"OriginState"`                      // Origin State
	DestZIP                string          `xml:"DestZIP"`                          // Destination ZIP Code
	DestCity               string          `xml:"DestCity"`                         // Destination City
	DestState              string          `xml:"DestState"`                        // Destination State
	AcceptDate             string          `xml:"AcceptDate"`                       // Acceptance date at postal facility
	AcceptTime             string          `xml:"AcceptTime"`                       // Acceptance time at postal facility
	Weight                 string          `xml:"Weight,omitempty"`                 // Weight from request
	NonExpeditedOriginType string          `xml:"NonExpeditedOriginType,omitempty"` // Nonexpedited origin type from request
	Expedited              *expedited      `xml:"Expedited,omitempty"`              // Groups Priority Mail Express Mail and Priority Mail Commitments and Effective Acceptance Dates
	NonExpedited           []*nonExpedited `xml:"NonExpedited,omitempty"`           // Non-Expedited Mail Class elements
}

type expedited struct {
	EAD               string                 `xml:"EAD,omitempty"`               // Effective Acceptance Date
	Commitment        []*expeditedCommitment `xml:"Commitment,omitempty"`        // Priority Mail Express and Priority Mail commitment information
	ExpeditedMessage  string                 `xml:"ExpeditedMessage,omitempty"`  // Expedited messaging
	ExpeditedTransMsg []*transMsg            `xml:"ExpeditedTransMsg,omitempty"` // Expedited transportational messaging
}

type expeditedCommitment struct {
	MailClass string `xml:"MailClass,omitempty"` // Mail Class
	/*
		Commitment Name.

		Valid: "1-Day",	"2-Day", "3-Day", "DPO", "Military"

		Note: If a scenario exceeds a 3-day commitment then no value will be returned in the response field
	*/
	CommitmentName string `xml:"CommitmentName,omitempty"`
	CommitmentTime string `xml:"CommitmentTime,omitempty"` // Commitment Time. Value will only be returned for Priority Mail
	/*
		Commitment Sequence.

		Valid:

		"A0118", "B0118" = 1-Day at 6:00 PM

		"A0218", "B0218" = 2-Day at 6:00 PM
	*/
	CommitmentSeq string               `xml:"CommitmentSeq,omitempty"`
	Locations     []*expeditedLocation `xml:"Location,omitempty"` // Drop off location information
}

type expeditedLocation struct {
	SDD     string `xml:"SDD,omitempty"`     // Scheduled Delivery Date
	COT     string `xml:"COT,omitempty"`     // Cut-off Time
	FacType string `xml:"FacType,omitempty"` // Facility Type
	Street  string `xml:"Street,omitempty"`  // Facility Street
	City    string `xml:"City,omitempty"`    // Facility City
	State   string `xml:"State,omitempty"`   // Facility State
	ZIP     string `xml:"ZIP,omitempty"`     // Facility Zip Code
	/*
		Indicates if Guarantee is offered for Priority Mail Express.

		Valid:

		“1” = Guaranteed

		“2” = No Guarantee

		“3” = Temporary No Guarantee
	*/
	IsGuaranteed string `xml:"IsGuaranteed,omitempty"`
}

type nonExpedited struct {
	MailClass string `xml:"MailClass,omitempty"` // Mail Class
	/*
		Destination Type

		Valid:

		"1" = PO-Addressee - Street

		"2" = Destination - PO Box

		"3" = Hold for pickup (HFPU)
	*/
	NonExpeditedDestType   string                  `xml:"NonExpeditedDestType,omitempty"`
	EAD                    string                  `xml:"EAD,omitempty"`                    // Effective Acceptance Date
	SvcStdDays             string                  `xml:"SvcStdDays,omitempty"`             // Service Standard Days
	SchedDlvryDate         string                  `xml:"SchedDlvryDate,omitempty"`         // Scheduled Delivery Date
	HFPU                   *hfpu                   `xml:"HFPU,omitempty"`                   // Hold for Pickup information
	NonExpeditedExceptions *nonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"` // Nonexpedited Exception Elements
}

type nonExpeditedExceptions struct {
	NonExpeditedMsg      string    `xml:"NonExpeditedMsg,omitempty"`      // Additional exception messaging
	NonExpeditedTransMsg *transMsg `xml:"NonExpeditedTransMsg,omitempty"` // Nonexpedited transportational messaging
}

type hfpu struct {
	EAD             string           `xml:"EAD,omitempty"`              // Effective Acceptance Date
	COT             string           `xml:"COT,omitempty"`              // Cut-off time
	ServiceStandard *serviceStandard `xml:"ServiceStandard,omitempty"`  // Hold for Pickup Service Standard Elements
	GlobalExcept    *globalExcept    `xml:"HFPUGlobalExcept,omitempty"` // Hold for Pickup Exception Elements
}

type globalExcept struct {
	PostCOT                    string    `xml:"PostCOT,omitempty"`              // Indicates if the mailpiece arrived after the Cut-Off Time
	OverMaxResults             string    `xml:"OverMaxResults,omitempty"`       // Indicates if more than the number of returned Hold for Pickup locations was found
	NoHFPULocInd               string    `xml:"NoHFPULocInd,omitempty"`         // Indicates if there were no locations found in the RAU that support Hold for Pickup
	NonExpeditedWTMsg          []string  `xml:"NonExpeditedWTMsg,omitempty"`    // Additional exception messaging
	GlobalNonExpeditedTransMsg *transMsg `xml:"NonExpeditedTransMsg,omitempty"` // Non-Expedited Mail transportation messaging
}

type serviceStandard struct {
	ServiceStandardMessage string                `xml:"SvcStdMsg,omitempty"`  // Service Standard Message
	ServiceStandardDays    string                `xml:"SvcStdDays,omitempty"` // Service Standard Days
	Location               *nonExpeditedLocation `xml:"Location,omitempty"`   // Hold for Pickup Location information
}

type nonExpeditedLocation struct {
	ScheduledDeliveryDate string                  `xml:"SchedDlvryDate,omitempty"`         // Scheduled Delivery Date
	NonDeliveryDays       string                  `xml:"NonDeliveryDays,omitempty"`        // Non-Delivery Days
	RAUName               string                  `xml:"RAUName,omitempty"`                // RAU Name
	Street                string                  `xml:"Street,omitempty"`                 // RAU Street Address
	ZIP                   string                  `xml:"ZIP,omitempty"`                    // Facility ZIP Code
	City                  string                  `xml:"City,omitempty"`                   // RAU City
	State                 string                  `xml:"State,omitempty"`                  // RAU State.
	CloseTimes            *closeTimes             `xml:"CloseTimes,omitempty"`             // Close times by day of the week for Hold for Pickup location
	NonExpeditedMsg       *nonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"` // Additional exception messaging
}

type closeTimes struct {
	M  string `xml:"M,omitempty"`  // Close Time for Monday
	Tu string `xml:"Tu,omitempty"` // Close Time for Tuesday
	W  string `xml:"W,omitempty"`  // Close Time for Wednesday
	Th string `xml:"Th,omitempty"` // Close Time for Thursday
	F  string `xml:"F,omitempty"`  // Close Time for Friday
	Sa string `xml:"Sa,omitempty"` // Close Time for Saturday
	Su string `xml:"Su,omitempty"` // Close Time for Sunday
	H  string `xml:"H,omitempty"`  // Close Time for Holiday
}

type transMsg struct {
	MsgCode string `xml:"MsgCode,omitempty"` // Transportation messaging code
	Msg     string `xml:"Msg,omitempty"`     // Transportation messaging text
}