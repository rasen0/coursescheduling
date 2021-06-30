package model

type Classroom struct {
	Id int `json:"id" gorm:"not null"`
	RoomName string `json:"room_name" gorm:"not null,unique"`
	Address string `json:"address" gorm:"not null"`
	Used int `json:"used"`
}

type Curriculum struct {
	ID uint `json:"id" gorm:"not null"`
	Name string `json:"name" gorm:"not null,unique"`
}

type CoursePlan struct {
	ID uint `json:"id" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
	Explain string `json:"explain" gorm:"not null"`
}
