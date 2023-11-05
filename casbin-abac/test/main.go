package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	enforcer.EnableLog(true)

	// 任意の関数をCasbinに登録する
	enforcer.AddFunction("is_assigned_patient", func(args ...interface{}) (interface{}, error) {
		doctorID, ok := args[0].(float64)
		if !ok {
			return "", fmt.Errorf("doctor ID is invalid")
		}

		patientID, ok := args[1].(float64)
		if !ok {
			return "", fmt.Errorf("patient ID is invalid")
		}

		return isPatientAssignedToDoctor(int(doctorID), int(patientID)), nil
	})

	e.Use(authorize)
	e.GET("/patient/:patient_id/record", func(c echo.Context) error {
		return c.String(http.StatusOK, "Access Allowed\n\n")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	ID int
}

type Object struct {
	ID   int
	Path string
}

func authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		method := c.Request().Method

		userIDStr := c.QueryParam("user_id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		objectIDStr := c.Param("patient_id")
		objectID, err := strconv.Atoi(objectIDStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		obj := Object{ID: objectID, Path: path}

		ok, err := enforcer.Enforce(User{ID: userID}, obj, method)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if ok {
			return next(c)
		}
		return c.String(http.StatusForbidden, "Access denied\n\n")
	}
}

func isPatientAssignedToDoctor(doctorID int, patientID int) bool {
	if doctorID == 1 && patientID == 1 {
		return true
	}
	return false
}
