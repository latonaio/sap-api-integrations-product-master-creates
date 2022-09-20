package main

import (
	sap_api_caller "sap-api-integrations-product-master-creates/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-product-master-creates/SAP_API_Input_Reader"
	"sap-api-integrations-product-master-creates/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	fr := sap_api_input_reader.NewFileReader()
	pc := sap_api_request_client_header_setup.NewSAPRequestClientWithOption(conf.SAP)
	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		"100",
		pc,
		l,
	)
	inputSDC := fr.ReadSDC("./Inputs/SDC_Product_Master_General_sample.json")
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&inputSDC)

	accepter := getAccepter(inputSDC)
	general := inputSDC.ConvertToGeneral()
	plant := inputSDC.ConvertToPlant()
	mrpArea := inputSDC.ConvertToMRPArea()
	procurement := inputSDC.ConvertToProcurement()
	workScheduling := inputSDC.ConvertToWorkScheduling()
	salesPlant := inputSDC.ConvertToSalesPlant()
	accounting := inputSDC.ConvertToAccounting()
	salesOrganization := inputSDC.ConvertToSalesOrganization()
	productDesc := inputSDC.ConvertToProductDesc()
	quality := inputSDC.ConvertToQuality()

	caller.AsyncPostProductMaster(
		general,
		plant,
		mrpArea,
		procurement,
		workScheduling,
		salesPlant,
		accounting,
		salesOrganization,
		productDesc,
		quality,
		accepter,
	)
}

func getAccepter(sdc sap_api_input_reader.SDC) []string {
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"General", "Plant", "MRPArea", "Procurement",
			"WorkScheduling", "SalesPlant",
			"Accounting", "SalesOrganization", "ProductDesc",
			"Quality",
		}
	}
	return accepter
}
