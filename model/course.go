package model

type Curriculum struct {
	ID uint `gorm:"not null"`
	Name string `gorm:"not null"`
}

type Course struct {
	TeacherUUID string `gorm:"not null"`
	CurriculumNumber uint `gorm:"not null"`
	ClassroomNumber int `gorm:"not null"`
	StudentGroup int  `gorm:"not null"`// 学生编号集合
	ScheduledTime string
	UpdateTime string `gorm:"not null"`
}
