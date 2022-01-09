# sap-api-integrations-maintenance-order-reads-rmq-kube
sap-api-integrations-maintenance-order-reads-rmq-kube は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 保全指図データを取得するマイクロサービスです。    
sap-api-integrations-maintenance-order-reads-rmq-kube には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-maintenance-order-reads-rmq-kube は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。     
https://api.sap.com/api/OP_API_MAINTENANCEORDER_0001/overview   

## 動作環境  
sap-api-integrations-maintenance-order-reads-rmq-kube は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）  
・ RabbitMQ on Kubernetes  
・ RabbitMQ Client         　

## クラウド環境での利用
sap-api-integrations-maintenance-order-reads-rmq-kube は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## RabbitMQ からの JSON Input

sap-api-integrations-maintenance-order-reads-rmq-kube は、Inputとして、RabbitMQ からのメッセージをJSON形式で受け取ります。 
Input の サンプルJSON は、Inputs フォルダ内にあります。  

## RabbitMQ からのメッセージ受信による イベントドリヴン の ランタイム実行

sap-api-integrations-maintenance-order-reads-rmq-kube は、RabbitMQ からのメッセージを受け取ると、イベントドリヴンでランタイムを実行します。  
AION の仕様では、Kubernetes 上 の 当該マイクロサービスPod は 立ち上がったまま待機状態で当該メッセージを受け取り、（コンテナ起動などの段取時間をカットして）即座にランタイムを実行します。　

## RabbitMQ への JSON Output

sap-api-integrations-maintenance-order-reads-rmq-kube は、Outputとして、RabbitMQ へのメッセージをJSON形式で出力します。  
Output の サンプルJSON は、Outputs フォルダ内にあります。  

## RabbitMQ の マスタサーバ環境

