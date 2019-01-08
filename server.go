package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	e.POST("/server/save", func(c echo.Context) error {
		response := &jsonResponse{
			Status:  "OK",
			NumRows: "69",
		}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8081"))

}

type jsonResponse struct {
	Status  string `json:"status"`
	NumRows string `json:"num_rows"`
}
