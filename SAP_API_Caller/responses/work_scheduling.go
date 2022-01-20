package responses

type WorkScheduling struct {
	D struct {
		Product                       string `json:"Product"`
		Plant                         string `json:"Plant"`
		ProductionInvtryManagedLoc    string `json:"ProductionInvtryManagedLoc"`
		ProductProcessingTime         string `json:"ProductProcessingTime"`
		ProductionSupervisor          string `json:"ProductionSupervisor"`
		ProductProductionQuantityUnit string `json:"ProductProductionQuantityUnit"`
		ProdnOrderIsBatchRequired     string `json:"ProdnOrderIsBatchRequired"`
		MatlCompIsMarkedForBackflush  string `json:"MatlCompIsMarkedForBackflush"`
		ProductionSchedulingProfile   string `json:"ProductionSchedulingProfile"`
	} `json:"d"`
}
