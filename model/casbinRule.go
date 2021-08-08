package model

type CasbinRule struct {
	Id int `json:"id" gorm:"autoIncrement not null"`
	PType string `json:"p_type" gorm:"type:varchar(100); not null"`
	V0 string `json:"v0" gorm:"type:varchar(100); not null"`
	V1 string `json:"v1" gorm:"type:varchar(100); not null"`
	V2 string `json:"v2" gorm:"type:varchar(100); not null"`
	V3 string `json:"v3" gorm:"type:varchar(100); not null"`
	V4 string `json:"v4" gorm:"type:varchar(100); not null"`
	V5 string `json:"v5" gorm:"type:varchar(100); not null"`
}
