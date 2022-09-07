package requests

type MRPArea struct {
	Product                       string  `json:"Product"`
	Plant                         string  `json:"Plant"`
	MRPArea                       string  `json:"MRPArea"`
	MRPType                       *string `json:"MRPType,omitempty"`
	MRPResponsible                *string `json:"MRPResponsible,omitempty"`
	MRPGroup                      *string `json:"MRPGroup,omitempty"`
	ReorderThresholdQuantity      *string `json:"ReorderThresholdQuantity,omitempty"`
	PlanningTimeFence             *string `json:"PlanningTimeFence,omitempty"`
	LotSizingProcedure            *string `json:"LotSizingProcedure,omitempty"`
	LotSizeRoundingQuantity       *string `json:"LotSizeRoundingQuantity,omitempty"`
	MinimumLotSizeQuantity        *string `json:"MinimumLotSizeQuantity,omitempty"`
	MaximumLotSizeQuantity        *string `json:"MaximumLotSizeQuantity,omitempty"`
	MaximumStockQuantity          *string `json:"MaximumStockQuantity,omitempty"`
	ProcurementSubType            *string `json:"ProcurementSubType,omitempty"`
	DfltStorageLocationExtProcmt  *string `json:"DfltStorageLocationExtProcmt,omitempty"`
	MRPPlanningCalendar           *string `json:"MRPPlanningCalendar,omitempty"`
	SafetyStockQuantity           *string `json:"SafetyStockQuantity,omitempty"`
	SafetyDuration                *string `json:"SafetyDuration,omitempty"`
	FixedLotSizeQuantity          *string `json:"FixedLotSizeQuantity,omitempty"`
	PlannedDeliveryDurationInDays *string `json:"PlannedDeliveryDurationInDays,omitempty"`
	StorageLocation               *string `json:"StorageLocation,omitempty"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion,omitempty"`
}
