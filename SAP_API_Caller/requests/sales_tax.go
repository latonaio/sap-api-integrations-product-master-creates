package requests

type SalesTax struct {
	Product           string  `json:"Product"`
	Country           string  `json:"Country"`
	TaxCategory       string  `json:"TaxCategory"`
	TaxClassification *string `json:"TaxClassification,omitempty"`
}
