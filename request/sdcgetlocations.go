package request

func (r *SDCGetLocationsRequest) ToHTTP() (string, error) {
	return createRequest("SDCGetLocations", r)
}

// SDCGetLocationsRequest is the request sent to the USPS ServiceDeliveryGetLocations API
type SDCGetLocationsRequest struct {
	// Your WebTools ID.
	USERID          string `xml:"USERID,attr"`
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
	MailClass       string `xml:"MailClass"`
	 // The origin ZIP code. May be 3, 5, or 9 characters.
	OriginZIP       string `xml:"OriginZIP"`
	// The destination Zip code. Must be a valid 5-character zip code.
	DestinationZIP  string `xml:"DestinationZIP"`
	/*
	 The date the package will be mailed.
	 Acceptance date may be up to 30 days in advance.
	 Defaults to system date.
	 
	 Format: YYYY-MM-DD
	 
	 Example: "2014-09-23"
	*/
	AcceptDate      string `xml:"AcceptDate,omitempty"`
	/*
	Time package will be accepted at a postal facility.
	Default value returned when no value is included in the request will be current time.

	Format: HHMM

	Example: "1600"
	*/
	AcceptTime      string `xml:"AcceptTime,omitempty"`
	/*
	Returns additional detail for Non-Expedited mail classes when "True".
	Default: "False"
	Example:
	*/
	NonEMDetail     string `xml:"NonEMDetail,omitempty"`
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
	NonEMDestType   string `xml:"NonEMDestType,omitempty"`
	// Item weight
	Weight          string `xml:"Weight,omitempty"`
}
