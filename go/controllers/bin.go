package controllers

import (
	"io/ioutil"
	"json-bin/firebase"
	"json-bin/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewBin(c echo.Context) error {
	jsonBodyBytes, err := ioutil.ReadAll(c.Request().Body)
	jsonBody := string(jsonBodyBytes)

	if err != nil || !helpers.IsValidJSON(jsonBody) {
		return c.String(http.StatusBadRequest, "Not a valid JSON!")
	}

	id := firebase.NewBin(jsonBody)
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
