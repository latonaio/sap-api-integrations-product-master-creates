package responses

type SalesOrganization struct {
	D struct {
		Product                        string `json:"Product"`
		ProductSalesOrg                string `json:"ProductSalesOrg"`
		ProductDistributionChnl        string `json:"ProductDistributionChnl"`
		SupplyingPlant                 string `json:"SupplyingPlant"`
		PriceSpecificationProductGroup string `json:"PriceSpecificationProductGroup"`
		AccountDetnProductGroup        string `json:"AccountDetnProductGroup"`
		ItemCategoryGroup              string `json:"ItemCategoryGroup"`
		SalesMeasureUnit               string `json:"SalesMeasureUnit"`
		ProductHierarchy               string `json:"ProductHierarchy"`
		IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
	} `json:"d"`
}
