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
	Loader(c echo.Context) error
	GetInfo(c echo.Context) error
	// Tester(c echo.Context) error
}

type loadController struct {
	service LoadServiceInterface
}

func NewloadController(ser LoadServiceInterface) LoadControllerInterface {
	return &loadController{
		ser,
	}
}
func (controller loadController) Loader(c echo.Context) error {
	url := c.QueryParam("url")
	loads, err3 := controller.service.Loader(url)
	if err3 != nil {
		return c.JSON(err3.Code(), err3)
	}
	return c.JSON(http.StatusOK, loads)
}
func (controller loadController) GetInfo(c echo.Context) error {
	loads, err3 := controller.service.GetInfo()
	if err3 != nil {
		return c.JSON(err3.Code(), err3)
	}
	return c.JSON(http.StatusOK, loads)
}
// func (controller loadController) Tester(c echo.Context) error {
// 	loads, err3 := controller.service.Tester()
// 	if err3 != nil {
// 		return c.JSON(err3.Code(), err3)
// 	}
// 	return c.JSON(http.StatusOK, loads)
// }

