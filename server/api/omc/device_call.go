package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

type listDeviceMethodCallsData struct {
	Data  []omc.DeviceMethodCall `json:"Data"`
	Total int64                  `json:"Total"`
}

// HandleListDeviceMethodCalls lists all device method calls.
// @Summary List device method calls
// @ID list-device-method-calls
// @Produce json
// @Security Bearer
// @Param offset query int false "Offset" default(0)
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param q query string false "Query" default()
// @Param order_by query string false "Sort order" default()
// @Success 200 {object} listDeviceMethodCallsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-method-calls [get]
// @Tags OMC Device MethodCalls
func HandleListDeviceMethodCalls(c echo.Context) error {
	selectNames := []string{
		"time", "schema", "device_id", "oui", "product_class", "serial_number", "product_type",
		"updated", "state", "command_key", "method_name", "fault_code", "fault_string",
	}
	data, total, err := echohandler.ListObjects[omc.DeviceMethodCall](store.GetDB(), c, selectNames, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeviceMethodCallsData{Data: data, Total: total})
}

// HandleGetDeviceMethodCall retrieves a single device method call.
// @Summary Get device method call
// @ID get-device-method-call
// @Produce json
// @Security Bearer
// @Param ts path int true "Timestamp"
// @Success 200 {object} omc.DeviceMethodCall
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-method-calls/{ts} [get]
// @Tags OMC Device MethodCalls
func HandleGetDeviceMethodCall(c echo.Context) error {
	call := echohandler.GetObjectFromEchoContext[omc.DeviceMethodCall](c)
	return c.JSON(http.StatusOK, call)
}
