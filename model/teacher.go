package model

import "time"

type Teacher struct {
	ID string `json:"id" gorm:"not null;primary_key"`  // 类别号+年分后两位+月份+日号+四位递增数
	Name string `json:"name" gorm:"not null"`
	Age int `json:"age" gorm:"not null"`
	Gender string `json:"gender" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null"`
	RelativeName string `json:"relative_name"`
	Relationship string `json:"relationship"`
	RelativePhone string `json:"relative_phone"`
	Desc string `json:"desc"`
	UpdateTime time.Time `gorm:"not null"`
}
