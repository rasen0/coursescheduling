package model

type Classroom struct {
	RoomName string `json:"room_name" gorm:"not null,unique"`
	Address string `json:"address" gorm:"not null"`
	Used int `json:"used"`
}

type Curriculum struct {
	ID uint `json:"id" gorm:"type:int;not null"`
	Name string `json:"name" gorm:"not null,unique"`
}

type CoursePlan struct {
	ID uint `json:"id" gorm:"autoIncrement;not null"`
	Name string `json:"name" gorm:"not null"`
	Explain string `json:"explain"`
}
