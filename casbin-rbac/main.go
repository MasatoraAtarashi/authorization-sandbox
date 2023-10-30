package main

import (
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
	deviceApi := e.Group("/device")
	deviceApi.GET("/list", func(c echo.Context) error {
		return c.String(http.StatusOK, "device list\n")
	})
	deviceApi.GET("/:deviceID", func(c echo.Context) error {
		return c.String(http.StatusOK, "device detail\n")
	})
	deviceApi.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "device update\n")
	})
	deviceApi.DELETE("/:deviceID", func(c echo.Context) error {
		return c.String(http.StatusOK, "device delete\n")
	})

	adminApi := e.Group("/admin")
	adminApi.GET("/user/list", func(c echo.Context) error {
		return c.String(http.StatusOK, "admin user list\n")
	})
}

func authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.QueryParam("user_name")
		role := getRoleByUserName(userName)
		ok, err := enforcer.Enforce(role, c.Path(), c.Request().Method)
		if err != nil {
			return err
		}

		if ok {
			return next(c)
		} else {
			return echo.ErrForbidden
		}
	}
}

func getRoleByUserName(userName string) string {
	var userRoleMap = map[string]string{
		"alice": "admin",
		"bob":   "writer",
		"john":  "reader",
	}

	return userRoleMap[userName]
}
