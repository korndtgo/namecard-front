package repository

import (
	"card-service/internal/config"
	"fmt"
	"math"
	"time"
)

type TimeRepository struct {
	Locale *time.Location
}
type ITimeRepository interface {
	CurrentTime() time.Time
	CurrentDateAtTime() time.Time
	CurrentDate() string
	CurrentTomorrowDate() string
	GetTimeOffset(ct time.Time) string
}

func (t *TimeRepository) CurrentTime() time.Time {
	return time.Now().In(t.Locale)
}

func (t *TimeRepository) CurrentDateAtTime() time.Time {
	ctime, _ := time.Parse("20060102", time.Now().In(t.Locale).Format("20060102"))
	return ctime
}

func (t *TimeRepository) CurrentDate() string {
	return time.Now().In(t.Locale).Format("20060102")
}

func (t *TimeRepository) CurrentTomorrowDate() string {
	return time.Now().In(t.Locale).AddDate(0, 0, 1).Format("20060102")
}

func (t *TimeRepository) CurrentDateTime() string {
	return time.Now().In(t.Locale).Format("20060102150405")
}

func (t *TimeRepository) CurrentWithFormat(format string) string {
	return time.Now().In(t.Locale).Format(format)
}

func (t *TimeRepository) GetTimeOffset(ct time.Time) string {
	_, currentOffset := ct.Zone()
	h, _ := time.ParseDuration(fmt.Sprintf("%vs", currentOffset))
	timeOffset := fmt.Sprintf("%02d:00", int8(math.Abs(h.Hours())))
	if h.Hours() > 0 {
		timeOffset = "+" + timeOffset
	} else {
		timeOffset = "-" + timeOffset
	}
	return timeOffset
}

func NewTimeRepository(cfg *config.Configuration) ITimeRepository {
	loc, err := time.LoadLocation(cfg.Locale)
	if err != nil {
		panic("Cannot Load Locale" + cfg.Locale)
	}
	return &TimeRepository{
		Locale: loc,
	}
}
