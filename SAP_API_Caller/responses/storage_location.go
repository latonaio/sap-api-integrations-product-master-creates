package responses

type StorageLocation struct {
	D struct {
	    Product                        string `json:"Product"`
	    Plant                          string `json:"Plant"`
	    StorageLocation                string `json:"StorageLocation"`
	    WarehouseStorageBin            string `json:"WarehouseStorageBin"`
	    MaintenanceStatus              string `json:"MaintenanceStatus"`
	    PhysicalInventoryBlockInd      string `json:"PhysicalInventoryBlockInd"`
	    CreationDate                   string `json:"CreationDate"`
	    IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
	    DateOfLastPostedCntUnRstrcdStk string `json:"DateOfLastPostedCntUnRstrcdStk"`
	    InventoryCorrectionFactor      string `json:"InventoryCorrectionFactor"`
	    InvtryRestrictedUseStockInd    string `json:"InvtryRestrictedUseStockInd"`
	    InvtryCurrentYearStockInd      string `json:"InvtryCurrentYearStockInd"`
	    InvtryQualInspCurrentYrStkInd  string `json:"InvtryQualInspCurrentYrStkInd"`
	    InventoryBlockStockInd         string `json:"InventoryBlockStockInd"`
	    InvtryRestStockPrevPeriodInd   string `json:"InvtryRestStockPrevPeriodInd"`
	    InventoryStockPrevPeriod       string `json:"InventoryStockPrevPeriod"`
	    InvtryStockQltyInspPrevPeriod  string `json:"InvtryStockQltyInspPrevPeriod"`
	    HasInvtryBlockStockPrevPeriod  string `json:"HasInvtryBlockStockPrevPeriod"`
	    FiscalYearCurrentPeriod        string `json:"FiscalYearCurrentPeriod"`
	    FiscalMonthCurrentPeriod       string `json:"FiscalMonthCurrentPeriod"`
	    FiscalYearCurrentInvtryPeriod  string `json:"FiscalYearCurrentInvtryPeriod"`
	    LeanWrhsManagementPickingArea  string `json:"LeanWrhsManagementPickingArea"`
	} `json:"d"`
}
