package requests

type SalesOrganization struct {
	Product                        string  `json:"Product"`
	ProductSalesOrg                string  `json:"ProductSalesOrg"`
	ProductDistributionChnl        string  `json:"ProductDistributionChnl"`
	SupplyingPlant                 *string `json:"SupplyingPlant"`
	PriceSpecificationProductGroup *string `json:"PriceSpecificationProductGroup"`
	AccountDetnProductGroup        *string `json:"AccountDetnProductGroup"`
	ItemCategoryGroup              *string `json:"ItemCategoryGroup"`
	SalesMeasureUnit               *string `json:"SalesMeasureUnit"`
	ProductHierarchy               *string `json:"ProductHierarchy"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
	ToSalesTax                     struct {
		ToSalesTaxResults []*SalesTax `json:"results"`
	} `json:"to_SalesTax"`
}