sap-api-integrations-maintenance-order-reads-rmq-kube が利用する RabbitMQ のマスタサーバ環境は、[rabbitmq-on-kubernetes](https://github.com/latonaio/rabbitmq-on-kubernetes) です。  
当該マスタサーバ環境は、同じエッジコンピューティングデバイスに配置されても、別の物理(仮想)サーバ内に配置されても、どちらでも構いません。

## RabbitMQ の Golang Runtime ライブラリ
sap-api-integrations-maintenance-order-reads-rmq-kube は、RabbitMQ の Golang Runtime ライブラリ として、[rabbitmq-golang-client](https://github.com/latonaio/rabbitmq-golang-client)を利用しています。

## デプロイ・稼働
sap-api-integrations-maintenance-order-reads-rmq-kube の デプロイ・稼働 を行うためには、aion-service-definitions の services.yml に、本レポジトリの services.yml を設定する必要があります。

kubectl apply - f 等で Deployment作成後、以下のコマンドで Pod が正しく生成されていることを確認してください。
```
$ kubectl get pods
```

## 本レポジトリ が 対応する API サービス
sap-api-integrations-maintenance-order-reads-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MAINTENANCEORDER_0001/overview    
* APIサービス名(=baseURL): API_MAINTENANCEORDER

## 本レポジトリ に 含まれる API名
sap-api-integrations-maintenance-order-reads-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* MaintenanceOrder（保全指図 - ヘッダ）※保全指図関連データを取得するために、ToObjectListItem、ToOperationData、ToOperationComponentData、と合わせて利用されます。
* ToObjectListItem（保全指図 - 対象一覧明細）
* ToOperation（保全指図 - オペレーション）
* ToOperationComponent（保全指図 - オペレーション構成品目）

## API への 値入力条件 の 初期値
sap-api-integrations-maintenance-order-reads-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.MaintenanceOrder.MaintenanceOrder（保全指図）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "sap.s4.beh.maintenanceorder.v1.MaintenanceOrder.Created.v1",
	"accepter": ["Header"],
	"maintenance_order": "4000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.maintenanceorder.v1.MaintenanceOrder.Created.v1",
	"accepter": ["All"],
	"maintenance_order": "4000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetMaintenanceOrder(maintenanceOrder string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(maintenanceOrder)
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
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 保全指図 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"MaintenanceOrder" ～ "to_MaintOrderObjectListItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-maintenance-order-reads-rmq-kube/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-maintenance-order-reads-rmq-kube/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"MaintenanceOrder": "4000000",
			"MaintOrderRoutingNumber": "1000000063",
			"MaintenanceOrderType": "YBA2",
			"MaintenanceOrderDesc": "Mechanical Inspection for Pump",
			"MaintOrdBasicStartDateTime": "/Date(1501027200000+0000)/",
			"MaintOrdBasicEndDateTime": "/Date(1501113600000+0000)/",
			"MaintOrdBasicStartDate": "/Date(1501027200000)/",
			"MaintOrdBasicStartTime": "PT00H00M00S",
			"MaintOrdBasicEndDate": "/Date(1501113600000)/",
			"MaintOrdBasicEndTime": "PT00H00M00S",
			"MaintOrdSchedldBscStrtDateTime": "/Date(1501048800000+0000)/",
			"MaintOrdSchedldBscEndDateTime": "/Date(1501052625000+0000)/",
			"ScheduledBasicStartDate": "/Date(1501027200000)/",
			"ScheduledBasicStartTime": "PT06H00M00S",
			"ScheduledBasicEndDate": "/Date(1501027200000)/",
			"ScheduledBasicEndTime": "PT07H03M45S",
			"MaintenanceNotification": "10000002",
			"OrdIsNotSchedldAutomatically": "",
			"MainWorkCenterInternalID": "10000063",
			"MainWorkCenterTypeCode": "A",
			"MainWorkCenter": "RES-0300",
			"MainWorkCenterPlant": "1710",
			"WorkCenterInternalID": "0",
			"WorkCenterTypeCode": "",
			"WorkCenter": "",
			"MaintenancePlanningPlant": "1710",
			"MaintenancePlant": "1710",
			"Assembly": "",
			"MaintOrdProcessPhaseCode": "",
			"MaintOrdProcessSubPhaseCode": "",
			"BusinessArea": "",
			"CompanyCode": "1710",
			"CostCenter": "17101301",
			"ReferenceElement": "",
			"FunctionalArea": "",
			"AdditionalDeviceData": "",
			"Equipment": "217100002",
			"EquipmentName": "Control Valve",
			"FunctionalLocation": "1710-SPA-SAC-PLAR1-INLT-SCTV",
			"MaintenanceOrderPlanningCode": "1",
			"MaintenancePlannerGroup": "930",
			"MaintenanceActivityType": "YB4",
			"MaintPriority": "",
			"MaintPriorityType": "PM",
			"OrderProcessingGroup": "0",
			"ProfitCenter": "YB110",
			"ResponsibleCostCenter": "17101301",
			"MaintenanceRevision": "",
			"SerialNumber": "",
			"SuperiorProjectNetwork": "",
			"OperationSystemCondition": "",
			"WBSElement": "",
			"WBSElementInternalID": "0",
			"ControllingObjectClass": "OC",
			"MaintenanceOrderInternalID": "OR000004000000",
			"MaintenanceObjectList": "203",
			"MaintObjectLocAcctAssgmtNmbr": "000000001913",
			"AssetLocation": "YB_1701",
			"AssetRoom": "",
			"PlantSection": "YOH",
			"BasicSchedulingType": "1",
			"LatestAcceptableCompletionDate": "",
			"MaintOrdPersonResponsible": "",
			"LastChangeDateTime": "",
			"ControllingSettlementProfile": "YEAM01",
			"SystemStatusText": "REL  CNF  GMPS JBFI MANC PRC  SETC TTJL",
			"UserStatusText": "",
			"TechnicalObject": "000000000217100002",
			"TechnicalObjectLabel": "217100002",
			"TechObjIsEquipOrFuncnlLoc": "EAMS_EQUI",
			"to_MaintenanceOrderOperation": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MAINTENANCEORDER/MaintenanceOrder('4000000')/to_MaintenanceOrderOperation",
			"to_MaintOrderObjectListItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MAINTENANCEORDER/MaintenanceOrder('4000000')/to_MaintOrderObjectListItem"
		}
	],
	"time": "2021-12-24T19:00:39.583823+09:00"
}
```
