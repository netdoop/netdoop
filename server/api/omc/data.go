package omc

import (
	"net/http"

	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/tsdb"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// HandleQueryData query series.
// @Summary Query data series
// @ID query-data
// @Produce json
// @Security Bearer
// @Param body body tsdb.TSQueryCommand true "Query Command"
// @Success 200 {object} tsdb.TSResult
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/data [post]
// @Tags OMC Data
func HandleQueryData(c echo.Context) error {
	schema := cast.ToString(c.Get("schema"))
	var cmd tsdb.TSQueryCommand
	if err := c.Bind(&cmd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	cmd.Schema = schema

	result, err := tsdb.HandleTSQueryCommand(store.GetDB(), cmd)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
