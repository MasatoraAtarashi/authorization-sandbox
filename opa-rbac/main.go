package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/open-policy-agent/opa/rego"
)

func main() {
	e := echo.New()

	e.Use(authorize)

	setRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func setRouter(e *echo.Echo) {
	articleApi := e.Group("/articles")
	articleApi.GET("/", func(c echo.Context) error {
		return c.String(200, "article list\n")
	})
	articleApi.POST("/", func(c echo.Context) error {
		return c.String(200, "article create\n")
	})
	articleApi.GET("/:articleID", func(c echo.Context) error {
		return c.String(200, "article detail\n")
	})
	articleApi.PUT("/:articleID", func(c echo.Context) error {
		return c.String(200, "article update\n")
	})
	articleApi.DELETE("/:articleID", func(c echo.Context) error {
		return c.String(200, "article delete\n")
	})
}

type Input struct {
	Method string   `json:"method"`
	Path   []string `json:"path"`
	Roles  []string `json:"roles"`
}

type Result struct {
	Result bool `json:"result"`
}

const (
	policyFile = "policy.rego"
	dataFile   = "data.json"
)

func authorize(next echo.HandlerFunc) echo.HandlerFunc {
	ctx := context.Background()
	query, err := rego.New(
		rego.Query("x = data.app.rbac.allow"),
		rego.Load([]string{policyFile, dataFile}, nil),
	).PrepareForEval(ctx)
	if err != nil {
		panic(err)
	}

	return func(c echo.Context) error {
		trimed := strings.Trim(c.Path(), "/")
		p := strings.Split(trimed, "/")
		userName := c.Request().Header.Get("user_name")

		in := Input{
			Method: c.Request().Method,
			Path:   p,
			Roles:  getUserRoles(userName),
		}
		results, err := query.Eval(ctx, rego.EvalInput(in))
		fmt.Printf("%+v\n", results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if len(results) == 0 {
			return c.JSON(http.StatusForbidden, "Access denied")
		}

		allow, ok := results[0].Bindings["x"].(bool)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "error")
		}

		if !allow {
			return c.JSON(http.StatusForbidden, "Access denied")
		}

		return next(c)
	}
}

func getUserRoles(userName string) []string {
	var userRoleMap = map[string][]string{
		"alice": {"admin", "articles.admin"},
		"bob":   {"article.editor"},
		"john":  {"article.viewer"},
	}
	return userRoleMap[userName]
}
