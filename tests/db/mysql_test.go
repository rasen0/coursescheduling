package db_test

import (
	"coursesheduling/lib/util"
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
	month1, month2 := util.DurationMonth(time.Now())
	var courses []model.CommonCourse
	db.Where("course_date > ? and course_date < ?",month1.Format("2006-01-02"),month2.Format("2006-01-02")).
		Find(&courses)
	for i,r := range courses{
		fmt.Println("i:",i," r:",r)
	}
}
