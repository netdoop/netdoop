package system

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/netdoop/netdoop/server/global"
	"github.com/netdoop/netdoop/utils"
)

type systemInfoData struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Build   string `json:"Build"`
}

// HandleGetSystemInfo get information of system.
// @Summary Get information of system
// @ID get-system-info
// @Produce json
// @Security Bearer
// @Success 200 {object} systemInfoData
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /system/info [get]
// @Tags System
func HandleGetSystemInfo(c echo.Context) error {
	env := utils.GetEnv()
	return c.JSON(http.StatusOK, systemInfoData{
		Name:    env.GetString("app_name"),
		Version: global.Version,
		Build:   global.Build,
	})
}

type systemTimeData struct {
	Current int64 `json:"Current"`
}

// HandleGetSystemTime get current time of system.
// @Summary Get current time of system
// @ID get-system-time
// @Produce json
// @Security Bearer
// @Success 200 {object} systemTimeData
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /system/time [get]
// @Tags System
func HandleGetSystemTime(c echo.Context) error {
	now := time.Now()
	return c.JSON(http.StatusOK, systemTimeData{Current: now.Unix()})
}
