package types

func (r *HFPFacilityInfoRequest) ToHTTP() (string, error) {
	return createRequest("HFPFacilityInfo", r)
}

type HFPFacilityInfoRequest struct {
	Request
	PickupCity  string `xml:"PickupCity"`
	PickupState string `xml:"PickupState"`
	PickupZIP   string `xml:"PickupZIP"`
	PickupZIP4  string `xml:"PickupZIP4"`
	Service     string `xml:"Service,omitempty"`
}

type HFPFacilityInfoResponse struct {
	*HFPFacilityInfoRequest
	USERID     string `xml:"-"`
	Facility   []*Facility
	LogMessage string `xml:"LogMessage,omitempty"`
}

type Facility struct {
	FacilityID        string `xml:"FacilityID"`
	FacilityName      string `xml:"FacilityName"`
	FacilityAddress   string `xml:"FacilityAddress"`
	FacilityCity      string `xml:"FacilityCity"`
	FacilityState     string `xml:"FacilityState"`
	FacilityZIP       string `xml:"FacilityZIP"`
	FacilityZIP4      string `xml:"FacilityZIP4"`
	Has10amCommitment string `xml:"Has10amCommitment"`
}
