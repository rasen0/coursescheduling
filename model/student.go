package model

import "time"

type Student struct {
	ID string `json:"id" gorm:"not null;primary_key"`  // 类别号+年分后两位+月份+日号+四位递增数
	Name string `json:"name" gorm:"not null"`
	Age  uint `json:"age" gorm:"not null"`
	//RelateGroupID string `json:"relate_group_id"`
	Gender string `json:"gender" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null"`
	RelativeName string `json:"relative_name"`
	Relationship string `json:"relationship"`
	RelativePhone string `json:"relative_phone"`
	CoursePlan string `json:"course_plan" gorm:"not null"`
	StartDate time.Time `json:"start_date" gorm:"type:date"`
	EndDate time.Time `json:"end_date"  gorm:"type:date"`
	Desc string `json:"desc"`
	UpdateTime time.Time `gorm:"not null"`
}
