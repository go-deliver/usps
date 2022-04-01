package types

// Shared between CarrierPickupChange and CarrierPickupAvailability

type CarrierPickupPackage struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}