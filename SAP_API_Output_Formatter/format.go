package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-master-creates-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	general := &General{
		Product:             data.Product,
		IndustrySector:      data.IndustrySector,
		ProductType:         data.ProductType,
		BaseUnit:            data.BaseUnit,
		ValidityStartDate:   data.ValidityStartDate,
		ProductGroup:        data.ProductGroup,
		Division:            data.Division,
		GrossWeight:         data.GrossWeight,
		WeightUnit:          data.WeightUnit,
		SizeOrDimensionText: data.SizeOrDimensionText,
		ProductStandardID:   data.ProductStandardID,
		CreationDate:        data.CreationDate,
		LastChangeDate:      data.LastChangeDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
		NetWeight:           data.NetWeight,
		ChangeNumber:        data.ChangeNumber,
		ToProductDesc:       data.ToProductDesc.ToProductDescResults[0],
	}

	return general, nil
}

func ConvertToPlant(raw []byte, l *logger.Logger) (*Plant, error) {
	p := &responses.Plant{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Plant. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	plant := &Plant{
		Product:                       data.Product,
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

	return plant, nil
}

func ConvertToMRPArea(raw []byte, l *logger.Logger) (*MRPArea, error) {
	p := &responses.MRPArea{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MRPArea. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	mrpArea := &MRPArea{
		Product:                       data.Product,
		Plant:                         data.Plant,
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

	return mrpArea, nil
}

func ConvertToProcurement(raw []byte, l *logger.Logger) (*Procurement, error) {
	p := &responses.Procurement{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Procurement. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	procurement := &Procurement{
		Product:                     data.Product,
		Plant:                       data.Plant,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
	}

	return procurement, nil
}

func ConvertToWorkScheduling(raw []byte, l *logger.Logger) (*WorkScheduling, error) {
	p := &responses.WorkScheduling{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WorkScheduling. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	workScheduling := &WorkScheduling{
		Product:                       data.Product,
		Plant:                         data.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
	}

	return workScheduling, nil
}

func ConvertToSalesPlant(raw []byte, l *logger.Logger) (*SalesPlant, error) {
	p := &responses.SalesPlant{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPlant. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	salesPlant := &SalesPlant{
		Product:               data.Product,
		Plant:                 data.Plant,
		LoadingGroup:          data.LoadingGroup,
		AvailabilityCheckType: data.AvailabilityCheckType,
	}

	return salesPlant, nil
}

func ConvertToAccounting(raw []byte, l *logger.Logger) (*Accounting, error) {
	p := &responses.Accounting{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Accounting. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	accounting := &Accounting{
		Product:             data.Product,
		ValuationArea:       data.ValuationArea,
		ValuationClass:      data.ValuationClass,
		StandardPrice:       data.StandardPrice,
		PriceUnitQty:        data.PriceUnitQty,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		PlannedPrice:        data.PlannedPrice,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}

	return accounting, nil
}

func ConvertToSalesOrganization(raw []byte, l *logger.Logger) (*SalesOrganization, error) {
	p := &responses.SalesOrganization{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesOrganization. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	salesOrganization := &SalesOrganization{
		Product:                        data.Product,
		ProductSalesOrg:                data.ProductSalesOrg,
		ProductDistributionChnl:        data.ProductDistributionChnl,
		SupplyingPlant:                 data.SupplyingPlant,
		PriceSpecificationProductGroup: data.PriceSpecificationProductGroup,
		AccountDetnProductGroup:        data.AccountDetnProductGroup,
		ItemCategoryGroup:              data.ItemCategoryGroup,
		SalesMeasureUnit:               data.SalesMeasureUnit,
		ProductHierarchy:               data.ProductHierarchy,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}

	return salesOrganization, nil
}

func ConvertToProductDesc(raw []byte, l *logger.Logger) (*ProductDesc, error) {
	p := &responses.ProductDesc{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductDesc. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	productDesc := &ProductDesc{
		Product:            data.Product,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}

	return productDesc, nil
}

func ConvertToQuality(raw []byte, l *logger.Logger) (*Quality, error) {
	p := &responses.Quality{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Quality. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	quality := &Quality{
		Product:                        data.Product,
		Plant:                          data.Plant,
		MaximumStoragePeriod:           data.MaximumStoragePeriod,
		QualityMgmtCtrlKey:             data.QualityMgmtCtrlKey,
		MatlQualityAuthorizationGroup:  data.MatlQualityAuthorizationGroup,
		HasPostToInspectionStock:       data.HasPostToInspectionStock,
		InspLotDocumentationIsRequired: data.InspLotDocumentationIsRequired,
		SuplrQualityManagementSystem:   data.SuplrQualityManagementSystem,
		RecrrgInspIntervalTimeInDays:   data.RecrrgInspIntervalTimeInDays,
		ProductQualityCertificateType:  data.ProductQualityCertificateType,
	}

	return quality, nil
}
