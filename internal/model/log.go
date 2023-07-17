package model

import "time"

type LogCardDto struct {
	Id         int          `json:"id"`
	Event      string       `json:"event"`
	Causer     LogCauserDto `json:"causer"`
	Properties string       `json:"properties"`
	CreateAt   time.Time    `json:"createAt"`
}

type LogCauserDto struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type QueryLogs struct {
	SubjectType string `json:"subjectType"`

	CompanyId int `json:"companyId"`

	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
