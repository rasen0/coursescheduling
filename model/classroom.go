package model

type Classroom struct {
	HouseNumber string `gorm:"not null"`
	Address string `gorm:"not null"`
}
