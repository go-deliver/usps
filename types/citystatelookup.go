package types

func (r *CityStateLookupRequest) ToHTTP() (string, error) {
	return createRequest("CityStateLookup", r)
}

type Zip struct {
	ID   string `xml:"ID,attr,omitempty"`
	ZIP5 string `xml:"Zip5"`
}

type CityStateLookupRequest struct {
	Request
	ZIPCode []*Zip `xml:"ZipCode"`
}

type ZipResponse struct {
	*Zip
	City  string `xml:"City"`
	State string `xml:"State"`
}

type CityStateLookupResponse struct {
	ZipResponse []*ZipResponse `xml:"ZipCode"`
}
