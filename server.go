package main

import (
	"github.com/labstack/echo"
	"github.com/gestire/src/v1/customers"
)

func main() {

	e := echo.New()

	//CUSTOMERS
	e.GET("/v1/customers", customers.GET)

	e.Logger.Fatal(e.Start(":1323"))
}