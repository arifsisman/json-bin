package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewRecord(c echo.Context) error {
	jsonFromRequest := c.FormValue("json")

	log.Printf("jsonFromRequest", jsonFromRequest)

	if !json.Valid([]byte(jsonFromRequest)) {
		return c.String(http.StatusBadRequest, "Not a valid json!")
	}

	uuid := uuid.New()
	return c.String(http.StatusOK, uuid.String())
}
