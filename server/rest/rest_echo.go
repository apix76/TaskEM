package rest

import (
	"TaskEM/conf"
	"TaskEM/entities"
	"TaskEM/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Server_Echo(conf conf.Conf) {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}

		err := usecase.Post(*u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}
		users, err := usecase.Get(*u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
	})

	e.PATCH("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}
		usecase.Patch(*u)

		return c.JSON(http.StatusOK, u)
	})

	e.DELETE("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}
		usecase.Delete(*u)

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(conf.HttpPort))
}
