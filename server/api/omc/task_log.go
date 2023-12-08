package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

type listTaskDeviceLogsData struct {
	Data  []omc.TaskDeviceLog `json:"Data"`
	Total int64               `json:"Total"`
}

// HandleListTaskDeviceLogs lists all log device logs.
// @Summary List log device logs
// @ID list-device-logs
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listTaskDeviceLogsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-logs [get]
// @Tags OMC TaskDeviceLogs
func HandleListTaskDeviceLogs(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.TaskDeviceLog](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listTaskDeviceLogsData{Data: data, Total: total})
}

// HandleGetTaskDeviceLog retrieves a single log device log.
// @Summary Get log device log
// @ID get-device-log
// @Produce json
// @Security Bearer
// @Param ts path int true "Timestamp"
// @Success 200 {object} omc.TaskDeviceLog
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-logs/{ts} [get]
// @Tags OMC TaskDeviceLogs
func HandleGetTaskDeviceLog(c echo.Context) error {
	log := echohandler.GetObjectFromEchoContext[omc.TaskDeviceLog](c)
	return c.JSON(http.StatusOK, log)
}
