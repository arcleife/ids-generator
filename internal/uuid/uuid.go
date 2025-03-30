package internal

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

type StructuredUUID struct {
	timehigh string
	timemid  string
	version  string
	variant  string
	random   string
}

func New() StructuredUUID {
	value, err := uuid.NewV7()

	if err != nil {
		log.Fatal("Could not create UUID")
	}

	all := value.String()
	generated_uuid := strings.Split(all, "-")

	return StructuredUUID{
		timehigh: generated_uuid[0],
		timemid:  generated_uuid[1],
		version:  generated_uuid[2],
		variant:  generated_uuid[3],
		random:   generated_uuid[4],
	}
}

func (u StructuredUUID) All(delimiter string) string {
	return u.timehigh + delimiter +
		u.timemid + delimiter +
		u.version + delimiter +
		u.variant + delimiter +
		u.random
}

func Bulk(n int, delimiter string) []string {
	bulk := make([]string, n)

	for i := range bulk {
		generated := New().All(delimiter)

		bulk[i] = generated
	}

	return bulk
}
