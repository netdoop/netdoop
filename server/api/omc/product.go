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

type listProductsData struct {
	Data  []omc.Product `json:"Data"`
	Total int64         `json:"Total"`
}

// HandleListProducts lists all products.
// @Summary List products
// @ID list-products
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listProductsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products [get]
// @Tags OMC Products
func HandleListProducts(c echo.Context) error {
	c.Set("preload", "SupportedAlarms")
	data, total, err := echohandler.ListObjects[omc.Product](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listProductsData{Data: data, Total: total})
}

type createProductBody struct {
	Oui          string
	ProductClass string
	Manufacturer string
	ModelName    string
}

// HandleCreateProduct creates a new product.
// @Summary Create product
// @ID create-product
// @Produce json
// @Security Bearer
// @Param body body createProductBody true "Product"
// @Success 200 {object} omc.Product
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products [post]
// @Tags OMC Products
func HandleCreateProduct(c echo.Context) error {
	var data createProductBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}

	db := store.GetDB()
	product := omc.Product{
		Schema:       cast.ToString(c.Get("schema")),
		Oui:          data.Oui,
		ProductClass: data.ProductClass,
		Manufacturer: data.Manufacturer,
		ModelName:    data.ModelName,
	}

	result := db.Create(&product)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, product)
}

// HandleGetProduct retrieves a single product.
// @Summary Get product
// @ID get-product
// @Produce json
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 200 {object} omc.Product
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products/{id} [get]
// @Tags OMC Products
func HandleGetProduct(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	return c.JSON(http.StatusOK, product)
}

// HandleListProductFirmwares
// @Summary List product firmwares
// @ID list-product-firmwares
// @Produce json
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 200 {array} omc.Firmware "Firmwares"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products/{id}/firmwares [get]
// @Tags OMC Products
func HandleListProductFirmwares(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	db := store.GetDB().Model(product).Preload("Firmwares")
	if result := db.Where("schema = ? AND id = ?", product.Schema, product.ID).First(product); result.Error != nil {
		return nil
	}
	return c.JSON(http.StatusOK, product.Firmwares)
}

type updateProductInfoBody struct {
	Manufacturer string
}

// HandleUpdateProductInfo updates a product's name and manufacturer.
// @Summary Update product info
// @ID update-product-info
// @Produce json
// @Security Bearer
// @Param id path int true "Product ID"
// @Param body body updateProductInfoBody true "Product"
// @Success 200 {object} omc.Product
// @Failure 400 {object} echo.HTTPError "Bad Request: invalid input parameter"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products/{id} [put]
// @Tags OMC Products
func HandleUpdateProductInfo(c echo.Context) error {
	var data updateProductInfoBody
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid input parameter").Error())
	}
	updateColumns := []string{"manufacturer"}
	updateData := &omc.Product{
		Manufacturer: data.Manufacturer,
	}

	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	db := store.GetDB()
	if result := db.Model(product).Select(updateColumns).Updates(updateData); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, product)
}

// HandleSetProductDisable godoc
// @Summary Set product disable
// @Tags OMC Products
// @ID set-product-disable
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 200 {object} omc.Product
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/products/{id}/disable [put]
func HandleSetProductDisable(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	if product.Default {
		return echo.NewHTTPError(http.StatusForbidden, "default product")
	}
	if product.Enable {
		db := store.GetDB()
		if result := db.Model(product).Update("enable", false); result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
		}
	}
	return c.JSON(http.StatusOK, product)
}

// HandleSetProductEnable godoc
// @Summary Set product enable
// @Tags OMC Products
// @ID set-product-enable
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 200 {object} omc.Product
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 403 {object} echo.HTTPError "Forbidden"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /omc/products/{id}/enable [put]
func HandleSetProductEnable(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	if product.Default {
		return echo.NewHTTPError(http.StatusForbidden, "default product")
	}
	if !product.Enable {
		db := store.GetDB()
		if result := db.Model(product).Update("actenableive", true); result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
		}
	}
	return c.JSON(http.StatusOK, product)
}

// HandleDeleteProduct deletes the product.
// @Summary Delete product
// @Tags OMC Products
// @ID delete-product
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/products/{id} [delete]
func HandleDeleteProduct(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	db := store.GetDB()
	if result := db.Delete(product); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
