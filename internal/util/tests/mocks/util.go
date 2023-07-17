package util_mock

import (
	"campaign-service/internal/util"
	"log"
	"time"
)

type MockUtils struct {
	GenerateUUID []string
	TimeNow      []string
}

func NewUtils(u *util.Utils, m MockUtils) *util.Utils {

	log.Println("time:::", u.TimeNow().String())
	if m.GenerateUUID != nil {
		u.GenerateUUID = func() string {
			mock := m.GenerateUUID[0]
			if len(m.GenerateUUID) > 1 {
				m.GenerateUUID = m.GenerateUUID[1:]
			}
			return mock
		}
	}

	if m.TimeNow != nil {
		u.TimeNow = func() time.Time {
			mock := m.TimeNow[0]
			if len(m.TimeNow) > 1 {
				m.TimeNow = m.TimeNow[1:]
			}
			t, _ := time.Parse("2006-01-02T15:04:05", mock)
			return t
		}
	}
	return u
}
