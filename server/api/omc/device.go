package omc

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/jsontype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type listDevicesData struct {
	Data  []omc.Device `json:"Data"`
	Total int64        `json:"Total"`
}

// HandleListDevices godoc
// @Summary List devices
// @Tags OMC Devices
// @ID list-devices
// @Accept json
// @Produce json
// @Param q query string false "Query" default()
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Security Bearer
// @Success 200 {object} listDevicesData
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices [get]
func HandleListDevices(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.Device](store.GetDB(), c, omc.SimpleDeviceInfoName, omc.DeviceSearchHanleFuncs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDevicesData{Data: data, Total: total})
}

type createDeviceBody struct {
	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerailNumber"`
	Name         string `json:"Name"`
}

// HandleCreateDevice godoc
// @Summary Create a device
// @Tags OMC Devices
// @ID create-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body createDeviceBody true "Device"
// @Success 200 {object} omc.Device
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 404 {object} echo.HTTPError "Not Found: invalid product"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices [post]
func HandleCreateDevice(c echo.Context) error {
	var data createDeviceBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	schema := cast.ToString(c.Get("schema"))
	db := store.GetDB()
	product := omc.GetProduct(schema, data.Oui, data.ProductClass)
	if product == nil {
		return echo.NewHTTPError(http.StatusNotFound, "invalid product")
	}

	device := omc.Device{
		Schema:       cast.ToString(c.Get("schema")),
		Oui:          data.Oui,
		ProductClass: data.ProductClass,
		SerialNumber: data.SerialNumber,
		Name:         data.Name,
		ProductID:    product.ID,
		ProductType:  product.ProductType,
	}

	result := db.Create(&device)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, device)
}

// HandleGetDevice godoc
// @Summary Get a device
// @Tags OMC Devices
// @ID get-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} omc.Device
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id} [get]
func HandleGetDevice(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.Device](c)
	return c.JSON(http.StatusOK, obj)
}

type updateDeviceInfoBody struct {
	Name string `json:"Name"`
}

// HandleUpdateDeviceInfo godoc
// @Summary Update device info
// @Tags OMC Devices
// @ID update-device-info
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body updateDeviceInfoBody true "Device"
// @Success 200 {object} omc.Device
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id} [put]
func HandleUpdateDeviceInfo(c echo.Context) error {
	var data updateDeviceInfoBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}

	updateColumns := []string{"name"}
	updateData := &omc.Device{
		Name: data.Name,
	}

	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	db := store.GetDB()
	if result := db.Model(device).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, device)
}

type setGroupForDeviceBody struct {
	GroupID uint `json:"GroupId"`
}

// HandleSetGroupForDevice godoc
// @Summary Set group for device
// @Tags OMC Devices
// @ID set-group-for-device
// @Security Bearer
// @Param id path int true "Device ID"
// @Param group_id body setGroupForDeviceBody true "Group ID"
// @Success 200 {object} omc.Device
// @Failure 400 {object} echo.HTTPError "Bad Request"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 404 {object} echo.HTTPError "Not Found"
// @Failure 500 {object} echo.HTTPError "Internal Server Error"
// @Router /omc/devices/{id}/group [put]
func HandleSetGroupForDevice(c echo.Context) error {
	var data setGroupForDeviceBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	db := store.GetDB()

	group := new(omc.Group)
	schema := cast.ToString(c.Get("schema"))
	if err := db.Where("schema = ? AND id = ?", schema, data.GroupID).First(group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "invalid group")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	updateColumns := []string{"group_id"}
	updateData := &omc.Device{
		GroupID: data.GroupID,
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	if result := db.Model(device).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, device)
}

// HandlesSetDeviceDisable godoc
// @Summary Set a device disable
// @Tags OMC Devices
// @ID set-device-disable
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} echo.Map
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/disable [put]
func HandlesSetDeviceDisable(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	db := store.GetDB()
	if result := db.Model(device).Update("enable", false); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, nil)
}

// HandlesSetDeviceEnable godoc
// @Summary Set a device enable
// @Tags OMC Devices
// @ID set-device-enable
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} echo.Map
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/enable [put]
func HandlesSetDeviceEnable(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	db := store.GetDB()
	if result := db.Model(device).Update("enable", true); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, nil)
}

// HandleDeleteDevice godoc
// @Summary Delete a device
// @Tags OMC Devices
// @ID delete-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id} [delete]
func HandleDeleteDevice(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	db := store.GetDB()
	if result := db.Delete(device); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}

