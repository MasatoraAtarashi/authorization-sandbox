package main

import (
	"fmt"
	"net/http"
	"time"

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

	enforcer.AddFunction("working_hours", func(args ...interface{}) (interface{}, error) {
		now := time.Now()
		workingStart := 9
		workingEnd := 18
		return now.Hour() >= workingStart && now.Hour() < workingEnd, nil
	})

	enforcer.AddFunction("is_assigned_patient", func(args ...interface{}) (interface{}, error) {
		doctorID, ok := args[0].(string)
		if !ok {
			return "", fmt.Errorf("doctor ID is invalid")
		}

		patientID, ok := args[1].(string)
		if !ok {
			return "", fmt.Errorf("patient ID is invalid")
		}

		return isPatientAssignedToDoctor(doctorID, patientID), nil
	})

	e.Use(authorize)
	e.GET("/patient/:patient_id/record", func(c echo.Context) error {
		return c.String(http.StatusOK, "Access Allowed\n\n")
	})
	e.GET("/patient/:patient_id/record/summary", func(c echo.Context) error {
		return c.String(http.StatusOK, "Access Allowed\n\n")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := getUserByName(c.Request().Header.Get("user"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		path := c.Request().URL.Path
		method := c.Request().Method
		objectID := c.Param("patient_id")
		obj := Object{ID: objectID, Path: path}

		ok, err := enforcer.Enforce(user, obj, method)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if ok {
			return next(c)
		}
		return c.String(http.StatusForbidden, "Access denied\n\n")
	}
}

type User struct {
	ID   string
	Role string
}

type Object struct {
	ID   string
	Path string
}

func getUserByName(name string) (*User, error) {
	users := []User{
		{ID: "1", Role: "doctor"},
		{ID: "2", Role: "nurse"},
		{ID: "3", Role: "family"},
	}

	for _, user := range users {
		if user.Role == name {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func isPatientAssignedToDoctor(doctorID string, patientID string) bool {
	if doctorID == "1" && patientID == "123" {
		return true
	}
	return false
}
