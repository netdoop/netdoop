package omc

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/s3"
	"github.com/heypkg/store/echohandler"
	"github.com/heypkg/store/jsontype"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type listFirmwaresData struct {
	Data  []omc.Firmware `json:"Data"`
	Total int64          `json:"Total"`
}

// HandleListFirmwares lists all firmwares.
// @Summary List firmwares
// @ID list-firmwares
// @Produce json
// @Security Bearer
// @Accept json
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listFirmwaresData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/firmwares [get]
// @Tags OMC Firmwares
func HandleListFirmwares(c echo.Context) error {
	c.Set("preload", "S3Object,Products")
	data, total, err := echohandler.ListObjects[omc.Firmware](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listFirmwaresData{Data: data, Total: total})
}

// HandleGetFirmware retrieves a single firmware.
// @Summary Get firmware
// @ID get-firmware
// @Produce json
// @Security Bearer
// @Param id path int true "Firmware ID"
// @Success 200 {object} omc.Firmware
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/firmwares/{id} [get]
// @Tags OMC Firmwares
func HandleGetFirmware(c echo.Context) error {
	c.Set("preload", "S3Object,Products")
	firmware := echohandler.GetObjectFromEchoContext[omc.Firmware](c)
	return c.JSON(http.StatusOK, firmware)
}

// HandleCreateFirmware creates a new firmware.
// @Summary Create firmware
// @ID create-firmware
// @Produce json
// @Security Bearer
// @Accept  multipart/form-data
// @Param Version formData string true "Version"
// @Param ProductType formData string true "ProductType"
// @Param Products formData string true "Product list"
// @Param File formData file true "Firmware file to upload"
// @Success 200 {object} omc.Firmware
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/firmwares [post]
// @Tags OMC Firmwares
func HandleCreateFirmware(c echo.Context) error {
	schema := cast.ToString(c.Get("schema"))
	version := c.FormValue("Version")
	productType := c.FormValue("ProductType")

	modelNames := strings.Split(c.FormValue("Products"), ",")

	db := store.GetDB()
	products := []*omc.Product{}
	for _, model := range modelNames {
		if model != "" {
			product := omc.GetProductByModelName(db, schema, model)
			if product == nil {
				return echo.NewHTTPError(http.StatusBadRequest, "unknow product model name")
			}
			if product.ProductType != productType {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid product type of product model")
			}
			products = append(products, product)
		}
	}

	key := version
	file, err := c.FormFile("File")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "read Object").Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "open Object").Error())
	}
	defer src.Close()

	object, err := s3.PutObject(schema, omc.FirmwareBucket, key, file.Filename, src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "put object").Error())
	}

	firmware := omc.Firmware{
		Schema:      schema,
		Version:     version,
		Name:        object.FileName,
		Uploader:    cast.ToString(c.Get("loginName")),
		UploadTime:  jsontype.JSONTime(time.Now()),
		S3ObjectID:  object.ID,
		ProductType: productType,
		Products:    products,
	}

	result := db.Create(&firmware)
	if result.Error != nil {
		s3.RemoveObjectById(schema, object.ID)
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, firmware)
}

// HandleDeleteFirmware deletes the firmware.
// @Summary Delete firmware
// @Tags OMC Firmwares
// @ID delete-firmware
// @Security Bearer
// @Accept json
// @Param id path int true "Firmware ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/firmwares/{id} [delete]
func HandleDeleteFirmware(c echo.Context) error {
	firmware := echohandler.GetObjectFromEchoContext[omc.Firmware](c)

	db := store.GetDB().Unscoped()
	if result := db.Select("Products").Delete(firmware); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}

	return c.NoContent(http.StatusNoContent)
}

type setFirmwareProductsBody struct {
	ModelNames []string `json:"ModelNames"`
}

// HandleSetFirmwareProducts godoc
// @Summary Set firmware products
// @Tags OMC Firmwares
// @ID set-firmware-products
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Device ID"
// @Param body body setFirmwareProductsBody true "Body"
// @Success 200 {object} omc.Device
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/firmwares/{id}/products [put]
func HandleSetFirmwareProducts(c echo.Context) error {
	var data setFirmwareProductsBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}

	schema := cast.ToString(c.Get("schema"))

	db := store.GetDB()
	products := []*omc.Product{}
	for _, model := range data.ModelNames {
		if product := omc.GetProductByModelName(db, schema, model); product != nil {
			products = append(products, product)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "unknow product modal name")
		}
	}

	firmware := echohandler.GetObjectFromEchoContext[omc.Firmware](c)
	updateColumns := []string{"products"}
	updateData := &omc.Firmware{
		Products: products,
	}
	if result := db.Model(firmware).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, firmware)
}
