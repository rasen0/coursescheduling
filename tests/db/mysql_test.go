package db_test

import (
	"coursesheduling/model"
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMysql(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/courseschedule?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil{
		fmt.Println("db err:",err)
	}
	// See "Important settings" section.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	err = db.AutoMigrate(&model.Curriculum{})
	if err != nil{
		fmt.Println("auto err:",err)
	}
	curriculums := [9]model.Curriculum{
		{
			1,
			"语文",
		},{
			2,
			"数学",
		},{
			3,
			"物理",
		},{
			4,
			"历史",
		},{
			5,
			"英语",
		},{
			6,
			"化学",
		},{
			7,
			"地理",
		},{
			8,
			"政治",
		},{
			9,
			"生物",
		},
	}
	db.CreateInBatches(&curriculums,9)

}
