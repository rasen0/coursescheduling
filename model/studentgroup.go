package model

type StudentGroup struct {
	ID int `json:"id" gorm:"type:integer  primary key autoincrement"`
	GroupName string `json:"GroupName" gorm:"not null,unique"`
	Desc string `json:"desc"`
}

type RelateGroup struct {
	ID string `json:"id" gorm:"autoIncrement;not null"`
	StudentGroupID string `json:"student_group_id" gorm:"not null"`
	TeacherID string `json:"teacher_id" gorm:"not null"`
}
