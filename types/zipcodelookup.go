package types

type ZipCodeLookupRequest struct {
	USERID  string     `xml:"USERID,attr"` // Your Web Tools ID.
	Address []*Address `xml:"Address"`     // Up to 5 address verifications can be included per transaction.
}

