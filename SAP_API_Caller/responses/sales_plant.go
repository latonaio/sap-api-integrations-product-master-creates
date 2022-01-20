package responses

type SalesPlant struct {
	D struct {
		Product               string `json:"Product"`
		Plant                 string `json:"Plant"`
		LoadingGroup          string `json:"LoadingGroup"`
		AvailabilityCheckType string `json:"AvailabilityCheckType"`
	} `json:"d"`
}
