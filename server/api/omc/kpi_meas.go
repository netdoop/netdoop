package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type listKPIMeasuresData struct {
	Data  []omc.KPIMeas `json:"Data"`
	Total int64         `json:"Total"`
}

// HandleListKPIMeasures lists all kpi measures.
// @Summary List kpi measures
// @ID list-kpi-measures
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listKPIMeasuresData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/measures [get]
// @Tags OMC KPI
func HandleListKPIMeasures(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.KPIMeas](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listKPIMeasuresData{Data: data, Total: total})
}

type createKPIMeasureBody struct {
	ProductType string
	MeasTypeID  string
	MeasTypeSet string
	Name        string
	Unit        string
	Formula     string
	StatsType   string
}

// HandleCreateKPIMeasure creates a new kpi measure.
// @Summary Create kpi measure
// @ID create-kpi-measure
// @Produce json
// @Security Bearer
// @Param body body createKPIMeasureBody true "KPIMeas"
// @Success 200 {object} omc.KPIMeas
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/measures [post]
// @Tags OMC KPI
func HandleCreateKPIMeasure(c echo.Context) error {
	var data createKPIMeasureBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	if data.ProductType != "enb" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid product type")
	}
	if data.MeasTypeID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid measure type id")
	}
	db := store.GetDB()
	meas := omc.KPIMeas{
		Schema:      cast.ToString(c.Get("schema")),
		ProductType: data.ProductType,
		MeasTypeID:  data.MeasTypeID,
		MeasTypeSet: data.MeasTypeSet,
		Name:        data.Name,
		Unit:        data.Unit,
		Formula:     data.Formula,
		StatsType:   data.StatsType,
	}

	result := db.Create(&meas)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, meas)
}

// HandleGetKPIMeasure retrieves a single kpi measure.
// @Summary Get kpi measure
// @ID get-kpi-measure
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Success 200 {object} omc.KPIMeas
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/measures/{id} [get]
// @Tags OMC KPI
func HandleGetKPIMeasure(c echo.Context) error {
	meas := echohandler.GetObjectFromEchoContext[omc.KPIMeas](c)
	return c.JSON(http.StatusOK, meas)
}

type updateKPIMeasureInfoBody struct {
	Name string
}

// HandleUpdateKPIMeasureInfo updates a kpi measure info
// @Summary Update kpi measure info
// @ID update-kpi-measure-info
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param body body updateKPIMeasureInfoBody true "KPIMeas"
// @Success 200 {object} omc.KPIMeas
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/measures/{id} [put]
// @Tags OMC KPI
func HandleUpdateKPIMeasureInfo(c echo.Context) error {
	var data updateKPIMeasureInfoBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	updateColumns := []string{"manufacturer"}
	updateData := &omc.KPIMeas{
		Name: data.Name,
	}

	meas := echohandler.GetObjectFromEchoContext[omc.KPIMeas](c)
	db := store.GetDB()
	if result := db.Model(meas).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, meas)
}

// HandleSetKPIMeasureDisable godoc
// @Summary Set kpi measure disable
// @Tags OMC KPI
// @ID set-kpi-measure-disable
// @Security Bearer
// @Param id path int true "ID"
// @Success 200 {object} omc.KPIMeas
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/kpi/measures/{id}/disable [put]
func HandleSetKPIMeasureDisable(c echo.Context) error {
	meas := echohandler.GetObjectFromEchoContext[omc.KPIMeas](c)
	if meas.Default {
		return echo.NewHTTPError(http.StatusForbidden, "default kpi measure")
	}
	if meas.Enable {
		db := store.GetDB()
		if result := db.Model(meas).Update("enable", false); result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
		}
	}
	return c.JSON(http.StatusOK, meas)
}

// HandleSetKPIMeasureEnable godoc
// @Summary Set kpi measure enable
// @Tags OMC KPI
// @ID set-kpi-measure-enable
// @Security Bearer
// @Param id path int true "ID"
// @Success 200 {object} omc.KPIMeas
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/kpi/measures/{id}/enable [put]
func HandleSetKPIMeasureEnable(c echo.Context) error {
	meas := echohandler.GetObjectFromEchoContext[omc.KPIMeas](c)
	if meas.Default {
		return echo.NewHTTPError(http.StatusForbidden, "default kpi measure")
	}
	if !meas.Enable {
		db := store.GetDB()
		if result := db.Model(meas).Update("actenableive", true); result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
		}
	}
	return c.JSON(http.StatusOK, meas)
}

// HandleDeleteKPIMeasure deletes the kpi measure.
// @Summary Delete kpi measure
// @Tags OMC KPI
// @ID delete-kpi-measure
// @Security Bearer
// @Param id path int true "ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/measures/{id} [delete]
func HandleDeleteKPIMeasure(c echo.Context) error {
	meas := echohandler.GetObjectFromEchoContext[omc.KPIMeas](c)
	db := store.GetDB().Unscoped()
	if result := db.Unscoped().Delete(meas); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
