package requests

type General struct {
	Product             string  `json:"Product"`
	ProductType         *string `json:"ProductType,omitempty"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion,omitempty"`
	GrossWeight         *string `json:"GrossWeight,omitempty"`
	WeightUnit          *string `json:"WeightUnit,omitempty"`
	ProductGroup        *string `json:"ProductGroup,omitempty"`
	BaseUnit            *string `json:"BaseUnit,omitempty"`
	Division            *string `json:"Division,omitempty"`
	NetWeight           *string `json:"NetWeight,omitempty"`
	ValidityStartDate   *string `json:"ValidityStartDate,omitempty"`
	SizeOrDimensionText *string `json:"SizeOrDimensionText,omitempty"`
	ProductStandardID   *string `json:"ProductStandardID,omitempty"`
	IndustrySector      *string `json:"IndustrySector,omitempty"`
	ToProductDesc       *struct {
		ToProductDescResults []*ProductDesc `json:"results"`
	} `json:"to_Description"`
}
