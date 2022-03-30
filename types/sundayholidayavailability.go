package types

func (r *SundayHolidayRequest) ToHTTP() (string, error) {
	return createRequest("SundayHolidayAvailability", r)
}

type SundayHolidayRequest struct {
	request
	SundayHoliday string `xml:"SundayHoliday"` // Select Sunday or Holiday
	FromZipCode   string `xml:"FromZipCode"` // 5 digit zip code from ZipCode
	ToZipCode     string `xml:"ToZipCode"` // 5 digit zip code to ZipCode
}

type SundayHolidayResponse struct {
	DeliveryAvailability string `xml:"DeliveryAvailability"` // Returns whether the lane Is eligible for delivery on day type selected.
	Error string `xml:"Error,omitempty"` // Returns error message if any.
}