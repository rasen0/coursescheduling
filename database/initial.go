package database

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"coursesheduling/common"
	"coursesheduling/lib/config"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var appDB *CourseDB

type CourseDB struct {
	*gorm.DB
	*casbin.Enforcer
}

func(c *CourseDB) ConfigCasbin()  {
	enforcer, err := InitCasbin(c.DB)
	if err != nil{
		log.Error("init casbin error.",err)
		return
	}
	appDB.Enforcer = enforcer
}

func ConnectDB(dbInfo config.DBInfo) (courseDB *CourseDB){
	var err error
	var db *gorm.DB
	if !strings.HasSuffix(dbInfo.DBName,".db"){
		err = errors.New("empty")
		//dsn := dbInfo.DBUser+":"+dbInfo.DBPassword+"@tcp("+dbInfo.IpAddress+")/"+dbInfo.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"
		//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		//db, err = mysqlLoader(dsn)
	} else {
		db,err = sqliteLoader(common.ConfigPath+"/"+dbInfo.DBName)
	}
	if err != nil {
		log.Error("open database fail.",err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil{
		fmt.Println("db err:",err)
	}
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	courseDB = &CourseDB{
		DB:db,
	}
	appDB = courseDB
	return
}

func(courseDB *CourseDB) InitialData() (errWrapper error){
	//rows,err := courseDB.Raw("show tables").Rows()
	//if err != nil{
	//	errWrapper = fmt.Errorf("rows db err:%w",err)
	//	return
	//}
	//for idx:= 0; rows.Next();idx++ {
	//	var t string
	//	rows.Scan(&t)
	//	ts = append(ts,t)
	//}

	sqlms := make([]sqliteMaster,0)
	courseDB.Find(&sqlms)
	ts := make([]string,0)
	for idx:= 0; idx < len(sqlms);idx++ {
		ts = append(ts,sqlms[idx].TblName)
	}
	isAccount := false
	isRoleItem := false
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
			if table == RoleTable{
				isRoleItem = true
			}
			if table == AccountTable{
				isAccount = true
			}
			err := courseDB.AutoMigrate(mod)
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

	courseDB.ConfigCasbin()
	if isRoleItem {
		log.Print("insert role_item data")
		roles := [2]model.RoleItem{
			{Role:"admin"},
			{Role:"common_user"},
		}
		// 插入数据库
		courseDB.CreateInBatches(&roles,2)
		courseDB.Enforcer.AddGroupingPolicy(Root,Admin)
		//courseDB.Enforcer.AddPolicies()
		//InsertCasbinRule(sqliteadapter.CasbinRule{Ptype: "g",V0:"root",V1:"admin"})
	}
	if isAccount {
		accounts := [2]model.Account{
			{"root","root",common.AdminRole,"","2021-08-15 15:04:05"},
			{"guest","guest",common.CommonRole,"","2021-08-15 15:04:05"},
		}
		InsertAccounts(accounts[:])
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