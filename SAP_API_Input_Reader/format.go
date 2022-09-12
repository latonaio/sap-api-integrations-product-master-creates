package sap_api_input_reader

import (
	"sap-api-integrations-product-master-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.Product
	return &requests.General{
		Product:             data.Product,
		ProductType:         data.ProductType,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
		GrossWeight:         data.GrossWeight,
		WeightUnit:          data.WeightUnit,
		ProductGroup:        data.ProductGroup,
		BaseUnit:            data.BaseUnit,
		NetWeight:           data.NetWeight,
		Division:            data.Division,
		ValidityStartDate:   data.ValidityStartDate,
		SizeOrDimensionText: data.SizeOrDimensionText,
		ProductStandardID:   data.ProductStandardID,
		IndustrySector:      data.IndustrySector,
		ToProductDesc: &struct {
			ToProductDescResults []*requests.ProductDesc `json:"results"`
		}{
			ToProductDescResults: []*requests.ProductDesc{
				sdc.ConvertToProductDesc(),
			},
		},
	}
}

func (sdc *SDC) ConvertToPlant() *requests.Plant {
	dataProduct := sdc.Product
	data := sdc.Product.Plant
	return &requests.Plant{
		Product:                       dataProduct.Product,
		Plant:                         data.Plant,
		PurchasingGroup:               data.PurchasingGroup,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		AvailabilityCheckType:         data.AvailabilityCheckType,
		ProfitCenter:                  data.ProfitCenter,
		GoodsReceiptDuration:          data.GoodsReceiptDuration,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		IsBatchManagementRequired:     data.IsBatchManagementRequired,
		ProcurementType:               data.ProcurementType,
		IsInternalBatchManaged:        data.IsInternalBatchManaged,
		GoodsIssueUnit:                data.GoodsIssueUnit,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		ProductLogisticsHandlingGroup: data.ProductLogisticsHandlingGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToMRPArea() *requests.MRPArea {
	dataProduct := sdc.Product
	dataPlant := sdc.Product.Plant
	data := sdc.Product.Plant.MRPArea
	return &requests.MRPArea{
		Product:                       dataProduct.Product,
		Plant:                         dataPlant.Plant,
		MRPArea:                       data.MRPArea,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MRPGroup:                      data.MRPGroup,
		ReorderThresholdQuantity:      data.ReorderThresholdQuantity,
		PlanningTimeFence:             data.PlanningTimeFence,
		LotSizingProcedure:            data.LotSizingProcedure,
		LotSizeRoundingQuantity:       data.LotSizeRoundingQuantity,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		MaximumStockQuantity:          data.MaximumStockQuantity,
		ProcurementSubType:            data.ProcurementSubType,
		DfltStorageLocationExtProcmt:  data.DfltStorageLocationExtProcmt,
		MRPPlanningCalendar:           data.MRPPlanningCalendar,
		SafetyStockQuantity:           data.SafetyStockQuantity,
		SafetyDuration:                data.SafetyDuration,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		PlannedDeliveryDurationInDays: data.PlannedDeliveryDurationInDays,
		StorageLocation:               data.StorageLocation,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToProcurement() *requests.Procurement {
	dataProduct := sdc.Product
	dataPlant := sdc.Product.Plant
	data := sdc.Product.Plant.Procurement
	return &requests.Procurement{
		Product:                     dataProduct.Product,
		Plant:                       dataPlant.Plant,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
	}
}

func (sdc *SDC) ConvertToWorkScheduling() *requests.WorkScheduling {
	dataProduct := sdc.Product
	dataPlant := sdc.Product.Plant
	data := sdc.Product.Plant.WorkScheduling
	return &requests.WorkScheduling{
		Product:                       dataProduct.Product,
		Plant:                         dataPlant.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
	}
}

func (sdc *SDC) ConvertToSalesPlant() *requests.SalesPlant {
	dataProduct := sdc.Product
	dataPlant := sdc.Product.Plant
	data := sdc.Product.Plant.SalesPlant
	return &requests.SalesPlant{
		Product:               dataProduct.Product,
		Plant:                 dataPlant.Plant,
		LoadingGroup:          data.LoadingGroup,
		AvailabilityCheckType: data.AvailabilityCheckType,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	dataProduct := sdc.Product
	data := sdc.Product.Accounting
	return &requests.Accounting{
		Product:             dataProduct.Product,
		ValuationArea:       data.ValuationArea,
		ValuationClass:      data.ValuationClass,
		StandardPrice:       data.StandardPrice,
		PriceUnitQty:        data.PriceUnitQty,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		PlannedPrice:        data.PlannedPrice,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToSalesOrganization() *requests.SalesOrganization {
	dataProduct := sdc.Product
	data := sdc.Product.SalesOrganization
	return &requests.SalesOrganization{
		Product:                        dataProduct.Product,
		ProductSalesOrg:                data.ProductSalesOrg,
		ProductDistributionChnl:        data.ProductDistributionChnl,
		SupplyingPlant:                 data.SupplyingPlant,
		PriceSpecificationProductGroup: data.PriceSpecificationProductGroup,
		AccountDetnProductGroup:        data.AccountDetnProductGroup,
		ItemCategoryGroup:              data.ItemCategoryGroup,
		SalesMeasureUnit:               data.SalesMeasureUnit,
		ProductHierarchy:               data.ProductHierarchy,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
		ToSalesTax: struct {
			ToSalesTaxResults []*requests.SalesTax `json:"results"`
		}{
			ToSalesTaxResults: []*requests.SalesTax{
				sdc.ConvertToSalesTax(),
			},
		},
	}
}

func (sdc *SDC) ConvertToSalesTax() *requests.SalesTax {
	dataProduct := sdc.Product
	data := sdc.Product.SalesTax
	return &requests.SalesTax{
		Product:           dataProduct.Product,
		Country:           data.Country,
		TaxCategory:       data.TaxCategory,
		TaxClassification: data.TaxClassification,
	}
}

func (sdc *SDC) ConvertToProductDesc() *requests.ProductDesc {
	dataProduct := sdc.Product
	data := sdc.Product.ProductDescription
	return &requests.ProductDesc{
		Product:            dataProduct.Product,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func (sdc *SDC) ConvertToQuality() *requests.Quality {
	dataProduct := sdc.Product
	dataPlant := sdc.Product.Plant
	data := sdc.Product.Plant.Quality
	return &requests.Quality{
		Product:                        dataProduct.Product,
		Plant:                          dataPlant.Plant,
		MaximumStoragePeriod:           data.MaximumStoragePeriod,
		QualityMgmtCtrlKey:             data.QualityMgmtCtrlKey,
		MatlQualityAuthorizationGroup:  data.MatlQualityAuthorizationGroup,
		HasPostToInspectionStock:       data.HasPostToInspectionStock,
		InspLotDocumentationIsRequired: data.InspLotDocumentationIsRequired,
		SuplrQualityManagementSystem:   data.SuplrQualityManagementSystem,
		RecrrgInspIntervalTimeInDays:   data.RecrrgInspIntervalTimeInDays,
		ProductQualityCertificateType:  data.ProductQualityCertificateType,
	}
}
