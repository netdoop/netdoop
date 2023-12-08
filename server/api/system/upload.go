package system

import (
	"io"
	"net/http"

	"github.com/heypkg/s3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func HandleUpload(c echo.Context) error {
	req := c.Request()
	contentType := req.Header.Get("Content-Type")

	var (
		src      io.Reader
		err      error
		filename string
	)
	switch contentType {
	case "text/plain":
		filename = "file"
		defer req.Body.Close()
		src = req.Body
	case "multipart/form-data":
		file, err := c.FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "read Object").Error())
		}
		filename = file.Filename
		f, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "open Object").Error())
		}
		defer f.Close()
		src = f
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "invalid content type "+contentType)
	}

	object, err := s3.PutObject("", "debug", "test", filename, src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "put object").Error())
	}
	return c.JSON(http.StatusOK, object)
}
