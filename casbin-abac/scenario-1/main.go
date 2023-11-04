package main

import (
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

var enforcer *casbin.Enforcer

func main() {
	e := echo.New()

	var err error
	enforcer, err = casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	e.Use(authorize)

	setRouter(e)
	e.Logger.Fatal(e.Start(":1323"))

}

func setRouter(e *echo.Echo) {
	e.POST("/buy_alcohol", func(c echo.Context) error {
		return c.String(http.StatusOK, "buy alcohol\n")
	})
}

func authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.QueryParam("user_name")
		user, err := getUserByUserName(userName)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ok, err := enforcer.Enforce(user, c.Path(), c.Request().Method)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if ok {
			return next(c)
		} else {
			return echo.ErrForbidden
		}
	}
}

type User struct {
	Name string
	Age  int
}

func getUserByUserName(userName string) (*User, error) {
	users := []User{
		{Name: "alice", Age: 17},
		{Name: "bob", Age: 22},
		{Name: "cathy", Age: 30},
	}

	for _, user := range users {
		if user.Name == userName {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
