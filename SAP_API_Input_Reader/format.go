package sap_api_input_reader

import (
	"sap-api-integrations-product-master-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	return &requests.General{
		Product:             sdc.Product.Product,
		ProductType:         sdc.Product.ProductType,
		IsMarkedForDeletion: sdc.Product.IsMarkedForDeletion,
		GrossWeight:         sdc.Product.GrossWeight,
		WeightUnit:          sdc.Product.WeightUnit,
		ProductGroup:        sdc.Product.ProductGroup,
		BaseUnit:            sdc.Product.BaseUnit,
		Division:            sdc.Product.Division,
		ValidityStartDate:   sdc.Product.ValidityStartDate,
		SizeOrDimensionText: sdc.Product.SizeOrDimensionText,
		ProductStandardID:   sdc.Product.ProductStandardID,
		IndustrySector:      sdc.Product.IndustrySector,
		ToProductDesc:       sdc.ConvertToToProductDesc(),
	}
}

func (sdc *SDC) ConvertToToProductDesc() *requests.ToProductDesc {
	return &requests.ToProductDesc{
		ToProductDescResults: []requests.ToProductDescResults{
			{
				Product:            sdc.Product.Product,
				Language:           sdc.Product.ProductDescription.Language,
				ProductDescription: sdc.Product.ProductDescription.ProductDescription,
			},
		},
	}
}

