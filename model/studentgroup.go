package model

type StudentGroup struct {
	ID int64 `gorm:"not null"`
	StudentUUID string `gorm:"not null"`
}
