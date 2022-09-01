package requests

type SalesOrganization struct {
	Product                        string  `json:"Product"`
	ProductSalesOrg                string  `json:"ProductSalesOrg"`
	ProductDistributionChnl        string  `json:"ProductDistributionChnl"`
	SupplyingPlant                 *string `json:"SupplyingPlant,omitempty"`
	PriceSpecificationProductGroup *string `json:"PriceSpecificationProductGroup,omitempty"`
	AccountDetnProductGroup        *string `json:"AccountDetnProductGroup,omitempty"`
	ItemCategoryGroup              *string `json:"ItemCategoryGroup,omitempty"`
	SalesMeasureUnit               *string `json:"SalesMeasureUnit,omitempty"`
	ProductHierarchy               *string `json:"ProductHierarchy,omitempty"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion,omitempty"`
	ToSalesTax                     struct {
		ToSalesTaxResults []*SalesTax `json:"results"`
	} `json:"to_SalesTax"`
}
