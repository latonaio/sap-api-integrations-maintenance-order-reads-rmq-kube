package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-order-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			MaintenanceOrder:               data.MaintenanceOrder,
			MaintOrderRoutingNumber:        data.MaintOrderRoutingNumber,
			MaintenanceOrderType:           data.MaintenanceOrderType,
			MaintenanceOrderDesc:           data.MaintenanceOrderDesc,
			MaintOrdBasicStartDateTime:     data.MaintOrdBasicStartDateTime,
			MaintOrdBasicEndDateTime:       data.MaintOrdBasicEndDateTime,
			MaintOrdBasicStartDate:         data.MaintOrdBasicStartDate,
			MaintOrdBasicStartTime:         data.MaintOrdBasicStartTime,
			MaintOrdBasicEndDate:           data.MaintOrdBasicEndDate,
			MaintOrdBasicEndTime:           data.MaintOrdBasicEndTime,
			MaintOrdSchedldBscStrtDateTime: data.MaintOrdSchedldBscStrtDateTime,
			MaintOrdSchedldBscEndDateTime:  data.MaintOrdSchedldBscEndDateTime,
			ScheduledBasicStartDate:        data.ScheduledBasicStartDate,
			ScheduledBasicStartTime:        data.ScheduledBasicStartTime,
			ScheduledBasicEndDate:          data.ScheduledBasicEndDate,
			ScheduledBasicEndTime:          data.ScheduledBasicEndTime,
			MaintenanceNotification:        data.MaintenanceNotification,
			OrdIsNotSchedldAutomatically:   data.OrdIsNotSchedldAutomatically,
			MainWorkCenterInternalID:       data.MainWorkCenterInternalID,
			MainWorkCenterTypeCode:         data.MainWorkCenterTypeCode,
			MainWorkCenter:                 data.MainWorkCenter,
			MainWorkCenterPlant:            data.MainWorkCenterPlant,
			WorkCenterInternalID:           data.WorkCenterInternalID,
			WorkCenterTypeCode:             data.WorkCenterTypeCode,
			WorkCenter:                     data.WorkCenter,
			MaintenancePlanningPlant:       data.MaintenancePlanningPlant,
			MaintenancePlant:               data.MaintenancePlant,
			Assembly:                       data.Assembly,
			MaintOrdProcessPhaseCode:       data.MaintOrdProcessPhaseCode,
			MaintOrdProcessSubPhaseCode:    data.MaintOrdProcessSubPhaseCode,
			BusinessArea:                   data.BusinessArea,
			CompanyCode:                    data.CompanyCode,
			CostCenter:                     data.CostCenter,
			ReferenceElement:               data.ReferenceElement,
			FunctionalArea:                 data.FunctionalArea,
			AdditionalDeviceData:           data.AdditionalDeviceData,
			Equipment:                      data.Equipment,
			EquipmentName:                  data.EquipmentName,
			FunctionalLocation:             data.FunctionalLocation,
			MaintenanceOrderPlanningCode:   data.MaintenanceOrderPlanningCode,
			MaintenancePlannerGroup:        data.MaintenancePlannerGroup,
			MaintenanceActivityType:        data.MaintenanceActivityType,
			MaintPriority:                  data.MaintPriority,
			MaintPriorityType:              data.MaintPriorityType,
			OrderProcessingGroup:           data.OrderProcessingGroup,
			ProfitCenter:                   data.ProfitCenter,
			ResponsibleCostCenter:          data.ResponsibleCostCenter,
			MaintenanceRevision:            data.MaintenanceRevision,
			SerialNumber:                   data.SerialNumber,
			SuperiorProjectNetwork:         data.SuperiorProjectNetwork,
			OperationSystemCondition:       data.OperationSystemCondition,
			WBSElement:                     data.WBSElement,
			WBSElementInternalID:           data.WBSElementInternalID,
			ControllingObjectClass:         data.ControllingObjectClass,
			MaintenanceOrderInternalID:     data.MaintenanceOrderInternalID,
			MaintenanceObjectList:          data.MaintenanceObjectList,
			MaintObjectLocAcctAssgmtNmbr:   data.MaintObjectLocAcctAssgmtNmbr,
			AssetLocation:                  data.AssetLocation,
			AssetRoom:                      data.AssetRoom,
			PlantSection:                   data.PlantSection,
			BasicSchedulingType:            data.BasicSchedulingType,
			LatestAcceptableCompletionDate: data.LatestAcceptableCompletionDate,
			MaintOrdPersonResponsible:      data.MaintOrdPersonResponsible,
			LastChangeDateTime:             data.LastChangeDateTime,
			ControllingSettlementProfile:   data.ControllingSettlementProfile,
			SystemStatusText:               data.SystemStatusText,
			UserStatusText:                 data.UserStatusText,
			TechnicalObject:                data.TechnicalObject,
			TechnicalObjectLabel:           data.TechnicalObjectLabel,
			TechObjIsEquipOrFuncnlLoc:      data.TechObjIsEquipOrFuncnlLoc,
			ToOperation:                    data.ToOperation.Deferred.URI,
			ToObjectListItem:               data.ToObjectListItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToObjectListItem(raw []byte, l *logger.Logger) ([]ToObjectListItem, error) {
	pm := &responses.ToObjectListItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToObjectListItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toObjectListItem := make([]ToObjectListItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toObjectListItem = append(toObjectListItem, ToObjectListItem{
			MaintenanceOrder:            data.MaintenanceOrder,
			MaintenanceOrderObjectList:  data.MaintenanceOrderObjectList,
			MaintenanceObjectListItem:   data.MaintenanceObjectListItem,
			Equipment:                   data.Equipment,
			MaintenanceNotification:     data.MaintenanceNotification,
			Assembly:                    data.Assembly,
			Material:                    data.Material,
			SerialNumber:                data.SerialNumber,
			UniqueItemIdentifier:        data.UniqueItemIdentifier,
			FunctionalLocation:          data.FunctionalLocation,
			MaintObjectListItemSequence: data.MaintObjectListItemSequence,
		})
	}

	return toObjectListItem, nil
}

