package responses

type MRPArea struct {
	D struct {
		Product                       string `json:"Product"`
		Plant                         string `json:"Plant"`
		MRPArea                       string `json:"MRPArea"`
		MRPType                       string `json:"MRPType"`
		MRPResponsible                string `json:"MRPResponsible"`
		MRPGroup                      string `json:"MRPGroup"`
		ReorderThresholdQuantity      string `json:"ReorderThresholdQuantity"`
		PlanningTimeFence             string `json:"PlanningTimeFence"`
		LotSizingProcedure            string `json:"LotSizingProcedure"`
		LotSizeRoundingQuantity       string `json:"LotSizeRoundingQuantity"`
		MinimumLotSizeQuantity        string `json:"MinimumLotSizeQuantity"`
		MaximumLotSizeQuantity        string `json:"MaximumLotSizeQuantity"`
		MaximumStockQuantity          string `json:"MaximumStockQuantity"`
		ProcurementSubType            string `json:"ProcurementSubType"`
		DfltStorageLocationExtProcmt  string `json:"DfltStorageLocationExtProcmt"`
		MRPPlanningCalendar           string `json:"MRPPlanningCalendar"`
		SafetyStockQuantity           string `json:"SafetyStockQuantity"`
		SafetyDuration                string `json:"SafetyDuration"`
		FixedLotSizeQuantity          string `json:"FixedLotSizeQuantity"`
		PlannedDeliveryDurationInDays string `json:"PlannedDeliveryDurationInDays"`
		StorageLocation               string `json:"StorageLocation"`
		IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
	} `json:"d"`
}