func (sdc *SDC) ConvertToPlant() *requests.Plant {
	return &requests.Plant{
		Product:                       sdc.Product.Product,
		Plant:                         sdc.Product.Plant.Plant,
		PurchasingGroup:               sdc.Product.Plant.PurchasingGroup,
		ProductionInvtryManagedLoc:    sdc.Product.Plant.ProductionInvtryManagedLoc,
		AvailabilityCheckType:         sdc.Product.Plant.AvailabilityCheckType,
		ProfitCenter:                  sdc.Product.Plant.ProfitCenter,
		GoodsReceiptDuration:          sdc.Product.Plant.GoodsReceiptDuration,
		MRPType:                       sdc.Product.Plant.MRPType,
		MRPResponsible:                sdc.Product.Plant.MRPResponsible,
		MinimumLotSizeQuantity:        sdc.Product.Plant.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        sdc.Product.Plant.MaximumLotSizeQuantity,
		FixedLotSizeQuantity:          sdc.Product.Plant.FixedLotSizeQuantity,
		IsBatchManagementRequired:     sdc.Product.Plant.IsBatchManagementRequired,
		ProcurementType:               sdc.Product.Plant.ProcurementType,
		IsInternalBatchManaged:        sdc.Product.Plant.IsInternalBatchManaged,
		GoodsIssueUnit:                sdc.Product.Plant.GoodsIssueUnit,
		MaterialFreightGroup:          sdc.Product.Plant.MaterialFreightGroup,
		ProductLogisticsHandlingGroup: sdc.Product.Plant.ProductLogisticsHandlingGroup,
		IsMarkedForDeletion:           sdc.Product.Plant.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	return &requests.MRPArea{
		Product:                       sdc.Product.Product,
		Plant:                         sdc.Product.Plant.Plant,
		MRPArea:                       sdc.Product.Plant.MRPArea.MRPArea,
		MRPType:                       sdc.Product.Plant.MRPArea.MRPType,
		MRPResponsible:                sdc.Product.Plant.MRPArea.MRPResponsible,
		MRPGroup:                      sdc.Product.Plant.MRPArea.MRPGroup,
		ReorderThresholdQuantity:      sdc.Product.Plant.MRPArea.ReorderThresholdQuantity,
		PlanningTimeFence:             sdc.Product.Plant.MRPArea.PlanningTimeFence,
		LotSizingProcedure:            sdc.Product.Plant.MRPArea.LotSizingProcedure,
		LotSizeRoundingQuantity:       sdc.Product.Plant.MRPArea.LotSizeRoundingQuantity,
		MinimumLotSizeQuantity:        sdc.Product.Plant.MRPArea.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        sdc.Product.Plant.MRPArea.MaximumLotSizeQuantity,
		MaximumStockQuantity:          sdc.Product.Plant.MRPArea.MaximumStockQuantity,
		ProcurementSubType:            sdc.Product.Plant.MRPArea.ProcurementSubType,
		DfltStorageLocationExtProcmt:  sdc.Product.Plant.MRPArea.DfltStorageLocationExtProcmt,
		MRPPlanningCalendar:           sdc.Product.Plant.MRPArea.MRPPlanningCalendar,
		SafetyStockQuantity:           sdc.Product.Plant.MRPArea.SafetyStockQuantity,
		SafetyDuration:                sdc.Product.Plant.MRPArea.SafetyDuration,
		FixedLotSizeQuantity:          sdc.Product.Plant.MRPArea.FixedLotSizeQuantity,
		PlannedDeliveryDurationInDays: sdc.Product.Plant.MRPArea.PlannedDeliveryDurationInDays,
		StorageLocation:               sdc.Product.Plant.MRPArea.StorageLocation,
		IsMarkedForDeletion:           sdc.Product.Plant.MRPArea.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProcurement() *requests.Procurement {
	return &requests.Procurement{
		Product:                       sdc.Product.Product,
		Plant:                         sdc.Product.Plant.Plant,
		IsAutoPurOrdCreationAllowed:   sdc.Product.Plant.Procurement.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:          sdc.Product.Plant.Procurement.IsSourceListRequired,
	}
}

func (sdc *SDC) ConvertToWorkScheduling() *requests.WorkScheduling {
	return &requests.WorkScheduling{
		Product:                       sdc.Product.Product,
		Plant:                         sdc.Product.Plant.Plant,
		ProductionInvtryManagedLoc:    sdc.Product.Plant.WorkScheduling.ProductionInvtryManagedLoc,
		ProductProcessingTime:         sdc.Product.Plant.WorkScheduling.ProductProcessingTime,
		ProductionSupervisor:          sdc.Product.Plant.WorkScheduling.ProductionSupervisor,
		ProductProductionQuantityUnit: sdc.Product.Plant.WorkScheduling.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     sdc.Product.Plant.WorkScheduling.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  sdc.Product.Plant.WorkScheduling.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   sdc.Product.Plant.WorkScheduling.ProductionSchedulingProfile,
	}
}

func (sdc *SDC) ConvertToSalesPlant() *requests.SalesPlant {
	return &requests.SalesPlant{
		Product:                       sdc.Product.Product,
		Plant:                         sdc.Product.Plant.Plant,
		LoadingGroup:                  sdc.Product.Plant.SalesPlant.LoadingGroup,
		AvailabilityCheckType:         sdc.Product.Plant.SalesPlant.AvailabilityCheckType,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	return &requests.Accounting{
		Product:                       sdc.Product.Product,
		ValuationArea:                 sdc.Product.Accounting.ValuationArea,
		ValuationClass:                sdc.Product.Accounting.ValuationClass,
		StandardPrice:                 sdc.Product.Accounting.StandardPrice,
		PriceUnitQty:                  sdc.Product.Accounting.PriceUnitQty,
		MovingAveragePrice:            sdc.Product.Accounting.MovingAveragePrice,
		PriceLastChangeDate:           sdc.Product.Accounting.PriceLastChangeDate,
		PlannedPrice:                  sdc.Product.Accounting.PlannedPrice,
		IsMarkedForDeletion:           sdc.Product.Accounting.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToSalesOrganization() *requests.SalesOrganization {
	return &requests.SalesOrganization{
		Product:                        sdc.Product.Product,
		ProductSalesOrg:                sdc.Product.SalesOrganization.ProductSalesOrg,
		ProductDistributionChnl:        sdc.Product.SalesOrganization.ProductDistributionChnl,
		SupplyingPlant:                 sdc.Product.SalesOrganization.SupplyingPlant,
		PriceSpecificationProductGroup: sdc.Product.SalesOrganization.PriceSpecificationProductGroup,
		AccountDetnProductGroup:        sdc.Product.SalesOrganization.AccountDetnProductGroup,
		ItemCategoryGroup:              sdc.Product.SalesOrganization.ItemCategoryGroup,
		SalesMeasureUnit:               sdc.Product.SalesOrganization.SalesMeasureUnit,
		ProductHierarchy:               sdc.Product.SalesOrganization.ProductHierarchy,
		IsMarkedForDeletion:            sdc.Product.SalesOrganization.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProductDesc() *requests.ProductDesc {
	return &requests.ProductDesc{
		Product:                        sdc.Product.Product,
		Language:                       sdc.Product.ProductDescription.Language,
		ProductDescription:             sdc.Product.ProductDescription.ProductDescription,
	}
}

func (sdc *SDC) ConvertToQuality() *requests.Quality {
	return &requests.Quality{
		Product:                        sdc.Product.Product,
		Plant:                          sdc.Product.Plant.Plant,
		MaximumStoragePeriod:           sdc.Product.Plant.Quality.MaximumStoragePeriod,
		QualityMgmtCtrlKey:             sdc.Product.Plant.Quality.QualityMgmtCtrlKey,
		MatlQualityAuthorizationGroup:  sdc.Product.Plant.Quality.MatlQualityAuthorizationGroup,
		HasPostToInspectionStock:       sdc.Product.Plant.Quality.HasPostToInspectionStock,
		InspLotDocumentationIsRequired: sdc.Product.Plant.Quality.InspLotDocumentationIsRequired,
		SuplrQualityManagementSystem:   sdc.Product.Plant.Quality.SuplrQualityManagementSystem,
		RecrrgInspIntervalTimeInDays:   sdc.Product.Plant.Quality.RecrrgInspIntervalTimeInDays,
		ProductQualityCertificateType:  sdc.Product.Plant.Quality.ProductQualityCertificateType,
	}
}