// HandleRebootDevice godoc
// @Summary Reboot a device
// @Tags OMC Devices
// @ID reboot-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/reboot [post]
func HandleRebootDevice(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	if !device.IsMethodSupported("Reboot") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	call, err := device.PushMethodCall(time.Now(), "Reboot", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type getDeviceParameterNamesBody struct {
	ParameterPath string `json:"ParameterPath"`
	NextLevel     bool   `json:"NextLevel"`
}

// HandleGetDeviceParameterNames godoc
// @Summary Get parameter names of a device
// @Tags OMC Devices
// @ID get-device-parameter-names
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body getDeviceParameterNamesBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/get-parameter-names [post]
func HandleGetDeviceParameterNames(c echo.Context) error {
	var data getDeviceParameterNamesBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	call, err := device.PushGetDeviceParameterNames(data.ParameterPath, data.NextLevel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type getDeviceParameterValuesBody struct {
	Names []string `json:"Names"`
}

// HandleGetDeviceParameterValues godoc
// @Summary Get parameter values of a device
// @Tags OMC Devices
// @ID get-device-parameter-values
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body getDeviceParameterValuesBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/get-parameter-values [post]
func HandleGetDeviceParameterValues(c echo.Context) error {
	var data getDeviceParameterValuesBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	if !device.IsMethodSupported("GetParameterValues") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	values := map[string]any{}
	for _, n := range data.Names {
		values[n] = ""
	}
	call, err := device.PushMethodCall(time.Now(), "GetParameterValues", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type setDeviceParameterValuesBody struct {
	Values map[string]any `json:"Values"`
}

// HandleSetDeviceParameterValues godoc
// @Summary Set parameter values of a device
// @Tags OMC Devices
// @ID set-device-parameter-values
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body setDeviceParameterValuesBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/set-parameter-values [post]
func HandleSetDeviceParameterValues(c echo.Context) error {
	var data setDeviceParameterValuesBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	if !device.IsMethodSupported("SetParameterValues") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	call, err := device.PushMethodCall(time.Now(), "SetParameterValues", data.Values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	values := map[string]any{}
	for n, _ := range data.Values {
		values[n] = ""
	}
	device.PushMethodCall(time.Now(), "GetParameterValues", values)
	device.SendConnectRequest()

	return c.JSON(http.StatusOK, call)
}

type addDeviceObejectBody struct {
	ObjectName   string `json:"ObjectName"`
	ParameterKey string `json:"ParameterKey"`
}

// HandleAddDeviceObjectNames godoc
// @Summary Add object of a device
// @Tags OMC Devices
// @ID add-device-object
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body addDeviceObejectBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/add-device-object [post]
func HandleAddDeviceObject(c echo.Context) error {
	var data addDeviceObejectBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	if !device.IsMethodSupported("AddObject") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	values := map[string]any{
		"ObjectName":   data.ObjectName,
		"ParameterKey": data.ParameterKey,
	}
	call, err := device.PushMethodCall(time.Now(), "AddObject", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type deleteDeviceObejectBody struct {
	ObjectName   string `json:"ObjectName"`
	ParameterKey string `json:"ParameterKey"`
}

// HandleDeleteDeviceObject godoc
// @Summary Delete object of a device
// @Tags OMC Devices
// @ID delete-device-object
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body deleteDeviceObejectBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/delete-device-object [post]
func HandleDeleteDeviceObject(c echo.Context) error {
	var data deleteDeviceObejectBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	if !device.IsMethodSupported("DeleteObject") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	values := map[string]any{
		"ObjectName":   data.ObjectName,
		"ParameterKey": data.ParameterKey,
	}
	call, err := device.PushMethodCall(time.Now(), "DeleteObject", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type uploadDeviceFileBody struct {
	// Url          string `json:"Url"`
	// Username     string `json:"Username"`
	// Password     string `json:"Password"`
	FileType     string `json:"FileType"`
	DelaySeconds uint   `json:"DelaySeconds"`
}

// HandleUploadDeviceFile godoc
// @Summary Upload file of device
// @Tags OMC Devices
// @ID upload-device-file
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body uploadDeviceFileBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/upload-file [post]
func HandleUploadDeviceFile(c echo.Context) error {
	var data uploadDeviceFileBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	db := store.GetDB()
	if !device.IsMethodSupported("Upload") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	fileType := ""
	switch data.FileType[0:2] {
	case "1 ":
		fileType = "ConfigurationFile"
	case "2 ":
		fileType = "LogFile"
	case "X ":
		fileType = strings.ReplaceAll(data.FileType[2:], " ", "")
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "file type not supported")
	}

	t := time.Unix(time.Now().Unix(), 0)
	env := utils.GetEnv()
	schema := cast.ToString(c.Get("schema"))
	key := fmt.Sprintf("%v.%v.%v_%v_%v", device.Oui, device.ProductClass, device.SerialNumber, fileType, t.Format("20060102030405"))
	url := fmt.Sprintf("%v/%v", env.GetString("omc_upload_url"), key)
	log := omc.DeviceTransferLog{
		Time:         jsontype.JSONTime(t),
		Oui:          device.Oui,
		ProductClass: device.ProductClass,
		SerialNumber: device.SerialNumber,
		ProductType:  device.ProductType,
		DeviceID:     device.ID,
		Schema:       schema,
		TransferType: omc.TransferTypeUpload,
		FileType:     data.FileType,
		FileName:     key,
		ObjectBucket: omc.DeviceUploadBucket,
		ObjectKey:    key,
	}
	result := db.Create(&log)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	values := map[string]any{
		// "Url":          data.Url,
		// "Username":     data.Username,
		// "Password":     data.Password,
		"Url":          url,
		"FileType":     data.FileType,
		"DelaySeconds": data.DelaySeconds,
	}
	call, err := device.PushMethodCall(t, "Upload", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

type upgradeDeviceBody struct {
	FirmwareID   uint `json:"FirmwareID"`
	DelaySeconds uint `json:"DelaySeconds"`
}

// HandleUpgradeDevice godoc
// @Summary Upgrade device
// @Tags OMC Devices
// @ID ugrade-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body upgradeDeviceBody true "Data"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 400 {object} echo.HTTPError "BadRequest"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/upgrade [post]
func HandleUpgradeDevice(c echo.Context) error {
	var data upgradeDeviceBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	db := store.GetDB()
	if !device.IsMethodSupported("Download") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	firmware := omc.GetFirmware(db, device.Schema, data.FirmwareID)
	if firmware == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "firmware not found")
	}
	supported := false
	for _, product := range firmware.Products {
		if product.ID == device.ProductID {
			supported = true
			break
		}
	}
	if !supported {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid frmware for device")
	}

	t := time.Unix(time.Now().Unix(), 0)
	env := utils.GetEnv()
	schema := cast.ToString(c.Get("schema"))
	url := fmt.Sprintf("%v%v", env.GetString("omc_base_url"), firmware.S3Object.DownloadUrl)
	log := omc.DeviceTransferLog{
		Time:         jsontype.JSONTime(t),
		Oui:          device.Oui,
		ProductClass: device.ProductClass,
		SerialNumber: device.SerialNumber,
		ProductType:  device.ProductType,
		DeviceID:     device.ID,
		Schema:       schema,
		TransferType: omc.TransferTypeUpload,
		FileType:     omc.DownloadFileTypeFirmwareUpgradeImage,
		FileName:     firmware.Name,
		ObjectBucket: firmware.S3Object.Bucket,
		ObjectKey:    firmware.S3Object.Key,
		S3ObjectID:   &firmware.S3ObjectID,
		FirmwareID:   &firmware.ID,
	}
	result := db.Create(&log)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	values := map[string]any{
		"Url":          url,
		"FileType":     "1 Firmware Upgrade Image",
		"FileSize":     firmware.S3Object.FileSize,
		"FileName":     firmware.S3Object.FileName,
		"DelaySeconds": data.DelaySeconds,
	}
	call, err := device.PushMethodCall(t, "Download", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

// HandleSetDevicePerfDisable godoc
// @Summary Set a device perf disable
// @Tags OMC Devices
// @ID set-device-perf-disable
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} echo.Map
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/perf-disable [post]
func HandleSetDevicePerfDisable(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	if !device.IsMethodSupported("SetParameterValues") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	index := cast.ToInt(device.GetProperty("PerfMgmtIndex"))
	if index < 1 {
		index = 1
	}
	prefix := fmt.Sprintf("Device.FAP.PerfMgmt.Config.%d.", index)
	values := map[string]any{
		prefix + "Enable": "0",
	}
	call, err := device.PushMethodCall(time.Now(), "SetParameterValues", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.PushMethodCall(time.Now(), "GetParameterValues", values)
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}

// HandleSetDevicePerfEnable godoc
// @Summary Set a device perf enable
// @Tags OMC Devices
// @ID set-device-perf-enable
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 200 {object} echo.Map
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/perf-enable [post]
func HandleSetDevicePerfEnable(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	if !device.IsMethodSupported("SetParameterValues") {
		return echo.NewHTTPError(http.StatusBadRequest, "method not supported")
	}
	env := utils.GetEnv()
	index := cast.ToInt(device.GetProperty("PerfMgmtIndex"))
	if index < 1 {
		index = 1
	}
	prefix := fmt.Sprintf("Device.FAP.PerfMgmt.Config.%d.", index)
	values := map[string]any{
		prefix + "Enable":                 "1",
		prefix + "PeriodicUploadInterval": "900",
		prefix + "PeriodicUploadTime":     "2011-11-01T00:00:00+08:00",
		prefix + "URL":                    env.GetString("omc_upload_url"),
		prefix + "Username":               env.GetString("omc_upload_username"),
		prefix + "Password":               env.GetString("omc_upload_password"),
	}
	call, err := device.PushMethodCall(time.Now(), "SetParameterValues", values)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	device.PushMethodCall(time.Now(), "GetParameterValues", values)
	device.SendConnectRequest()
	return c.JSON(http.StatusOK, call)
}
