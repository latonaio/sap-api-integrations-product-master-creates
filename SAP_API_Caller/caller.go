package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-product-master-creates-rmq-kube/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-product-master-creates-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	outputQueues    []string
	outputter       RMQOutputter
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		outputQueues:    outputQueueTo,
		outputter:       outputter,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostProductMaster(
	general *requests.General,
	plant *requests.Plant,
	mrpArea *requests.MRPArea,
	procurement *requests.Procurement,
	workScheduling *requests.WorkScheduling,
	salesPlant *requests.SalesPlant,
	accounting *requests.Accounting,
	salesOrganization *requests.SalesOrganization,
	productDesc *requests.ProductDesc,
	quality *requests.Quality,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Plant":
			func() {
				c.Plant(plant)
				wg.Done()
			}()
		case "MRPArea":
			func() {
				c.MRPArea(mrpArea)
				wg.Done()
			}()
		case "Procurement":
			func() {
				c.Procurement(procurement)
				wg.Done()
			}()
		case "WorkScheduling":
			func() {
				c.WorkScheduling(workScheduling)
				wg.Done()
			}()
		case "SalesPlant":
			func() {
				c.SalesPlant(salesPlant)
				wg.Done()
			}()
		case "Accounting":
			func() {
				c.Accounting(accounting)
				wg.Done()
			}()
		case "SalesOrganization":
			func() {
				c.SalesOrganization(salesOrganization)
				wg.Done()
			}()
		case "ProductDesc":
			func() {
				c.ProductDesc(productDesc)
				wg.Done()
			}()
		case "Quality":
			func() {
				c.Quality(quality)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(general *requests.General) {
	generalData, err := c.callProductSrvAPIRequirementGeneral("A_Product", general)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": generalData, "function": "ProductMasterGeneral"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementGeneral(api string, general *requests.General) (*sap_api_output_formatter.General, error) {
	body, err := json.Marshal(general)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Plant(plant *requests.Plant) {
	outputDataPlant, err := c.callProductSrvAPIRequirementPlant("A_ProductPlant", plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataPlant, "function": "ProductMasterPlant"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataPlant)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementPlant(api string, plant *requests.Plant) (*sap_api_output_formatter.Plant, error) {
	body, err := json.Marshal(plant)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MRPArea(mrpArea *requests.MRPArea) {
	outputDataMRPArea, err := c.callProductSrvAPIRequirementMRPArea("A_ProductPlantMRPArea", mrpArea)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataMRPArea, "function": "ProductMasterMRPArea"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataMRPArea)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementMRPArea(api string, mrpArea *requests.MRPArea) (*sap_api_output_formatter.MRPArea, error) {
	body, err := json.Marshal(mrpArea)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToMRPArea(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Procurement(procurement *requests.Procurement) {
	outputDataProcurement, err := c.callProductSrvAPIRequirementProcurement("A_ProductPlantProcurement", procurement)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataProcurement, "function": "ProductMasterProcurement"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataProcurement)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementProcurement(api string, procurement *requests.Procurement) (*sap_api_output_formatter.Procurement, error) {
	body, err := json.Marshal(procurement)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToProcurement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) WorkScheduling(workScheduling *requests.WorkScheduling) {
	outputDataWorkScheduling, err := c.callProductSrvAPIRequirementWorkScheduling("A_ProductWorkScheduling", workScheduling)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataWorkScheduling, "function": "ProductMasterWorkScheduling"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataWorkScheduling)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementWorkScheduling(api string, workScheduling *requests.WorkScheduling) (*sap_api_output_formatter.WorkScheduling, error) {
	body, err := json.Marshal(workScheduling)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToWorkScheduling(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesPlant(salesPlant *requests.SalesPlant) {
	outputDataSalesPlant, err := c.callProductSrvAPIRequirementSalesPlant("A_ProductPlantSales", salesPlant)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataSalesPlant, "function": "ProductMasterSalesPlant"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataSalesPlant)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementSalesPlant(api string, salesPlant *requests.SalesPlant) (*sap_api_output_formatter.SalesPlant, error) {
	body, err := json.Marshal(salesPlant)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToSalesPlant(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Accounting(accounting *requests.Accounting) {
	outputDataAccounting, err := c.callProductSrvAPIRequirementAccounting("A_ProductValuation", accounting)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataAccounting, "function": "ProductMasterAccounting"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataAccounting)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementAccounting(api string, accounting *requests.Accounting) (*sap_api_output_formatter.Accounting, error) {
	body, err := json.Marshal(accounting)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToAccounting(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) SalesOrganization(salesOrganization *requests.SalesOrganization) {
	outputDataSalesOrganization, err := c.callProductSrvAPIRequirementSalesOrganization("A_ProductSalesDelivery", salesOrganization)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataSalesOrganization, "function": "ProductMasterSalesOrganization"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataSalesOrganization)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementSalesOrganization(api string, salesOrganization *requests.SalesOrganization) (*sap_api_output_formatter.SalesOrganization, error) {
	body, err := json.Marshal(salesOrganization)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToSalesOrganization(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductDesc(productDesc *requests.ProductDesc) {
	outputDataProductDesc, err := c.callProductSrvAPIRequirementProductDesc("A_ProductDescription", productDesc)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataProductDesc, "function": "ProductMasterProductDescription"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataProductDesc)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementProductDesc(api string, productDesc *requests.ProductDesc) (*sap_api_output_formatter.ProductDesc, error) {
	body, err := json.Marshal(productDesc)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToProductDesc(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Quality(quality *requests.Quality) {
	outputDataQuality, err := c.callProductSrvAPIRequirementQuality("A_ProductPlantQualityMgmt", quality)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": outputDataQuality, "function": "ProductMasterQuality"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataQuality)
}

func (c *SAPAPICaller) callProductSrvAPIRequirementQuality(api string, quality *requests.Quality) (*sap_api_output_formatter.Quality, error) {
	body, err := json.Marshal(quality)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.requestClient.Request("POST", url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToQuality(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
