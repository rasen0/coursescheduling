package model

type StudentGroup struct {
	ID string `json:"id" gorm:"not null"` // 年分后两位+月份+日号+四位递增数
	GroupName string `json:"group_name" gorm:"not null,unique"`
}

type RelateGroup struct {
	ID string `json:"id" gorm:"not null"` // 年分后两位+月份+日号+四位递增数
	StudentGroupID string `json:"student_group_id" gorm:"not null"`
	TeacherID string `json:"teacher_id" gorm:"not null"`
}
