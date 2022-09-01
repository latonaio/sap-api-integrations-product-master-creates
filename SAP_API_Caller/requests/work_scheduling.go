package requests

type WorkScheduling struct {
	Product                       string  `json:"Product"`
	Plant                         string  `json:"Plant"`
	ProductionInvtryManagedLoc    *string `json:"ProductionInvtryManagedLoc,omitempty"`
	ProductProcessingTime         *string `json:"ProductProcessingTime,omitempty"`
	ProductionSupervisor          *string `json:"ProductionSupervisor,omitempty"`
	ProductProductionQuantityUnit *string `json:"ProductProductionQuantityUnit,omitempty"`
	ProdnOrderIsBatchRequired     *string `json:"ProdnOrderIsBatchRequired,omitempty"`
	MatlCompIsMarkedForBackflush  *string `json:"MatlCompIsMarkedForBackflush,omitempty"`
	ProductionSchedulingProfile   *string `json:"ProductionSchedulingProfile,omitempty"`
}
