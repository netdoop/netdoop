package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

type listDeviceTransferLogsData struct {
	Data  []omc.DeviceTransferLog `json:"Data"`
	Total int64                   `json:"Total"`
}

// HandleListDeviceTransferLogs lists all transfer logs.
// @Summary List transfer logs
// @ID list-device-transfer-logs
// @Tags OMC Device Transfer Logs
// @Produce json
// @Security Bearer
// @Accept json
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listDeviceTransferLogsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/transfer-logs [get]
func HandleListDeviceTransferLogs(c echo.Context) error {
	c.Set("preload", "S3Object")
	data, total, err := echohandler.ListObjects[omc.DeviceTransferLog](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeviceTransferLogsData{Data: data, Total: total})
}

// HandleGetDeviceTransferLog
// @Summary Get device transfer log
// @ID get-device-transfer-log
// @Tags OMC Device Transfer Logs
// @Produce json
// @Security Bearer
// @Param ts path int true "Timestamp"
// @Success 200 {object} omc.DeviceTransferLog
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/transfer-logs/{ts} [get]
func HandleGetDeviceTransferLog(c echo.Context) error {
	c.Set("preload", "S3Object")
	obj := echohandler.GetObjectFromEchoContext[omc.DeviceTransferLog](c)
	return c.JSON(http.StatusOK, obj)
}

// HandleDeleteDeviceTransferLog
// @Summary Delete device transfer log
// @Tags OMC Device Transfer Logs
// @ID delete-device-transfer-log
// @Security Bearer
// @Accept json
// @Param ts path int true "Timestamp"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/transfer-logs/{ts} [delete]
func HandleDeleteDeviceTransferLog(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DeviceTransferLog](c)
	db := store.GetDB().Unscoped()
	if result := db.Delete(obj); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
