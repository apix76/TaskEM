package rest

import (
	"TaskEM/conf"
	"TaskEM/entities"
	"TaskEM/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Server_Echo(conf conf.Conf, use *usecase.Usecase) {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}

		result, err := use.Post(*u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("/", func(c echo.Context) error {
		var u map[string]string
		if err := c.Bind(u); err != nil {
			return err
		}
		users, err := use.Get(u)
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
		use.Patch(*u)

		return c.JSON(http.StatusOK, u)
	})

	e.DELETE("/", func(c echo.Context) error {
		u := new(entities.UserRequest)
		if err := c.Bind(u); err != nil {
			return err
		}
		err := use.Delete(u.Id)

		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(conf.HttpPort))
}
