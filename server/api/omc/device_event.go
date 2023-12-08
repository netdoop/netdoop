package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

type listDeviceEventsData struct {
	Data  []omc.DeviceEvent `json:"Data"`
	Total int64             `json:"Total"`
}

// HandleListDeviceEvents lists all device events.
// @Summary List device events
// @ID list-device-events
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listDeviceEventsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-events [get]
// @Tags OMC DeviceEvents
func HandleListDeviceEvents(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.DeviceEvent](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeviceEventsData{Data: data, Total: total})
}

// HandleGetDeviceEvent retrieves a single device event.
// @Summary Get device event
// @ID get-device-event
// @Produce json
// @Security Bearer
// @Param ts path int true "Timestamp"
// @Success 200 {object} omc.DeviceEvent
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/device-events/{ts} [get]
// @Tags OMC DeviceEvents
func HandleGetDeviceEvent(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DeviceEvent](c)
	return c.JSON(http.StatusOK, obj)
}
