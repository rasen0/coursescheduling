package model

type RoleItem struct {
	Id int `json:"id" gorm:"type:INTEGER primary key autoincrement"`
	Role string `json:"role" gorm:"type:varchar(30) not null"`
}

type Account struct {
	UserName string `json:"userName" gorm:"type:varchar(30) not null"`
	TeacherID string `json:"teacherId" gorm:"type:varchar(100) not null"`
	Password string `json:"password" grom:"type:varchar(30) not null"`
	Role string `json:"role" gorm:"type:varchar(30) not null"`
	Token string `json:"token" gorm:"type:varchar(200)"`
	Update string `json:"update" gorm:"type:date not null"`
}
