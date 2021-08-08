package model

type RoleItem struct {
	Id int `json:"id" gorm:"type:tinyint not null auto_increment"`
	Role string `json:"role" gorm:"type:varchar(30) not null"`
}

type Account struct {
	UserName string `json:"user_name" gorm:"type:varchar(30) not null"`
	Password string `json:"password" grom:"type:varchar(30) not null"`
	Role string `json:"role" gorm:"type:varchar(30) not null"`
	Update string `json:"update" gorm:"type:date not null"`
}

