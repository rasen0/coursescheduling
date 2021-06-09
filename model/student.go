package model

import "time"

type Student struct {
	UUID string `gorm:"not null"`
	Name string `gorm:"not null"`
	Gender uint `gorm:"not null"`
	Age  uint `gorm:"not null"`
	Phone string `gorm:"not null"`
	CoursePlan string `gorm:"not null"`
	RelativeName string
	Relationship string
	RelativePhone string
	StartDate time.Time
	EndDate time.Time
	Desc string
	UpdateTime string `gorm:"not null"`
}
