package model

type Curriculum struct {
	Name string
}

type Course struct {
	TeacherNumber int
	CurriculumNumber int
	ClassroomNumber int
	StudentGroup string // 学生编号集合
	ScheduledTime string
	UpdateTime string
}
