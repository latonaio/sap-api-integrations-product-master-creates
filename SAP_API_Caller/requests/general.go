package requests

type General struct {
	Product             string  `json:"Product"`
	ProductType         *string `json:"ProductType"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
	GrossWeight         *string `json:"GrossWeight"`
	WeightUnit          *string `json:"WeightUnit"`
	ProductGroup        *string `json:"ProductGroup"`
	BaseUnit            *string `json:"BaseUnit"`
	Division            *string `json:"Division"`
	NetWeight           *string `json:"NetWeight"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	SizeOrDimensionText *string `json:"SizeOrDimensionText"`
	ProductStandardID   *string `json:"ProductStandardID"`
	IndustrySector      *string `json:"IndustrySector"`
	ToProductDesc       *struct {
		ToProductDescResults []*ProductDesc `json:"results"`
	} `json:"to_Description"`
}
