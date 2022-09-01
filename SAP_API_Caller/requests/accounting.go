package requests

type Accounting struct {
	Product             string  `json:"Product"`
	ValuationArea       string  `json:"ValuationArea"`
	ValuationClass      *string `json:"ValuationClass,omitempty"`
	StandardPrice       *string `json:"StandardPrice,omitempty"`
	PriceUnitQty        *string `json:"PriceUnitQty,omitempty"`
	MovingAveragePrice  *string `json:"MovingAveragePrice,omitempty"`
	PriceLastChangeDate *string `json:"PriceLastChangeDate,omitempty"`
	PlannedPrice        *string `json:"PlannedPrice,omitempty"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion,omitempty"`
}
