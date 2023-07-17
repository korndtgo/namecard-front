package model

import "time"

type CreateCardDto struct {
	EmployeeId     string `json:"employeeId"`
	NameTh         string `json:"nameTh"`
	NameEn         string `json:"nameEn"`
	SurnameTh      string `json:"surnameTh"`
	SurnameEn      string `json:"surnameEn"`
	NicknameTh     string `json:"nicknameTh"`
	NicknameEn     string `json:"nicknameEn"`
	Email          string `json:"email"`
	ContactNumber1 string `json:"contactNumber1"`
	ContactNumber2 string `json:"contactNumber2"`
	LineId         string `json:"lineId"`
	PositionTh     string `json:"positionTh"`
	PositionEn     string `json:"positionEn"`
	DepartmentTh   string `json:"departmentTh"`
	DepartmentEn   string `json:"departmentEn"`
}

type ImportCardDto struct {
	File string `json:"file"`
}

type QueryCard struct {
	CompanyId  string `form:"companyId" json:"companyId"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	EmployeeId string `json:"employeeId"`
	Search     string `json:"search"`
	SortBy     string `json:"sortBy"`
	OrderBy    string `json:"orderBy"`
}

type QueryImportCard struct {
	Token string `json:"token"`
}

type UpdateCardDto struct {
	Id             string `json:"cardId" form:"cardId"`
	NameTh         string `json:"nameTh"`
	NameEn         string `json:"nameEn"`
	SurnameTh      string `json:"surnameTh"`
	SurnameEn      string `json:"surnameEn"`
	NicknameTh     string `json:"nicknameTh"`
	NicknameEn     string `json:"nicknameEn"`
	Email          string `json:"email"`
	ContactNumber1 string `json:"contactNumber1"`
	ContactNumber2 string `json:"contactNumber2"`
	LineId         string `json:"lineId"`
	PositionTh     string `json:"positionTh"`
	PositionEn     string `json:"positionEn"`
	DepartmentTh   string `json:"departmentTh"`
	DepartmentEn   string `json:"departmentEn"`
}

type CardRequest struct {
	Id string `json:"cardId" form:"cardId"`
}

type CardDto struct {
	Id int `json:"id"`

	Uuid string `json:"uuid"`

	EmployeeId string `json:"employeeId"`

	NameTh string `json:"nameTh"`

	NameEn string `json:"nameEn"`

	SurnameTh string `json:"surnameTh"`

	SurnameEn string `json:"surnameEn"`

	NicknameTh string `json:"nicknameTh"`

	NicknameEn string `json:"nicknameEn"`

	Email string `json:"email"`

	ContactNumber1 string `json:"contactNumber1"`

	ContactNumber2 string `json:"contactNumber2"`

	LineId string `json:"lineId"`

	PositionTh string `json:"positionTh"`

	PositionEn string `json:"positionEn"`

	DepartmentTh string `json:"departmentTh"`

	DepartmentEn string `json:"departmentEn"`

	CompanyId int `json:"companyId"`

	Company CompanyDto `json:"company"`

	CreatedAt time.Time `json:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt"`

	DeletedAt time.Time `json:"deletedAt"`

	DownloadUrlTh string `json:"download_url_th"`

	DownloadUrlEn string `json:"download_url_en"`
}

type CompanyDto struct {
	Logo   string `json:"logo"`
	NameTh string `json:"name_th"`
	NameEn string `json:"name_en"`
}
