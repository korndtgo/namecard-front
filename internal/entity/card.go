package entity

import "time"

type Card struct {
	Id int `gorm:"primarykey;column:id;"`

	Uuid string `gorm:"column:uuid;"`

	EmployeeId string `gorm:"column:employee_id;"`

	NameTh string `gorm:"column:name_th;"`

	NameEn string `gorm:"column:name_en;"`

	SurnameTh string `gorm:"column:surname_th;"`

	SurnameEn string `gorm:"column:surname_en;"`

	NicknameTh string `gorm:"column:nickname_th;"`

	NicknameEn string `gorm:"column:nickname_en;"`

	Email string `gorm:"column:email;"`

	ContactNumber1 string `gorm:"column:contact_number_1;"`

	ContactNumber2 string `gorm:"column:contact_number_2;"`

	LineId string `gorm:"column:line_id;"`

	PositionTh string `gorm:"column:position_th;"`

	PositionEn string `gorm:"column:position_en;"`

	DepartmentTh string `gorm:"column:department_th;"`

	DepartmentEn string `gorm:"column:department_en;"`

	CompanyId int `gorm:"column:company_id;"`

	IsActive bool `gorm:"column:is_active;"`

	CreatedAt time.Time `gorm:"column:created_at;"`

	UpdatedAt time.Time `gorm:"column:updated_at;"`

	DeletedAt time.Time `gorm:"column:deleted_at;"`
}
