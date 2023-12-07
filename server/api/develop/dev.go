package develop

import (
	"bytes"
	"net/http"

	"github.com/netdoop/netdoop/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func HandlePost(c echo.Context) error {
	logger := utils.GetLogger().Named("develop")
	body := c.Request().Body
	defer body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	logger.Info("post", zap.String("content", buf.String()))
	return c.JSON(http.StatusOK, map[string]string{})
}