func ConvertToToOperation(raw []byte, l *logger.Logger) ([]ToOperation, error) {
	pm := &responses.ToOperation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToOperation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toOperation := make([]ToOperation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toOperation = append(toOperation, ToOperation{
			MaintenanceOrder:               data.MaintenanceOrder,
			MaintenanceOrderOperation:      data.MaintenanceOrderOperation,
			MaintenanceOrderSubOperation:   data.MaintenanceOrderSubOperation,
			OperationControlKey:            data.OperationControlKey,
			OperationWorkCenterInternalID:  data.OperationWorkCenterInternalID,
			WorkCenter:                     data.WorkCenter,
			Plant:                          data.Plant,
			OperationStandardTextCode:      data.OperationStandardTextCode,
			OperationDescription:           data.OperationDescription,
			MaintOrderRoutingNumber:        data.MaintOrderRoutingNumber,
			MaintenanceOrderRoutingNode:    data.MaintenanceOrderRoutingNode,
			SuperiorOperationInternalID:    data.SuperiorOperationInternalID,
			OperationWorkCenterTypeCode:    data.OperationWorkCenterTypeCode,
			Language:                       data.Language,
			NumberOfTimeTickets:            data.NumberOfTimeTickets,
			OperationPurgInfoRecdSearchTxt: data.OperationPurgInfoRecdSearchTxt,
			OperationSupplier:              data.OperationSupplier,
			CostElement:                    data.CostElement,
			OperationPurchasingInfoRecord:  data.OperationPurchasingInfoRecord,
			PurchasingOrganization:         data.PurchasingOrganization,
			PurchasingGroup:                data.PurchasingGroup,
			MaterialGroup:                  data.MaterialGroup,
			OpPurchaseOutlineAgreement:     data.OpPurchaseOutlineAgreement,
			OpPurchaseOutlineAgreementItem: data.OpPurchaseOutlineAgreementItem,
			OperationRequisitionerName:     data.OperationRequisitionerName,
			OperationTrackingNumber:        data.OperationTrackingNumber,
			NumberOfCapacities:             data.NumberOfCapacities,
			OperationWorkPercent:           data.OperationWorkPercent,
			OperationCalculationControl:    data.OperationCalculationControl,
			ActivityType:                   data.ActivityType,
			OperationSystemCondition:       data.OperationSystemCondition,
			OperationGoodsRecipientName:    data.OperationGoodsRecipientName,
			OperationUnloadingPointName:    data.OperationUnloadingPointName,
			OperationPersonResponsible:     data.OperationPersonResponsible,
			DeliveryTimeInDays:             data.DeliveryTimeInDays,
			MaintOrderOperationDuration:    data.MaintOrderOperationDuration,
			MaintOrdOperationDurationUnit:  data.MaintOrdOperationDurationUnit,
			OpBscStartDateConstraintType:   data.OpBscStartDateConstraintType,
			OpBscEndDateConstraintType:     data.OpBscEndDateConstraintType,
			MaintOrdOperationWorkDuration:  data.MaintOrdOperationWorkDuration,
			MaintOrdOpWorkDurationUnit:     data.MaintOrdOpWorkDurationUnit,
			MaintOrdOpConstraintStrtDteTme: data.MaintOrdOpConstraintStrtDteTme,
			ConstraintDateForBscStartDate:  data.ConstraintDateForBscStartDate,
			ConstraintTimeForBscStartTime:  data.ConstraintTimeForBscStartTime,
			MaintOrdOpCstrtFinishDteTme:    data.MaintOrdOpCstrtFinishDteTme,
			ConstraintDateForBscFinishDate: data.ConstraintDateForBscFinishDate,
			ConstraintTimeForBscFinishTime: data.ConstraintTimeForBscFinishTime,
			MaintOrdOperationExecutionRate: data.MaintOrdOperationExecutionRate,
			Equipment:                      data.Equipment,
			FunctionalLocation:             data.FunctionalLocation,
			MaintenanceActivityType:        data.MaintenanceActivityType,
			BusinessArea:                   data.BusinessArea,
			ProfitCenter:                   data.ProfitCenter,
			FunctionalArea:                 data.FunctionalArea,
			MaintControllingObjectClass:    data.MaintControllingObjectClass,
			WrkCtrIntCapRqmtsDistr:         data.WrkCtrIntCapRqmtsDistr,
			MaintOrdOperationOverheadCode:  data.MaintOrdOperationOverheadCode,
			MaintOrderOperationQuantity:    data.MaintOrderOperationQuantity,
			MaintOrdOperationQuantityUnit:  data.MaintOrdOperationQuantityUnit,
			Assembly:                       data.Assembly,
			MaintOperationExecStageCode:    data.MaintOperationExecStageCode,
			WBSElement:                     data.WBSElement,
			IsMarkedForDeletion:            data.IsMarkedForDeletion,
			MaintOrderOperationInternalID:  data.MaintOrderOperationInternalID,
			MaintenanceObjectListItem:      data.MaintenanceObjectListItem,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			OpErlstSchedldExecStrtDte:      data.OpErlstSchedldExecStrtDte,
			OpErlstSchedldExecStrtTme:      data.OpErlstSchedldExecStrtTme,
			OpErlstSchedldExecEndDte:       data.OpErlstSchedldExecEndDte,
			OpErlstSchedldExecEndTme:       data.OpErlstSchedldExecEndTme,
			OpLtstSchedldExecStrtDte:       data.OpLtstSchedldExecStrtDte,
			OpLtstSchedldExecStrtTme:       data.OpLtstSchedldExecStrtTme,
			OpLtstSchedldExecEndDte:        data.OpLtstSchedldExecEndDte,
			OpLtstSchedldExecEndTme:        data.OpLtstSchedldExecEndTme,
			OpActualExecutionStartDate:     data.OpActualExecutionStartDate,
			OpActualExecutionStartTime:     data.OpActualExecutionStartTime,
			OpActualExecutionEndDate:       data.OpActualExecutionEndDate,
			OpActualExecutionEndTime:       data.OpActualExecutionEndTime,
			ForecastWorkQuantity:           data.ForecastWorkQuantity,
			ActualWorkQuantity:             data.ActualWorkQuantity,
			MaintOrdOpProcessPhaseCode:     data.MaintOrdOpProcessPhaseCode,
			MaintOrdOpProcessSubPhaseCode:  data.MaintOrdOpProcessSubPhaseCode,
			SystemStatusText:               data.SystemStatusText,
			UserStatusText:                 data.UserStatusText,
			ToOperationComponent:           data.ToOperationComponent.Deferred.URI,
		})
	}

	return toOperation, nil
}

