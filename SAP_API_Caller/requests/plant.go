package requests

type Plant struct {
	Product                       string  `json:"Product"`
	Plant                         string  `json:"Plant"`
	PurchasingGroup               *string `json:"PurchasingGroup,omitempty"`
	ProductionInvtryManagedLoc    *string `json:"ProductionInvtryManagedLoc,omitempty"`
	AvailabilityCheckType         *string `json:"AvailabilityCheckType,omitempty"`
	ProfitCenter                  *string `json:"ProfitCenter,omitempty"`
	GoodsReceiptDuration          *string `json:"GoodsReceiptDuration,omitempty"`
	MRPType                       *string `json:"MRPType,omitempty"`
	MRPResponsible                *string `json:"MRPResponsible,omitempty"`
	MinimumLotSizeQuantity        *string `json:"MinimumLotSizeQuantity,omitempty"`
	MaximumLotSizeQuantity        *string `json:"MaximumLotSizeQuantity,omitempty"`
	FixedLotSizeQuantity          *string `json:"FixedLotSizeQuantity,omitempty"`
	IsBatchManagementRequired     *bool   `json:"IsBatchManagementRequired,omitempty"`
	ProcurementType               *string `json:"ProcurementType,omitempty"`
	IsInternalBatchManaged        *bool   `json:"IsInternalBatchManaged,omitempty"`
	GoodsIssueUnit                *string `json:"GoodsIssueUnit,omitempty"`
	MaterialFreightGroup          *string `json:"MaterialFreightGroup,omitempty"`
	ProductLogisticsHandlingGroup *string `json:"ProductLogisticsHandlingGroup,omitempty"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion,omitempty"`
}
