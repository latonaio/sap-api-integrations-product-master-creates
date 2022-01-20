package responses

type Procurement struct {
	D struct {
		Product                     string `json:"Product"`
		Plant                       string `json:"Plant"`
		IsAutoPurOrdCreationAllowed bool   `json:"IsAutoPurOrdCreationAllowed"`
		IsSourceListRequired        bool   `json:"IsSourceListRequired"`
	} `json:"d"`
}
