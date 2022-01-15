package responses

type General struct {
	D struct {
		Product             string `json:"Product"`
		BaseUnit            string `json:"BaseUnit"`
		ValidityStartDate   string `json:"ValidityStartDate"`
		ProductGroup        string `json:"ProductGroup"`
		Division            string `json:"Division"`
		GrossWeight         string `json:"GrossWeight"`
		WeightUnit          string `json:"WeightUnit"`
		SizeOrDimensionText string `json:"SizeOrDimensionText"`
		ProductStandardID   string `json:"ProductStandardID"`
		ToProductDesc       struct {
			ToProductDescResults []struct {
				Product            string `json:"Product"`
				Language           string `json:"Language"`
				ProductDescription string `json:"ProductDescription"`
			} `json:"results"`
		} `json:"to_Description"`
	} `json:"d"`
}
