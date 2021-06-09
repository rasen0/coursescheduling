package model

type Teacher struct {
	UUID string `gorm:"not null"`
	Name string `gorm:"not null"`
	Age int `gorm:"not null"`
	Gender int `gorm:"not null"`
	Phone string `gorm:"not null"`
	RelativeName string
	Relationship string
	RelativePhone string
	UpdateTime string `gorm:"not null"`
}
