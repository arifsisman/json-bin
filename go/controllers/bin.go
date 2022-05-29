package controllers

import (
	"json-bin/firebase"
	"json-bin/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewBin(c echo.Context) error {
	json := c.FormValue("json")

	if !helpers.IsValidJSON(json) {
		return c.String(http.StatusBadRequest, "Not a valid JSON!")
	}

	id := firebase.NewBin(json)
	return c.String(http.StatusOK, id)
}

func GetBin(c echo.Context) error {
	uuid := c.Param("uuid")

	if !helpers.IsValidUUID(uuid) {
		return c.String(http.StatusBadRequest, "Not a valid UUID!")
	}

	json := firebase.GetBin(uuid)
	return c.String(http.StatusOK, json)
}
