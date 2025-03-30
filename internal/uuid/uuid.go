package internal

import (
	"strings"

	"github.com/google/uuid"
)

type StructuredUUID struct {
	all      string
	timehigh string
	timemid  string
	version  string
	variant  string
	random   string
}

func New(uuid uuid.UUID) StructuredUUID {
	all := uuid.String()
	generated_uuid := strings.Split(all, "-")

	return StructuredUUID{
		all:      all,
		timehigh: generated_uuid[0],
		timemid:  generated_uuid[1],
		version:  generated_uuid[2],
		variant:  generated_uuid[3],
		random:   generated_uuid[4],
	}
}

func (u StructuredUUID) ToMap() map[string]string {
	return map[string]string{
		"timehigh": u.timehigh,
		"timemid":  u.timemid,
		"version":  u.version,
		"variant":  u.variant,
		"random":   u.random,
	}
}
