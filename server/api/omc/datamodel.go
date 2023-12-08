package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/jsontype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type listDataModelData struct {
	Data  []omc.DataModel `json:"Data"`
	Total int64           `json:"Total"`
}

// HandleListDataModel lists all data models.
// @Summary List data models
// @ID list-datamodels
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listDataModelData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels [get]
// @Tags OMC DataModel
func HandleListDataModel(c echo.Context) error {
	db := store.GetDB()
	data, total, err := echohandler.ListObjects[omc.DataModel](db, c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDataModelData{Data: data, Total: total})
}

type createDataModelBody struct {
	ProductType   string
	Name          string
	ParameterPath string
	MetaData      *jsontype.Tags
}

// HandleCreateDataModel creates a new data model.
// @Summary Create data model
// @ID create-datamodel
// @Produce json
// @Security Bearer
// @Param body body createDataModelBody true "DataModel"
// @Success 200 {object} omc.DataModel
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels [post]
// @Tags OMC DataModel
func HandleCreateDataModel(c echo.Context) error {
	var data createDataModelBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	if data.ProductType != "enb" && data.ProductType != "cpe" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid product type")
	}
	if data.ParameterPath == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter path")
	}
	db := store.GetDB()
	obj := omc.DataModel{
		Default:       false,
		Schema:        cast.ToString(c.Get("schema")),
		ProductType:   data.ProductType,
		Name:          data.Name,
		ParameterPath: data.ParameterPath,
		MetaData:      data.MetaData,
	}

	result := db.Create(&obj)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, obj)
}

// HandleGetDataModel retrieves a single data model.
// @Summary Get data model
// @ID get-datamodel
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Success 200 {object} omc.DataModel
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id} [get]
// @Tags OMC DataModel
func HandleGetDataModel(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	return c.JSON(http.StatusOK, obj)
}

type updateDataModelInfoBody struct {
	Name string
}

// HandleDeleteDataModel deletes the data model.
// @Summary Delete data model
// @Tags OMC DataModel
// @ID delete-datamodel
// @Security Bearer
// @Param id path int true "ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id} [delete]
func HandleDeleteDataModel(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	db := store.GetDB()
	if result := db.Delete(obj); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
