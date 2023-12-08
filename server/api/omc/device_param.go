package omc

import (
	"net/http"

	"github.com/netdoop/netdoop/models/omc"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
)

// HandleGetDeviceMethods godoc
// @Summary Get the methods of a device
// @Tags OMC Devices
// @ID list-device-methods
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Device ID"
// @Success 200 {object} omc.Methods
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/methods [get]
func HandleGetDeviceMethods(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	data := device.Methods
	if data == nil {
		data = &omc.Methods{}
	}
	return c.JSON(http.StatusOK, data)
}

// HandleGetDeviceParameters godoc
// @Summary Get the parameters of a device
// @Tags OMC Devices
// @ID list-device-parameters
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Device ID"
// @Success 200 {object} omc.ParameterValues
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/devices/{id}/parameters [get]
func HandleGetDeviceParameters(c echo.Context) error {
	device := echohandler.GetObjectFromEchoContext[omc.Device](c)

	data := device.ParameterValues
	if data == nil {
		data = &omc.ParameterValues{}
	}
	return c.JSON(http.StatusOK, data)
}
