package requests

type Procurement struct {
	Product                     string `json:"Product"`
	Plant                       string `json:"Plant"`
	IsAutoPurOrdCreationAllowed *bool  `json:"IsAutoPurOrdCreationAllowed,omitempty"`
	IsSourceListRequired        *bool  `json:"IsSourceListRequired,omitempty"`
}
