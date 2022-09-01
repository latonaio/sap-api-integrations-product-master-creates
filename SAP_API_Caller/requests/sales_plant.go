package requests

type SalesPlant struct {
	Product               string  `json:"Product"`
	Plant                 string  `json:"Plant"`
	LoadingGroup          *string `json:"LoadingGroup,omitempty"`
	AvailabilityCheckType *string `json:"AvailabilityCheckType,omitempty"`
}
