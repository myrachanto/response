package load

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// loadController ...
var (
	LoadController LoadControllerInterface = loadController{}
	Bizname        string
)

type LoadControllerInterface interface {
	GetURL(c echo.Context) error
}

type loadController struct {
	service LoadServiceInterface
}

func NewloadController(ser LoadServiceInterface) LoadControllerInterface {
	return &loadController{
		ser,
	}
}
func (controller loadController) GetURL(c echo.Context) error {
	url := c.QueryParam("url")
	loads := controller.service.GetURL(url)
	return c.JSON(http.StatusOK, loads)
}
