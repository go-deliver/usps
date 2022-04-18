package types

func (r *CityStateLookupRequest) API() string {
	return "CityStateLookup"
}

type CityStateLookupRequest struct {
	request
	ZIPCode []*Zip `xml:"ZipCode"`
}

type Zip struct {
	ID   string `xml:"ID,attr,omitempty"`
	ZIP5 string `xml:"Zip5"`
}

type ZipResponse struct {
	*Zip
	City  string `xml:"City"`
	State string `xml:"State"`
}

type CityStateLookupResponse struct {
	response
	ZipResponse []*ZipResponse `xml:"ZipCode"`
}
