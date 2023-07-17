package util

import (
	"card-service/internal/config"
	"time"
)

type Utils struct {
	GenerateUUID func() string
	TimeNow      func() time.Time
}

type IUtils interface{}

func NewUtils(cfg *config.Configuration) *Utils {
	loc, err := time.LoadLocation(cfg.Locale)
	if err != nil {
		panic("Cannot Load Locale" + cfg.Locale)
	}
	return &Utils{
		GenerateUUID: GenerateUUID,
		TimeNow: func() time.Time {
			return time.Now().In(loc)
		},
	}
}
