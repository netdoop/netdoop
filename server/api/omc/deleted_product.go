package omc

import (
	"fmt"
	"net/http"

	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

type listDeletedProductsData struct {
	Data  []omc.Product `json:"Data"`
	Total int64         `json:"Total"`
}

// HandleListDeletedProducts lists all products.
// @Summary List deleted products
// @Tags OMC Deleted Products
// @ID list-deleted-products
// @Produce json
// @Security Bearer
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param order_by query string false "Sort order" default()
// @Param q query string false "Query" default()
// @Success 200 {object} listDeletedProductsData
// @Header 200 {int} X-Total "Total number"
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/deleted-products [get]
func HandleListDeletedProducts(c echo.Context) error {
	data, total, err := echohandler.ListDeletedObjects[omc.Product](store.GetDB(), c, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("X-Total", fmt.Sprintf("%v", total))
	return c.JSON(http.StatusOK, listDeletedProductsData{Data: data, Total: total})
}

// HandleDeleteDeletedProduct deletes the product.
// @Summary Delete deleted product
// @Tags OMC Deleted Products
// @ID delete-deleted-product
// @Security Bearer
// @Param id path int true "Product ID"
// @Success 204
// @Failure 401 {object} echo.HTTPError "Unauthorized"
// @Failure 500 {object} echo.HTTPError "Internal Server error"
// @Router /omc/deleted-products/{id} [delete]
func HandleDeleteDeletedProduct(c echo.Context) error {
	product := echohandler.GetObjectFromEchoContext[omc.Product](c)
	db := store.GetDB().Unscoped()
	db = db.Select(clause.Associations)
	if result := db.Delete(product); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.NoContent(http.StatusNoContent)
}
