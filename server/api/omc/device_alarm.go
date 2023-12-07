package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

type listDeviceAlarmsData struct {
	Data  []omc.DeviceAlarm `json:"Data"`
	Total int64             `json:"Total"`
}

// HandleListDeviceAlarms lists all device alarms.
// @Summary List device alarms
// @ID list-device-alarms
// @Produce json
// @Security Bearer
// @Param offset query int false "Offset" default(0)
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param q query string false "Query" default()
// @Param order_by query string false "Sort order" default()
// @Success 200 {object} listDeviceAlarmsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-alarms [get]
// @Tags OMC Device Alarms(
func HandleListDeviceAlarms(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.DeviceAlarm](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeviceAlarmsData{Data: data, Total: total})
}

// HandleGetDeviceAlarm retrieves a single device alarm.
// @Summary Get device alarm
// @ID get-device-alarm
// @Produce json
// @Security Bearer
// @Param ts path int true "Timestamp"
// @Success 200 {object} omc.DeviceAlarm
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-alarms/{ts} [get]
// @Tags OMC Device Alarms
func HandleGetDeviceAlarm(c echo.Context) error {
	alarm := echohandler.GetObjectFromEchoContext[omc.DeviceAlarm](c)
	return c.JSON(http.StatusOK, alarm)
}
