package controllers

import (
	"encoding/json"
	"json-bin/firebase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewBin(c echo.Context) error {
	jsonFromRequest := c.FormValue("json")

	if !json.Valid([]byte(jsonFromRequest)) {
		return c.String(http.StatusBadRequest, "Not a valid json!")
	}

	id := firebase.NewBin(jsonFromRequest)
	return c.String(http.StatusOK, id)
}
