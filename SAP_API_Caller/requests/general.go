package requests

type General struct {
	Product             string         `json:"Product"`
	ProductType         *string        `json:"ProductType"`
	IsMarkedForDeletion *bool          `json:"IsMarkedForDeletion"`
	GrossWeight         *string        `json:"GrossWeight"`
	WeightUnit          *string        `json:"WeightUnit"`
	ProductGroup        *string        `json:"ProductGroup"`
	BaseUnit            *string        `json:"BaseUnit"`
	Division            *string        `json:"Division"`
	ValidityStartDate   *string        `json:"ValidityStartDate"`
	SizeOrDimensionText *string        `json:"SizeOrDimensionText"`
	ProductStandardID   *string        `json:"ProductStandardID"`
	IndustrySector      *string        `json:"IndustrySector"`
	ToProductDesc       *ToProductDesc `json:"to_Description"`
}

type ToProductDescResults struct {
	Product            string  `json:"Product"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type ToProductDesc struct {
	ToProductDescResults []ToProductDescResults `json:"results"`
}
