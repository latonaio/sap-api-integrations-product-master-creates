package requests

type SalesPlant struct {
	Product               string  `json:"Product"`
	Plant                 string  `json:"Plant"`
	LoadingGroup          *string `json:"LoadingGroup"`
	AvailabilityCheckType *string `json:"AvailabilityCheckType"`
}
