package internal

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

func Generate() map[string]string {
	uuid_v7, err := uuid.NewV7()
	if err != nil {
		log.Fatal("ERROR: Could not generate UUID")
	}

	generated_uuid := strings.Split(uuid_v7.String(), "-")
	result := map[string]string{
		"timehigh": generated_uuid[0],
		"timemid":  generated_uuid[1],
		"version":  generated_uuid[2],
		"variant":  generated_uuid[3],
		"random":   generated_uuid[4],
	}

	return result
}
