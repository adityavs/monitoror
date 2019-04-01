package http

import (
	"net/http"

	"github.com/jsdidierlaurent/monitoror/models/errors"
	"github.com/jsdidierlaurent/monitoror/monitorable/port/model"

	"github.com/jsdidierlaurent/monitoror/monitorable/port"

	"github.com/labstack/echo/v4"
)

type httpPortHandler struct {
	portUsecase port.Usecase
}

func NewHttpPortHandler(p port.Usecase) *httpPortHandler {
	return &httpPortHandler{p}
}

func (h *httpPortHandler) GetPort(c echo.Context) error {
	// Bind / Validate Params
	params := &model.PortParams{}
	err := c.Bind(params)
	if err != nil || !params.Validate() {
		return errors.NewQueryParamsError(err)
	}

	tile, err := h.portUsecase.Port(params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tile)
}