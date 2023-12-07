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

// HandleListDataModelTemplates godoc
// @Summary List the templates of a datamodel
// @Tags OMC DataModel
// @ID list-datamodel-templates
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Device ID"
// @Success 200 {object} omc.ParameterValues
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/datamodels/{id}/templates [get]
func HandleListDataModelTemplates(c echo.Context) error {
	c.Set("preload", "DataModelTemplate")
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	data := obj.Templates
	if data == nil {
		data = []*omc.DataModelTemplate{}
	}
	return c.JSON(http.StatusOK, data)
}

type createDataModelTemplateBody struct {
	Name           string
	ParameterNames *jsontype.Tags
	MetaData       *jsontype.Tags
}

// HandleCreateDataModelTemplate creates a new data model template.
// @Summary Create data model template
// @ID create-datamodel-template
// @Produce json
// @Security Bearer
// @Param body body createDataModelTemplateBody true "DataModelTemplate"
// @Success 200 {object} omc.DataModelTemplate
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input template"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/templates [post]
// @Tags OMC DataModel Template
func HandleCreateDataModelTemplate(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	var data createDataModelTemplateBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input template").Error())
	}
	db := store.GetDB()
	param := omc.DataModelTemplate{
		DataModelID:    obj.ID,
		Default:        false,
		Name:           data.Name,
		ParameterNames: data.ParameterNames,
		MetaData:       data.MetaData,
	}

	result := db.Create(&param)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, obj)
}

// HandleGetDataModelTemplate retrieves a single data model template.
// @Summary Get data model template
// @ID get-datamodel-template
// @Produce json
// @Security Bearer
// @Param id path int true "ID"
// @Param template_id path string true "TemplateID"
// @Success 200 {object} omc.DataModelTemplate
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/templates/{template_id} [get]
// @Tags OMC DataModel Template
func HandleGetDataModelTemplate(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	if obj == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	templateID := cast.ToUint(c.Param("template_id"))
	db := store.GetDB()
	template := omc.GetDataModelTemplate(db, obj, templateID)
	if template == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, template)
}

// HandleDeleteDataModelTemplate deletes the data model template.
// @Summary Delete data model template
// @Tags OMC DataModel Template
// @ID delete-datamodel-template
// @Security Bearer
// @Param id path int true "ID"
// @Param template_id path string true "TemplateID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/datamodels/{id}/template/{template_id} [delete]
func HandleDeleteDataModelTemplate(c echo.Context) error {
	obj := echohandler.GetObjectFromEchoContext[omc.DataModel](c)
	if obj == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	templateID := cast.ToUint(c.Param("template_id"))
	db := store.GetDB()
	template := omc.GetDataModelTemplate(db, obj, templateID)
	if template == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if result := db.Delete(template); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
