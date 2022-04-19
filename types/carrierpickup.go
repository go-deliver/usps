package types

// Shared between CarrierPickup APIs

type CarrierPickupPackage struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}