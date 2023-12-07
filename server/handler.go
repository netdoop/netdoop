package server

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/heypkg/store/jsontype"
	"github.com/netdoop/cwmp/acs"
	"github.com/netdoop/cwmp/proto"
	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type AcsHanlder struct {
	logger *zap.Logger
}

func NewAcsHanlder() acs.AcsHanlder {
	s := AcsHanlder{
		logger: utils.GetLogger().Named("acs-handler"),
	}
	return &s
}

// HandleFault implements acs.AcsHanlder.
func (*AcsHanlder) HandleFault(acsDevice acs.Device, id string, v *proto.SoapFault) error {
	logger := utils.GetLogger()
	if v == nil {
		return nil
	}
	device, ok := acsDevice.(*omc.Device)
	if !ok {
		return nil
	}

	logger.Debug("SoapFault", zap.Any("SoapFault", v))
	faultCode := 0
	faultString := ""
	values := jsontype.Tags{}
	if v.FaultCode == "Client" && v.FaultString == "CWMP Fault" {
		for _, v2 := range v.Detail.Fault.SetParameterValuesFaults {
			logger.Debug("SetParameterValuesFault",
				zap.Int("FaultCode", v2.FaultCode),
				zap.String("FaultString", v2.FaultString),
				zap.String("ParameterName", v2.ParameterName),
			)
			values[v2.ParameterName] = fmt.Sprintf("%d:%v", v2.FaultCode, v2.FaultString)
		}
		faultCode = cast.ToInt(v.Detail.Fault.FaultCode)
		faultString = v.Detail.Fault.FaultString
	} else {
		faultString = v.FaultString
	}
	device.UpdateMethodCallResponse(id, values, faultCode, faultString)
	return nil
}

// GetDevice implements acs.AcsHanlder.
func (*AcsHanlder) GetDevice(schema string, oui string, productClass string, serialNumber string) acs.Device {
	return omc.GetDevice(schema, oui, serialNumber)
}

// GetProduct implements acs.AcsHanlder.
func (*AcsHanlder) GetProduct(schema string, oui string, productClass string) acs.Product {
	return omc.GetProduct(schema, oui, productClass)
}

// HandleInform implements acs.AcsHanlder.
func (s *AcsHanlder) HandleInform(ctx context.Context, inform *proto.Inform) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if inform == nil {
		return nil
	}
	s.logger.Debug("Inform", zap.Any("Inform", inform))
	db := store.GetDB()

	product := omc.GetProduct("", inform.DeviceID.OUI, inform.DeviceID.ProductClass)
	if product == nil {
		return errors.New("unknow product")
	}
	now := time.Now()

	values := map[string]string{}
	for _, param := range inform.ParameterList.ParameterValueStructs {
		values[param.Name.Text] = param.Value.Text
	}

	lastOnlineStatus := false
	currentTime := omc.ParseTime(inform.CurrentTime)
	periodicOnly := true
	needSync := false
	for _, event := range inform.Event.Events {
		if event.EventCode != "2 PERIODIC" && event.EventCode != "10 AUTONOMOUS TRANSFER COMPLETE" {
			periodicOnly = false
		}
		if event.EventCode == "1 BOOT" || event.EventCode == "0 BOOTSTRAP" {
			needSync = true
		}
	}

	device := omc.GetDeviceWithProductClass("", inform.DeviceID.OUI, inform.DeviceID.ProductClass, inform.DeviceID.SerialNumber)
	if device == nil {
		device = &omc.Device{
			Oui:          inform.DeviceID.OUI,
			ProductClass: inform.DeviceID.ProductClass,
			SerialNumber: inform.DeviceID.SerialNumber,
		}
		device.Name = fmt.Sprintf("%v-%v", inform.DeviceID.OUI, inform.DeviceID.SerialNumber)
		device.UpdateMetaData("Manufacturer", inform.DeviceID.Manufacturer)
		device.UpdateMetaData("OUI", inform.DeviceID.OUI)
		device.UpdateMetaData("ProductClass", inform.DeviceID.ProductClass)
		device.UpdateMetaData("SerialNumber", inform.DeviceID.SerialNumber)

		device.UpdateParameterValues(values)
		device.UpdateByIgdValues(db, values)
		device.Methods = &omc.Methods{"GetRPCMethods"}
		device.ProductID = product.ID
		device.ProductType = product.ProductType
		device.Online = true
		device.Enable = true
		device.LastInformTime = jsontype.JSONTime(now)

		result := db.Create(device)
		if result.Error != nil {
			return errors.Wrap(result.Error, "create new device")
		}

		device.PushMethodCall(time.Now(), "GetRPCMethods", nil)
	} else {
		lastOnlineStatus = device.Online
		device.UpdateParameterValues(values)
		device.UpdateByIgdValues(db, values)
		device.FetchPmValues(db, now)

		updateColumns := []string{"online", "meta_data", "properties", "parameter_values", "last_inform_time"}
		updateData := &omc.Device{
			Online:          true,
			MetaData:        device.MetaData,
			Properties:      device.Properties,
			ParameterValues: device.ParameterValues,
			LastInformTime:  jsontype.JSONTime(now),
		}
		updateData.SaveData()
		result := db.Model(device).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update device")
		}
	}
	device.Product = product
	omc.HandleDeviceAlive(device, currentTime, lastOnlineStatus)

	if needSync {
		device.PushGetDeviceParameterNames("Device.", false)
	}

	if !periodicOnly {
		metaData := &jsontype.Tags{}
		for _, v := range inform.Event.Events {
			metaData.Set(v.EventCode, v.CommandKey)
		}
		if err := omc.InsertDevieEvent(db, device, "Inform", omc.ParseTime(inform.CurrentTime), metaData); err != nil {
			s.logger.Error("insert device event", zap.Error(err))
		}
	}
	return nil
}

// HandleGetRPCMethodsResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleGetRPCMethodsResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.GetRPCMethodsResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("GetRPCMethodsResponse", zap.Any("MethodList", resp.MethodList.Strings))

	db := store.GetDB()
	device.UpdateMethodCallResponse(id, jsontype.StringArrayToTags(resp.MethodList.Strings), 0, "")

	methodListStrings := resp.MethodList.Strings
	if len(methodListStrings) == 0 {
		methodListStrings = []string{
			"SetParameterValues",
			"GetParameterValues",
			"GetParameterNames",
			"SetParameterAttributes",
			"GetParameterAttributes",
			"AddObject",
			"DeleteObject",
			"Reboot",
			"Download",
			"Upload",
			"FactoryReset",
		}
	}
	methods := omc.Methods(methodListStrings)
	updateColumns := []string{"methods"}
	updateData := &omc.Device{
		Methods: &methods,
	}
	updateData.SaveData()
	result := db.Model(device).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update device")
	}
	return nil
}

// HandleGetParameterNamesResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleGetParameterNamesResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.GetParameterNamesResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("GetParameterNamesResponse", zap.Int("Length", len(resp.ParameterList.ParameterInfoStructs)))
	db := store.GetDB()

	values := map[string]bool{}
	for _, v := range resp.ParameterList.ParameterInfoStructs {
		values[v.Name] = cast.ToBool(v.Writable)
	}
	device.UpdateMethodCallResponse(id, jsontype.BoolMapToTags(values), 0, "")

	device.UpdateParameterWritables(values)
	updateColumns := []string{"parameter_writables", "meta_data"}
	updateData := &omc.Device{
		ParameterWritables: device.ParameterWritables,
		MetaData:           device.MetaData,
	}
	if device.IsMethodSupported("GetParameterValues") {
		values := map[string]any{}
		values["Device."] = ""
		device.PushMethodCall(time.Now(), "GetParameterValues", values)
	}
	updateData.SaveData()
	result := db.Model(device).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update device")
	}

	if product := omc.GetProduct("", device.Oui, device.ProductClass); product != nil {
		if dm := omc.GetDataModelByProduct(product); dm != nil {
			needReload := false
			if !GetState("sync_datamodel_param_writable", fmt.Sprintf("%v", product.ID)) {
				for _, v := range resp.ParameterList.ParameterInfoStructs {
					name := v.Name
					var writable *bool
					*writable = cast.ToBool(v.Writable)
					omc.UpsertDataModelParameter(db, dm, name, nil, writable, nil, nil)
					needReload = true
				}
			}
			if needReload {
				SetState("sync_datamodel_param_writable", fmt.Sprintf("%v", product.ID), true)
				go omc.ReloadAllDataModels()
			}
		}
	}

	return nil
}

// HandleGetParameterValuesResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleGetParameterValuesResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.GetParameterValuesResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}

	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("GetParameterValuesResponse", zap.Int("ParameterList", len(resp.ParameterList.ParameterValueStructs)))
	db := store.GetDB()

	values := map[string]string{}
	for _, v := range resp.ParameterList.ParameterValueStructs {
		values[v.Name.Text] = v.Value.Text
	}
	device.UpdateMethodCallResponse(id, jsontype.StringMapToTags(values), 0, "")

	device.UpdateParameterValues(values)
	device.UpdateByIgdValues(db, values)

	updateColumns := []string{"parameter_values", "properties", "meta_data"}
	updateData := &omc.Device{
		MetaData:        device.MetaData,
		Properties:      device.Properties,
		ParameterValues: device.ParameterValues,
	}
	updateData.SaveData()
	result := db.Model(device).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update device")
	}

	if product := omc.GetProduct("", device.Oui, device.ProductClass); product != nil {
		if dm := omc.GetDataModelByProduct(product); dm != nil {
			needReload := false
			if !GetState("sync_datamodel_param_type", fmt.Sprintf("%v", product.ID)) {
				for _, v := range resp.ParameterList.ParameterValueStructs {
					name := v.Name.Text
					var typ *string
					for _, attr := range v.Value.Attrs {
						if attr.Name.Local == "type" {
							typ = &attr.Value
						}
					}
					omc.UpsertDataModelParameter(db, dm, name, typ, nil, nil, nil)
					needReload = true
				}
			}
			if needReload {
				SetState("sync_datamodel_param_type", fmt.Sprintf("%v", product.ID), true)
				go omc.ReloadAllDataModels()
			}
		}
	}

	return nil
}

// HandleSetParameterValuesResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleSetParameterValuesResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.SetParameterValuesResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("SetParameterValuesResponse", zap.Int("Status", resp.Status))
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleGetParameterAttributesResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleGetParameterAttributesResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.GetParameterAttributesResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)

	s.logger.Debug("GetParameterAttributesResponse", zap.Int("Length", len(resp.ParameterList.ParameterAttributesStructs)))
	values := map[string]int{}
	for _, v := range resp.ParameterList.ParameterAttributesStructs {
		values[v.Name] = v.Notification
	}

	device.UpdateMethodCallResponse(id, jsontype.IntMapToTags(values), 0, "")
	db := store.GetDB()
	device.UpdateParameterNotifications(values)
	updateColumns := []string{"parameter_notifications"}
	updateData := &omc.Device{
		ParameterNotifications: device.ParameterNotifications,
	}
	updateData.SaveData()
	result := db.Model(device).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update device")
	}
	return nil
}

// HandleSetParameterAttributesResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleSetParameterAttributesResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.SetParameterAttributesResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("SetParameterAttributesResponse")
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleAddObjectResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleAddObjectResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.AddObjectResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("AddObjectResponse",
		zap.Uint("InstanceNumber", resp.InstanceNumber),
		zap.Int("Status", resp.Status),
	)
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleDeleteObjectResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleDeleteObjectResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.DeleteObjectResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("DeleteObjectResponse",
		zap.Int("Status", resp.Status),
	)
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleUploadResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleUploadResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.UploadResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)

	s.logger.Debug("UploadResponse",
		zap.Int("Status", resp.Status),
		zap.String("StartTime", resp.StartTime),
		zap.String("CompleteTime", resp.CompleteTime),
	)
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleDownloadResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleDownloadResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.DownloadResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)

	s.logger.Debug("DownloadResponse",
		zap.Int("Status", resp.Status),
		zap.String("StartTime", resp.StartTime),
		zap.String("CompleteTime", resp.CompleteTime),
	)
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleFactoryResetResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleFactoryResetResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.FactoryResetResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("FactoryResetResponse")
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleRebootResponse implements acs.AcsHanlder.
func (s *AcsHanlder) HandleRebootResponse(ctx context.Context, acsDevice acs.Device, id string, resp *proto.RebootResponse) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if resp == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)
	s.logger.Debug("RebootResponse")
	device.UpdateMethodCallResponse(id, nil, 0, "")
	return nil
}

