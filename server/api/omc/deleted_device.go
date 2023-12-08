package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type listDeletedDevicesData struct {
	Data  []omc.Device `json:"Data"`
	Total int64        `json:"Total"`
}

// HandleListDeletedDevices godoc
// @Summary List deleted devices
// @Tags OMC Deleted Devices
// @ID list-deleted-devices
// @Accept json
// @Produce json
// @Param q query string false "Query" default()
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Security Bearer
// @Success 200 {object} listDeletedDevicesData
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/deleted-devices [get]
func HandleListDeletedDevices(c echo.Context) error {
	db := store.GetDB()
	data, total, err := echohandler.ListDeletedObjects[omc.Device](db, c, omc.SimpleDeviceInfoName, omc.DeviceSearchHanleFuncs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeletedDevicesData{Data: data, Total: total})
}

// HandleDeleteDeletedDevice godoc
// @Summary Delete a device permanently
// @Tags OMC Deleted Devices
// @ID delete-deleted-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/deleted-devices/{id} [delete]
func HandleDeleteDeletedDevice(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	db := store.GetDB().Unscoped()
	db = db.Select(clause.Associations)
	if result := db.Delete(device); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}

// HandleRecoverDeletedDevice godoc
// @Summary Recover a deleted device
// @Tags OMC Deleted Devices
// @ID recover-deleted-device
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/deleted-devices/{id}/recover [post]
func HandleRecoverDeletedDevice(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)
	db := store.GetDB().Unscoped()
	if result := db.Model(device).Update("deleted", nil); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
