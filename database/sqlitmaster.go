package database

type sqliteMaster struct {
	Ttype string `json:"type" gorm:"column:type;"`
	Name string `json:"name" gorm:"name"`
	TblName string `json:"tbl_name" gorm:"tbl_name"`
	RootPage int `json:"rootpage" gorm:"rootpage"`
	Sql string `json:"sql" gorm:"sql"`
}

func (sqliteMaster) TableName() string {
	return "sqlite_master"
}
