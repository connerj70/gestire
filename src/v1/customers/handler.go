package customers

import (
	"fmt"
	"github.com/labstack/echo"
)

func GET(c echo.Context) error {
	fmt.Println("HI")

	return nil
}