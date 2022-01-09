package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenanceOrder               string      `json:"MaintenanceOrder"`
			MaintOrderRoutingNumber        string      `json:"MaintOrderRoutingNumber"`
			MaintenanceOrderType           string      `json:"MaintenanceOrderType"`
			MaintenanceOrderDesc           string      `json:"MaintenanceOrderDesc"`
			MaintOrdBasicStartDateTime     string      `json:"MaintOrdBasicStartDateTime"`
			MaintOrdBasicEndDateTime       string      `json:"MaintOrdBasicEndDateTime"`
			MaintOrdBasicStartDate         string      `json:"MaintOrdBasicStartDate"`
			MaintOrdBasicStartTime         string      `json:"MaintOrdBasicStartTime"`
			MaintOrdBasicEndDate           string      `json:"MaintOrdBasicEndDate"`
			MaintOrdBasicEndTime           string      `json:"MaintOrdBasicEndTime"`
			MaintOrdSchedldBscStrtDateTime string      `json:"MaintOrdSchedldBscStrtDateTime"`
			MaintOrdSchedldBscEndDateTime  string      `json:"MaintOrdSchedldBscEndDateTime"`
			ScheduledBasicStartDate        string      `json:"ScheduledBasicStartDate"`
			ScheduledBasicStartTime        string      `json:"ScheduledBasicStartTime"`
			ScheduledBasicEndDate          string      `json:"ScheduledBasicEndDate"`
			ScheduledBasicEndTime          string      `json:"ScheduledBasicEndTime"`
			MaintenanceNotification        string      `json:"MaintenanceNotification"`
			OrdIsNotSchedldAutomatically   string      `json:"OrdIsNotSchedldAutomatically"`
			MainWorkCenterInternalID       string      `json:"MainWorkCenterInternalID"`
			MainWorkCenterTypeCode         string      `json:"MainWorkCenterTypeCode"`
			MainWorkCenter                 string      `json:"MainWorkCenter"`
			MainWorkCenterPlant            string      `json:"MainWorkCenterPlant"`
			WorkCenterInternalID           string      `json:"WorkCenterInternalID"`
			WorkCenterTypeCode             string      `json:"WorkCenterTypeCode"`
			WorkCenter                     string      `json:"WorkCenter"`
			MaintenancePlanningPlant       string      `json:"MaintenancePlanningPlant"`
			MaintenancePlant               string      `json:"MaintenancePlant"`
			Assembly                       string      `json:"Assembly"`
			MaintOrdProcessPhaseCode       string      `json:"MaintOrdProcessPhaseCode"`
			MaintOrdProcessSubPhaseCode    string      `json:"MaintOrdProcessSubPhaseCode"`
			BusinessArea                   string      `json:"BusinessArea"`
			CompanyCode                    string      `json:"CompanyCode"`
			CostCenter                     string      `json:"CostCenter"`
			ReferenceElement               string      `json:"ReferenceElement"`
			FunctionalArea                 string      `json:"FunctionalArea"`
			AdditionalDeviceData           string      `json:"AdditionalDeviceData"`
			Equipment                      string      `json:"Equipment"`
			EquipmentName                  string      `json:"EquipmentName"`
			FunctionalLocation             string      `json:"FunctionalLocation"`
			MaintenanceOrderPlanningCode   string      `json:"MaintenanceOrderPlanningCode"`
			MaintenancePlannerGroup        string      `json:"MaintenancePlannerGroup"`
			MaintenanceActivityType        string      `json:"MaintenanceActivityType"`
			MaintPriority                  string      `json:"MaintPriority"`
			MaintPriorityType              string      `json:"MaintPriorityType"`
			OrderProcessingGroup           string      `json:"OrderProcessingGroup"`
			ProfitCenter                   string      `json:"ProfitCenter"`
			ResponsibleCostCenter          string      `json:"ResponsibleCostCenter"`
			MaintenanceRevision            string      `json:"MaintenanceRevision"`
			SerialNumber                   string      `json:"SerialNumber"`
			SuperiorProjectNetwork         string      `json:"SuperiorProjectNetwork"`
			OperationSystemCondition       string      `json:"OperationSystemCondition"`
			WBSElement                     string      `json:"WBSElement"`
			WBSElementInternalID           string      `json:"WBSElementInternalID"`
			ControllingObjectClass         string      `json:"ControllingObjectClass"`
			MaintenanceOrderInternalID     string      `json:"MaintenanceOrderInternalID"`
			MaintenanceObjectList          string      `json:"MaintenanceObjectList"`
			MaintObjectLocAcctAssgmtNmbr   string      `json:"MaintObjectLocAcctAssgmtNmbr"`
			AssetLocation                  string      `json:"AssetLocation"`
			AssetRoom                      string      `json:"AssetRoom"`
			PlantSection                   string      `json:"PlantSection"`
			BasicSchedulingType            string      `json:"BasicSchedulingType"`
			LatestAcceptableCompletionDate string      `json:"LatestAcceptableCompletionDate"`
			MaintOrdPersonResponsible      string      `json:"MaintOrdPersonResponsible"`
			LastChangeDateTime             string      `json:"LastChangeDateTime"`
			ControllingSettlementProfile   string      `json:"ControllingSettlementProfile"`
			SystemStatusText               string      `json:"SystemStatusText"`
			UserStatusText                 string      `json:"UserStatusText"`
			TechnicalObject                string      `json:"TechnicalObject"`
			TechnicalObjectLabel           string      `json:"TechnicalObjectLabel"`
			TechObjIsEquipOrFuncnlLoc      string      `json:"TechObjIsEquipOrFuncnlLoc"`
			ToOperation struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MaintenanceOrderOperation"`
			ToObjectListItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MaintOrderObjectListItem"`
		} `json:"results"`
	} `json:"d"`
}
