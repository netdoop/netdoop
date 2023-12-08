package omc

import (
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/jsontype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// HandleListDataModelParameters godoc
// @Summary List the parameters of a datamodel
// @Tags OMC DataModel
// @ID list-datamodel-parameters
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Device ID"
// @Success 200 {object} omc.ParameterValues
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/datamodels/{id}/parameters [get]
func HandleListDataModelParameters(c echo.Context) error {
	c.Set("preload", "DataModelParameter")
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	data := obj.Parameters
	if data == nil {
		data = []*omc.DataModelParameter{}
	}
	return c.JSON(http.StatusOK, data)
}

type createDataModelParameterBody struct {
	Name         string
	Type         string
	Writable     bool
	Description  string
	DefaultValue string
	MetaData     *jsontype.Tags
}

// HandleCreateDataModelParameter creates a new data model parameter.
// @Summary Create data model parameter
// @ID create-datamodel-parameter
// @Produce json
// @Security Bearer
// @Param body body createDataModelParameterBody true "DataModelParameter"
// @Success 200 {object} omc.DataModelParameter
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/parameters [post]
// @Tags OMC DataModel Parameter
func HandleCreateDataModelParameter(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	var data createDataModelParameterBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	db := store.GetDB()
	param := omc.DataModelParameter{
		DataModelID:  obj.ID,
		Default:      false,
		Name:         data.Name,
		Type:         data.Type,
		Description:  data.Description,
		DefaultValue: data.DefaultValue,
		MetaData:     data.MetaData,
	}

	result := db.Create(&param)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, obj)
}

// HandleGetDataModelParameter retrieves a single data model parameter.
// @Summary Get data model parameter
// @ID get-datamodel-parameter
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param parameter_id path string true "Parameter ID"
// @Success 200 {object} omc.DataModelParameter
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/parameters/{parameter_id} [get]
// @Tags OMC DataModel Parameter
func HandleGetDataModelParameter(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	if obj == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	parameterId := cast.ToUint(c.Param("parameter_id"))
	db := store.GetDB()
	param := omc.GetDataModelParameterByID(db, obj, parameterId)
	if param == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, param)
}

// HandleDeleteDataModelParameter deletes the data model parameter.
// @Summary Delete data model parameter
// @Tags OMC DataModel Parameter
// @ID delete-datamodel-parameter
// @Security Bearer
// @Param id path int true "ID"
// @Param parameter_id path string true "Parameter ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/parameter/{parameter_id} [delete]
func HandleDeleteDataModelParameter(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	if obj == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	parameterId := cast.ToUint(c.Param("parameter_id"))
	db := store.GetDB()
	param := omc.GetDataModelParameterByID(db, obj, parameterId)
	if param == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if result := db.Delete(param); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
