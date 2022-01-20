package responses

type Accounting struct {
	D struct {
		Product             string `json:"Product"`
		ValuationArea       string `json:"ValuationArea"`
		ValuationClass      string `json:"ValuationClass"`
		StandardPrice       string `json:"StandardPrice"`
		PriceUnitQty        string `json:"PriceUnitQty"`
		MovingAveragePrice  string `json:"MovingAveragePrice"`
		PriceLastChangeDate string `json:"PriceLastChangeDate"`
		PlannedPrice        string `json:"PlannedPrice"`
		IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
	} `json:"d"`
}
