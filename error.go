package usps

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

// Example:
/*
<Error>
	<Number>80040B19</Number>
	<Description>XML Syntax Error: Please check the XML request to see if it can be parsed.</Description>
	<Source>USPSCOM::DoAuth</Source>
</Error>
*/