// HandleTransferComplete implements acs.AcsHanlder.
func (s *AcsHanlder) HandleTransferComplete(ctx context.Context, acsDevice acs.Device, req *proto.TransferComplete) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if req == nil {
		return nil
	}
	s.logger.Debug("TransferComplete", zap.Any("CommandKey", req.CommandKey))
	device, _ := acsDevice.(*omc.Device)
	db := store.GetDB()
	call := device.GetMethodCall(req.CommandKey)
	if call != nil {
		if ts := cast.ToInt64(req.CommandKey); ts != 0 {
			if req.FaultStruct.FaultCode == 0 {
				if err := omc.UpdateDeviceTransferLogComplete(db, device, ts,
					omc.ParseTime(req.StartTime),
					omc.ParseTime(req.CompleteTime),
				); err != nil {
					s.logger.Error("update device transfer log with complete", zap.Error(err))
				}
			} else {
				if err := omc.UpdateDeviceTransferLogFault(db, device, ts,
					omc.ParseTime(req.StartTime),
					omc.ParseTime(req.CompleteTime),
					req.FaultStruct.FaultCode,
					req.FaultStruct.FaultString,
				); err != nil {
					s.logger.Error("update device transfer log with fault", zap.Error(err))
				}
			}
		}
	}

	return nil
}

// HandleAutonomousTransferComplete implements acs.AcsHanlder.
func (s *AcsHanlder) HandleAutonomousTransferComplete(ctx context.Context, acsDevice acs.Device, req *proto.AutonomousTransferComplete) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if req == nil {
		return nil
	}
	device, _ := acsDevice.(*omc.Device)

	s.logger.Debug("AutonomousTransferComplete")
	db := store.GetDB()
	filename := path.Base(req.TransferURL)
	if filename == "." || filename == "/" {
		s.logger.Error("update device transfer log with fault, invalid filename",
			zap.String("filename", filename),
			zap.String("url", req.TransferURL))
		return nil
	}

	fileType, startTime, _, _, _, _, err := omc.ParseFileName(filename)
	if err != nil {
		s.logger.Error("update device transfer log with fault, parse filename", zap.Error(err), zap.String("filename", filename))
		return nil
	}

	key := filename
	ts := startTime.UnixNano()
	if err := omc.InsertDeviceTransferLogComplete(db, device, ts,
		omc.DeviceUploadBucket,
		key,
		fileType,
		filename,
		omc.ParseTime(req.StartTime),
		omc.ParseTime(req.CompleteTime),
		req.FaultStruct.FaultCode,
		req.FaultStruct.FaultString,
	); err != nil {
		s.logger.Error("update device transfer log with fault", zap.Error(err))
	}
	return nil
}

// HandleMesureValues implements acs.AcsHanlder.
func (s *AcsHanlder) HandleMesureValues(acsDevice acs.Device, filename string, values map[string]any) {
	device, _ := acsDevice.(*omc.Device)
	db := store.GetDB()
	for typ, value := range values {
		if meas := omc.GetKPIMeasure(device.Schema, "enb", typ); meas != nil && meas.Enable {
			if v := cast.ToFloat64(value); v != 0 {
				omc.InsertDeviePerformanceValue(db, device, meas.MeasTypeID, filename, v, nil)
			}
		}
	}

	measList := omc.GetKPIMeasuresBySet(device.Schema, "enb", "Customize")
	for _, meas := range measList {
		if meas.Enable && meas.FormulaExpression != nil {
			if result, err := meas.FormulaExpression.Evaluate(values); err != nil {
				s.logger.Error("evaluate", zap.String("formula", meas.Formula))
			} else {
				v := fmt.Sprintf("%v", result)
				if v != "NaN" {
					omc.InsertDeviePerformanceValue(db, device, meas.MeasTypeID, filename, cast.ToFloat64(v), nil)
				}
			}
		}
	}
}
