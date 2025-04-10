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
		cond := entities.Cond{}
		b := echo.QueryParamsBinder(c)

		b.String("id", cond.Id).
			String("name", cond.Name).
			String("surname", cond.Surname).
			String("patronymic", cond.Patronymic).
			Int("ageGt", cond.AgeGt).
			Int("ageLt", cond.AgeLt).
			String("gender", cond.Gender).
			String("race", cond.Race).BindErrors()

		users, err := use.Get(&cond)
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
		user, err := use.Patch(*u)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
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
	//e.StartTLS(conf)
}
