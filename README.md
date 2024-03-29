# sap-api-integrations-product-master-creates  
sap-api-integrations-product-master-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で品目マスタデータを登録するマイクロサービスです。  
sap-api-integrations-product-master-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-product-master-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PRODUCT_SRV_0001/overview  

## 動作環境  
sap-api-integrations-product-master-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-product-master-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-product-master-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PRODUCT_SRV_0001/overview  
* APIサービス名(=baseURL): API_PRODUCT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-product-master-creates には、次の API をコールするためのリソースが含まれています。  

* A_Product（品目マスタ - 一般データ）
* A_ProductPlant（品目マスタ - プラントデータ）
* A_ProductStorageLocation（品目マスタ - 保管場所データ）
* A_ProductPlantMRPArea（品目マスタ - MRPエリアデータ）
* A_ProductPlantProcurement（品目マスタ - 購買データ）
* A_ProductWorkScheduling（品目マスタ - 作業計画データ）
* A_ProductPlantSales（品目マスタ - 販売プラントデータ）
* A_ProductValuationAccount（品目マスタ - 評価エリアデータ）
* A_ProductSalesDelivery（品目マスタ - 販売組織データ）
* A_ProductPlantQualityMgmt（品目マスタ - 品質管理データ）
* A_ProductSalesTax（品目マスタ - 販売税データ）
* A_ProductDescription（品目マスタ - テキストデータ）  

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General", "Plant", "Accounting" が指定されています。    
  
```
"api_schema": "SAPProductMasterCreates",
"accepter": ["General", "Plant", "Accounting"],
"material_code": "",
"deleted": false
```
  
* 全データを登録する際のsample.jsonの記載例(2)  

全データを登録する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPProductMasterCreates",
"accepter": ["All"],
"material_code": "",
"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostProductMaster(
	general *requests.General,
	plant *requests.Plant,
	storageLocation *requests.StorageLocation,
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
		case "StorageLocation":
			func() {
				c.StorageLocation(storageLocation)
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 品目マスタ の 一般データ が登録された結果の JSON の例です。  
以下の項目のうち、"Product" ～ "ProductDescription" は、/SAP_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/home/ampamman/go/src/sap-api-integrations-product-master-creates/SAP_API_Caller/caller.go#L113",
	"function": "sap-api-integrations-product-master-creates/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": {
		"Product": "52",
		"IndustrySector": "M",
		"ProductType": "FERT",
		"BaseUnit": "PC",
		"ValidityStartDate": "",
		"ProductGroup": "01",
		"Division": "",
		"GrossWeight": "0.000",
		"WeightUnit": "",
		"SizeOrDimensionText": "",
		"ProductStandardID": "",
		"CreationDate": "",
		"LastChangeDate": "",
		"IsMarkedForDeletion": false,
		"NetWeight": "0.000",
		"ChangeNumber": "",
		"to_Description": {
			"Product": "",
			"Language": "EN",
			"ProductDescription": "test"
		}
	},
	"time": "2022-01-20T16:37:40.635494424+09:00"
}
```
