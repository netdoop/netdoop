package omc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/tsdb"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type listKPITemplatesData struct {
	Data  []omc.KPITemplate `json:"Data"`
	Total int64             `json:"Total"`
}

// HandleListKPITemplates lists all kpi templates.
// @Summary List kpi templates
// @ID list-kpi-templates
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listKPITemplatesData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates [get]
// @Tags OMC KPI
func HandleListKPITemplates(c echo.Context) error {
	data, total, err := echohandler.ListObjects[omc.KPITemplate](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listKPITemplatesData{Data: data, Total: total})
}

type createKPITemplateBody struct {
	Name             string
	SelectType       string
	SelectIds        []uint
	MeasTypeIds      []string
	PeriodicInterval int64
}

// HandleCreateKPITemplate creates a new kpi template.
// @Summary Create kpi template
// @ID create-kpi-template
// @Produce json
// @Security Bearer
// @Param body body createKPITemplateBody true "KPITemplate"
// @Success 200 {object} omc.KPITemplate
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates [post]
// @Tags OMC KPI
func HandleCreateKPITemplate(c echo.Context) error {
	var data createKPITemplateBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	if data.SelectType != "Device" && data.SelectType != "DeviceGroup" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid select type")
	}
	if len(data.SelectIds) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "empty select ids")
	}
	if len(data.MeasTypeIds) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "empty measure type ids")
	}
	db := store.GetDB()
	temp := omc.KPITemplate{
		Schema:           cast.ToString(c.Get("schema")),
		Name:             data.Name,
		SelectType:       data.SelectType,
		SelectIds:        data.SelectIds,
		MeasTypeIds:      data.MeasTypeIds,
		PeriodicInterval: data.PeriodicInterval,
	}

	result := db.Create(&temp)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	where := ""
	if data.SelectType == "Device" && len(data.SelectIds) > 0 {
		values := []string{}
		for _, v := range data.SelectIds {
			values = append(values, fmt.Sprintf("'%v'", v))
		}
		where = fmt.Sprintf("device_id IN (%v)", strings.Join(values, ","))
	}
	if err := tsdb.CreateHyperTableAvgValuesView(db,
		"device_performance_values",
		temp.ViewTableName(),
		fmt.Sprintf("%vm", temp.PeriodicInterval/60),
		"name", "value",
		[]string{"schema", "device_id", "product_type", "oui", "product_class", "serial_number"},
		temp.MeasTypeIds,
		where,
	); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, temp)
}

// HandleGetKPITemplate retrieves a single kpi template.
// @Summary Get kpi template
// @ID get-kpi-template
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Success 200 {object} omc.KPITemplate
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates/{id} [get]
// @Tags OMC KPI
func HandleGetKPITemplate(c echo.Context) error {
	temp := echohandler.GetObjectFromEchoContext[omc.KPITemplate](c)
	return c.JSON(http.StatusOK, temp)
}

type updateKPITemplateBody struct {
	SelectType       string
	SelectIds        []uint
	MeasTypeIds      []string
	PeriodicInterval int64
}

// HandleUpdateKPITemplate updates a kpi template
// @Summary Update kpi template
// @ID update-kpi-template
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param body body updateKPITemplateBody true "KPITemplate"
// @Success 200 {object} omc.KPITemplate
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates/{id} [put]
// @Tags OMC KPI
func HandleUpdateKPITemplateInfo(c echo.Context) error {
	var data updateKPITemplateBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	updateColumns := []string{"periodic_interval", "select_type", "select_ids", "meas_type_ids"}
	updateData := &omc.KPITemplate{
		SelectType:       data.SelectType,
		SelectIds:        data.SelectIds,
		MeasTypeIds:      data.MeasTypeIds,
		PeriodicInterval: data.PeriodicInterval,
	}
	updateData.SaveData()

	temp := echohandler.GetObjectFromEchoContext[omc.KPITemplate](c)
	db := store.GetDB()
	if result := db.Model(temp).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, temp)
}

// HandleDeleteKPITemplate deletes the kpi template.
// @Summary Delete kpi template
// @Tags OMC KPI
// @ID delete-kpi-template
// @Security Bearer
// @Param id path int true "ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates/{id} [delete]
func HandleDeleteKPITemplate(c echo.Context) error {
	temp := echohandler.GetObjectFromEchoContext[omc.KPITemplate](c)
	db := store.GetDB().Unscoped()
	if result := db.Unscoped().Delete(temp); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	tsdb.DropHyperTableView(db, fmt.Sprintf("device_kpi_template_%v_view", temp.Name))
	return c.NoContent(http.StatusNoContent)
}

type listKPITemplateRecordsData struct {
	Data  []map[string]any `json:"Data"`
	Total int64            `json:"Total"`
}

// HandleListKPITemplateRecords lists all kpi templates.
// @Summary List kpi template records
// @ID list-kpi-template-records
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listKPITemplateRecordsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/kpi/templates/{id}/records [get]
// @Tags OMC KPI
func HandleListKPITemplateRecords(c echo.Context) error {
	temp := echohandler.GetObjectFromEchoContext[omc.KPITemplate](c)
	data, total, err := echohandler.ListAnyObjects(store.GetDB(), c, temp.ViewTableName(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listKPITemplateRecordsData{Data: data, Total: total})
}
