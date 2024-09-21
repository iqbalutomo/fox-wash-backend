package helpers

import (
	"email_service/models"
	"encoding/json"
	"log"
)

func AssertJsonToStruct(body []byte) models.UserCredential {
	var credential models.UserCredential

	if err := json.Unmarshal(body, &credential); err != nil {
		log.Fatalf("failed to unmarshaling data: %v", err)
	}

	return credential
}