func ConvertToToOperationComponent(raw []byte, l *logger.Logger) ([]ToOperationComponent, error) {
	pm := &responses.ToOperationComponent{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToOperationComponent. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toOperationComponent := make([]ToOperationComponent, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toOperationComponent = append(toOperationComponent, ToOperationComponent{
			MaintenanceOrder:               data.MaintenanceOrder,
			MaintenanceOrderOperation:      data.MaintenanceOrderOperation,
			MaintenanceOrderSubOperation:   data.MaintenanceOrderSubOperation,
			MaintenanceOrderComponent:      data.MaintenanceOrderComponent,
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			ReservationType:                data.ReservationType,
			MaintOrderRoutingNumber:        data.MaintOrderRoutingNumber,
			MaintOrderOperationCounter:     data.MaintOrderOperationCounter,
			Product:                        data.Product,
			MaintOrdOperationComponentText: data.MaintOrdOperationComponentText,
			MaintOrdOpCompRequiredQuantity: data.MaintOrdOpCompRequiredQuantity,
			BaseUnit:                       data.BaseUnit,
			QuantityInUnitOfEntry:          data.QuantityInUnitOfEntry,
			UnitOfEntry:                    data.UnitOfEntry,
			RequirementDate:                data.RequirementDate,
			RequirementTime:                data.RequirementTime,
			Supplier:                       data.Supplier,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			MaintOrdOpCompItemCategory:     data.MaintOrdOpCompItemCategory,
			GoodsMovementType:              data.GoodsMovementType,
			ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
			MaterialGroup:                  data.MaterialGroup,
			ProductTypeCode:                data.ProductTypeCode,
			ServicePerformer:               data.ServicePerformer,
			PerformancePeriodStartDateTime: data.PerformancePeriodStartDateTime,
			PerformancePeriodStartDate:     data.PerformancePeriodStartDate,
			PerformancePeriodEndDate:       data.PerformancePeriodEndDate,
			PerformancePeriodEndDateTime:   data.PerformancePeriodEndDateTime,
			PerformancePeriodStartTime:     data.PerformancePeriodStartTime,
			PerformancePeriodEndTime:       data.PerformancePeriodEndTime,
			LeanServiceDuration:            data.LeanServiceDuration,
			LeanServiceDurationUnit:        data.LeanServiceDurationUnit,
			DistributionFunction:           data.DistributionFunction,
			SrvcSchedgIsAlignedWthOpWrkCtr: data.SrvcSchedgIsAlignedWthOpWrkCtr,
			MaintOrderCompDebitCreditCode:  data.MaintOrderCompDebitCreditCode,
			GoodsMovementIsAllowed:         data.GoodsMovementIsAllowed,
			MaintenanceOrderComponentBatch: data.MaintenanceOrderComponentBatch,
			QuantityIsFixed:                data.QuantityIsFixed,
			MaintOrdOpComponentGLAccount:   data.MaintOrdOpComponentGLAccount,
			MaintOrdOpCompCostingRelevancy: data.MaintOrdOpCompCostingRelevancy,
			MaintCompAltvProdUsgeRateInPct: data.MaintCompAltvProdUsgeRateInPct,
			MaintOrderOpComponentSortText:  data.MaintOrderOpComponentSortText,
			MaintOrdOpCompIsBulkProduct:    data.MaintOrdOpCompIsBulkProduct,
			MaterialProvisionType:          data.MaterialProvisionType,
			MaintOrdOpCompAssgdWBSElmntInt: data.MaintOrdOpCompAssgdWBSElmntInt,
			MaintOrderOpComponentPrice:     data.MaintOrderOpComponentPrice,
			MaintOrdOpComponentCurrency:    data.MaintOrdOpComponentCurrency,
			MatlCompIsMarkedForBackflush:   data.MatlCompIsMarkedForBackflush,
			PurchasingGroup:                data.PurchasingGroup,
			DeliveryTimeInDays:             data.DeliveryTimeInDays,
			MaintOrdOpCompGdsRecipientName: data.MaintOrdOpCompGdsRecipientName,
			MaintOrdOpCompUnloadingPtTxt:   data.MaintOrdOpCompUnloadingPtTxt,
			GoodsReceiptDurationInWorkDays: data.GoodsReceiptDurationInWorkDays,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			OperationLeadTimeOffset:        data.OperationLeadTimeOffset,
			OpsLeadTimeOffsetUnit:          data.OpsLeadTimeOffsetUnit,
			MaintOrdOpCompRequisitioner:    data.MaintOrdOpCompRequisitioner,
			MaintOrdOpCompProcmtTrckgNmbr:  data.MaintOrdOpCompProcmtTrckgNmbr,
			ResponsiblePurchaseOrg:         data.ResponsiblePurchaseOrg,
			MaintOrdOpCompSpecialStockType: data.MaintOrdOpCompSpecialStockType,
			VariableSizeDimension1:         data.VariableSizeDimension1,
			VariableSizeDimensionUnit:      data.VariableSizeDimensionUnit,
			VariableSizeCompFormulaKey:     data.VariableSizeCompFormulaKey,
			VariableSizeDimension2:         data.VariableSizeDimension2,
			NumberOfVariableSizeItem:       data.NumberOfVariableSizeItem,
			VariableSizeDimension3:         data.VariableSizeDimension3,
			VariableSizeItemQuantity:       data.VariableSizeItemQuantity,
			VariableSizeComponentUnit:      data.VariableSizeComponentUnit,
			RqmtDateIsEnteredManually:      data.RqmtDateIsEnteredManually,
			SupplierProduct:                data.SupplierProduct,
			MaintOrdCompPurOutlineAgrmtItm: data.MaintOrdCompPurOutlineAgrmtItm,
			MaintOrderComponentInternalID:  data.MaintOrderComponentInternalID,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			OverallLimitAmount:             data.OverallLimitAmount,
			ExpectedOverallLimitAmount:     data.ExpectedOverallLimitAmount,
		})
	}

	return toOperationComponent, nil
}
