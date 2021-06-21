package database

import (
	"errors"
	"fmt"
	"time"

	"coursesheduling/lib/config"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var appDB *CourseDB

type CourseDB struct {
	*gorm.DB
}

func ConnectDB(dbInfo config.DBInfo) (courseDB *CourseDB){
	dsn := dbInfo.DBUser+":"+dbInfo.DBPassword+"@tcp("+dbInfo.IpAddress+")/"+dbInfo.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("open database fail.",err)
		return
	}
	courseDB = &CourseDB{
		db,
	}
	sqlDB, err := db.DB()
	if err != nil{
		fmt.Println("db err:",err)
	}
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	appDB = courseDB
	return
}

func(courseDB *CourseDB) InitialData() (errWrapper error){
	rows,err := courseDB.Raw("show tables").Rows()
	if err != nil{
		errWrapper = fmt.Errorf("rows db err:%w",err)
		return
	}
	ts := make([]string,0)
	for idx:= 0; rows.Next();idx++ {
		var t string
		rows.Scan(&t)
		ts = append(ts,t)
	}
	isCurriculumTable := false

	for table,mod := range dBTable{
		exist := false
		for _, t := range ts{
			if t == table {
				exist = true
				break
			}
		}
		if !exist{
			if table == CurriculumTable{
				isCurriculumTable = true
			}
			err = courseDB.AutoMigrate(mod)
			if err != nil{
				errWrapper = fmt.Errorf("auto err:%w",err)
			}
		}
	}
	if isCurriculumTable{
		log.Print("insert curriculum data")
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
		courseDB.CreateInBatches(&curriculums,9)
	}

	return
}

func InitDB(dbInfo config.DBInfo) (err error) {
	db := ConnectDB(dbInfo)
	if db == nil{
		err = errors.New("open database fail!")
		return
	}
	err = db.InitialData()
	return err
}