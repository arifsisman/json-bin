package helpers

import (
	"encoding/json"

	"github.com/google/uuid"
)

func IsValidJSON(j string) bool {
	return json.Valid([]byte(j))
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
