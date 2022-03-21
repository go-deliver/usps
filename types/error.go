package types

// Error is the response returned from USPS API when there is an issue
// with the request or other reasons.
type Error struct {
	Number      string `xml:"Number"`
	Description string `xml:"Description,omitempty"`
	Source      string `xml:"Source,omitempty"`
	HelpFile    string `xml:"HelpFile,omitempty"`
	HelpContext string `xml:"HelpContext,omitempty"`
}

// Error extends the default error with USPS Errors.
func (e *Error) Error() string {
	return e.Description
}
