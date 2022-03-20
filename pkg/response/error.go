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