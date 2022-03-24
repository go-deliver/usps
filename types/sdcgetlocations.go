package types

func (r *SDCGetLocationsRequest) ToHTTP() (string, error) {
	return createRequest("SDCGetLocations", r)
}

// SDCGetLocationsRequest is the request sent to the USPS ServiceDeliveryGetLocations API
type SDCGetLocationsRequest struct {
	Request
	/*
		Defines mail class for commitment data.

		Valid:

		"0" = All Mail Classes

		"1" = Priority Mail Express

		"2" = Priority Mail

		"3" = First Class Mail

		"4" = Marketing Mail

		"5" = Periodicals

		"6" = Package Services
	*/
	MailClass string `xml:"MailClass"`
	// The origin ZIP code. May be 3, 5, or 9 characters.
	OriginZIP string `xml:"OriginZIP"`
	// The destination Zip code. Must be a valid 5-character zip code.
	DestinationZIP string `xml:"DestinationZIP"`
	/*
	 The date the package will be mailed.
	 Acceptance date may be up to 30 days in advance.
	 Defaults to system date.

	 Format: YYYY-MM-DD

	 Example: "2014-09-23"
	*/
	AcceptDate string `xml:"AcceptDate,omitempty"`
	/*
		Time package will be accepted at a postal facility.
		Default value returned when no value is included in the request will be current time.

		Format: HHMM

		Example: "1600"
	*/
	AcceptTime string `xml:"AcceptTime,omitempty"`
	/*
		Returns additional detail for Non-Expedited mail classes when "True".
		Default: "False"
		Example:
	*/
	NonEMDetail string `xml:"NonEMDetail,omitempty"`
	/*
		Origin type indicator for non-Expedited shipments.

		Valid:

		"1" = Local Mail

		"2" = Destination Entered Mail
	*/
	NonEMOriginType string `xml:"NonEMOriginType,omitempty"`
	/*
		Destination type indicator for non-Expedited shipments.

		Valid:

		"1" = PO-Addressee - Street

		"2" = Destination - PO Box

		"3" = Hold for pickup (HFPU)
	*/
	NonEMDestType string `xml:"NonEMDestType,omitempty"`
	// Item weight
	Weight string `xml:"Weight,omitempty"`
}

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
	Expedited              *Expedited      `xml:"Expedited,omitempty"`              // Groups Priority Mail Express Mail and Priority Mail Commitments and Effective Acceptance Dates
	NonExpedited           []*NonExpedited `xml:"NonExpedited,omitempty"`           // Non-Expedited Mail Class elements
}

type Expedited struct {
	EAD               string                 `xml:"EAD,omitempty"`               // Effective Acceptance Date
	Commitment        []*ExpeditedCommitment `xml:"Commitment,omitempty"`        // Priority Mail Express and Priority Mail commitment information
	ExpeditedMessage  string                 `xml:"ExpeditedMessage,omitempty"`  // Expedited messaging
	ExpeditedTransMsg []*TransMsg            `xml:"ExpeditedTransMsg,omitempty"` // Expedited transportational messaging
}

type ExpeditedCommitment struct {
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
	CommitmentSeq string              `xml:"CommitmentSeq,omitempty"`
	Locations     []ExpeditedLocation `xml:"Location,omitempty"` // Drop off location information
}

type ExpeditedLocation struct {
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

type NonExpedited struct {
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
	HFPU                   *HFPU                   `xml:"HFPU,omitempty"`                   // Hold for Pickup information
	NonExpeditedExceptions *NonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"` // Nonexpedited Exception Elements
}

type NonExpeditedExceptions struct {
	NonExpeditedMsg      string    `xml:"NonExpeditedMsg,omitempty"`      // Additional exception messaging
	NonExpeditedTransMsg *TransMsg `xml:"NonExpeditedTransMsg,omitempty"` // Nonexpedited transportational messaging
}

type HFPU struct {
	EAD             string           `xml:"EAD,omitempty"`              // Effective Acceptance Date
	COT             string           `xml:"COT,omitempty"`              // Cut-off time
	ServiceStandard *ServiceStandard `xml:"ServiceStandard,omitempty"`  // Hold for Pickup Service Standard Elements
	GlobalExcept    *GlobalExcept    `xml:"HFPUGlobalExcept,omitempty"` // Hold for Pickup Exception Elements
}

type GlobalExcept struct {
	PostCOT                    string    `xml:"PostCOT,omitempty"`              // Indicates if the mailpiece arrived after the Cut-Off Time
	OverMaxResults             string    `xml:"OverMaxResults,omitempty"`       // Indicates if more than the number of returned Hold for Pickup locations was found
	NoHFPULocInd               string    `xml:"NoHFPULocInd,omitempty"`         // Indicates if there were no locations found in the RAU that support Hold for Pickup
	NonExpeditedWTMsg          []string  `xml:"NonExpeditedWTMsg,omitempty"`    // Additional exception messaging
	GlobalNonExpeditedTransMsg *TransMsg `xml:"NonExpeditedTransMsg,omitempty"` // Non-Expedited Mail transportation messaging
}

type ServiceStandard struct {
	ServiceStandardMessage string                `xml:"SvcStdMsg,omitempty"`  // Service Standard Message
	ServiceStandardDays    string                `xml:"SvcStdDays,omitempty"` // Service Standard Days
	Location               *NonExpeditedLocation `xml:"Location,omitempty"`   // Hold for Pickup Location information
}

type NonExpeditedLocation struct {
	ScheduledDeliveryDate string                  `xml:"SchedDlvryDate,omitempty"`         // Scheduled Delivery Date
	NonDeliveryDays       string                  `xml:"NonDeliveryDays,omitempty"`        // Non-Delivery Days
	RAUName               string                  `xml:"RAUName,omitempty"`                // RAU Name
	Street                string                  `xml:"Street,omitempty"`                 // RAU Street Address
	ZIP                   string                  `xml:"ZIP,omitempty"`                    // Facility ZIP Code
	City                  string                  `xml:"City,omitempty"`                   // RAU City
	State                 string                  `xml:"State,omitempty"`                  // RAU State.
	CloseTimes            *CloseTimes             `xml:"CloseTimes,omitempty"`             // Close times by day of the week for Hold for Pickup location
	NonExpeditedMsg       *NonExpeditedExceptions `xml:"NonExpeditedExceptions,omitempty"` // Additional exception messaging
}

type CloseTimes struct {
	M  string `xml:"M,omitempty"`  // Close Time for Monday
	Tu string `xml:"Tu,omitempty"` // Close Time for Tuesday
	W  string `xml:"W,omitempty"`  // Close Time for Wednesday
	Th string `xml:"Th,omitempty"` // Close Time for Thursday
	F  string `xml:"F,omitempty"`  // Close Time for Friday
	Sa string `xml:"Sa,omitempty"` // Close Time for Saturday
	Su string `xml:"Su,omitempty"` // Close Time for Sunday
	H  string `xml:"H,omitempty"`  // Close Time for Holiday
}

type TransMsg struct {
	MsgCode string `xml:"MsgCode,omitempty"` // Transportation messaging code
	Msg     string `xml:"Msg,omitempty"`     // Transportation messaging text
}
