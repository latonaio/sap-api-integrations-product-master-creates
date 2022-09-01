package requests

type Quality struct {
	Product                        string  `json:"Product"`
	Plant                          string  `json:"Plant"`
	MaximumStoragePeriod           *string `json:"MaximumStoragePeriod,omitempty"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey,omitempty"`
	MatlQualityAuthorizationGroup  *string `json:"MatlQualityAuthorizationGroup,omitempty"`
	HasPostToInspectionStock       *bool   `json:"HasPostToInspectionStock,omitempty"`
	InspLotDocumentationIsRequired *bool   `json:"InspLotDocumentationIsRequired,omitempty"`
	SuplrQualityManagementSystem   *string `json:"SuplrQualityManagementSystem,omitempty"`
	RecrrgInspIntervalTimeInDays   *string `json:"RecrrgInspIntervalTimeInDays,omitempty"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType,omitempty"`
}
