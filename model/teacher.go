package model

import "time"

type Teacher struct {
	ID int `json:"id" gorm:"autoIncrement"`
	SerialNumber string `json:"serial_number" gorm:"not null;unique"`   // 类别号+年份+月份+日号+四位递增数
	Name string `json:"name" gorm:"not null"`
	Age string `json:"age" gorm:"type:int;not null"`
	Gender string `json:"gender" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null"`
	Curriculum string `json:"curriculum"`
	RelativeName string `json:"relative_name"`
	Relationship string `json:"relationship"`
	RelativePhone string `json:"relative_phone"`
	Desc string `json:"desc"`
	UpdateTime time.Time `gorm:"not null"`
}
