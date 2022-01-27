package responses

type General struct {
	D struct {
		Product             string `json:"Product"`
		IndustrySector      string `json:"IndustrySector"`
		ProductType         string `json:"ProductType"`
		BaseUnit            string `json:"BaseUnit"`
		ValidityStartDate   string `json:"ValidityStartDate"`
		ProductGroup        string `json:"ProductGroup"`
		Division            string `json:"Division"`
		GrossWeight         string `json:"GrossWeight"`
		WeightUnit          string `json:"WeightUnit"`
		SizeOrDimensionText string `json:"SizeOrDimensionText"`
		ProductStandardID   string `json:"ProductStandardID"`
		CreationDate        string `json:"CreationDate"`
		LastChangeDate      string `json:"LastChangeDate"`
		IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
		NetWeight           string `json:"NetWeight"`
		ChangeNumber        string `json:"ChangeNumber"`
		ToProductDesc       struct {
			ToProductDescResults []ToProductDesc `json:"results"`
		} `json:"to_Description"`
	} `json:"d"`
}
