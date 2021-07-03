package model

import "time"

type Student struct {
	ID int `json:"id" gorm:"autoIncrement"`
	SerialNumber string `json:"serial_number" gorm:"not null;unique"`   // 类别号+年分+月份+日号+四位递增数
	Name string `json:"name" gorm:"not null"`
	Age  string `json:"age" gorm:"type:int;not null"`
	RelateGroupID int `json:"relate_group_id" gorm:"type:int(11)"`
	Gender string `json:"gender" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null"`
	RelativeName string `json:"relative_name"`
	Relationship string `json:"relationship"`
	RelativePhone string `json:"relative_phone"`
	CoursePlan int `json:"course_plan" gorm:"type:int(11);not null"`
	StartDate time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate time.Time `json:"end_date"  gorm:"type:date"`
	Desc string `json:"desc"`
	UpdateTime time.Time `gorm:"not null"`
}
