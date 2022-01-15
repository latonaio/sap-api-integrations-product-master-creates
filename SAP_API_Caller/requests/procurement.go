package requests

type Procurement struct {
 	Product                     string `json:"Product"`
 	Plant                       string `json:"Plant"`
 	IsAutoPurOrdCreationAllowed *bool  `json:"IsAutoPurOrdCreationAllowed"`
 	IsSourceListRequired        *bool  `json:"IsSourceListRequired"`
}
