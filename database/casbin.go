package database

import (
	"coursesheduling/common"
	"coursesheduling/lib/config"
	"coursesheduling/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	CasbinModle = "config/rbac_model.conf"
	CasbinPolicy = "config/rbac_policy.conf"

	Root = "root"
	Admin = "admin"
)

func InitCasbin0(dbInfo config.DBInfo) (enforcer *casbin.Enforcer,err error) {
	var adapter *gormadapter.Adapter
	//adapter, err = gormadapter.NewAdapter("mysql",dbInfo.DBUser+":"+dbInfo.DBPassword+"@tcp("+dbInfo.IpAddress+")/"+dbInfo.DBName,true)
	adapter, err = gormadapter.NewAdapter("sqlite3",common.ConfigPath+"/"+dbInfo.DBName,false)
	if err != nil{
		return nil, err
	}
	//e, err := casbin.NewEnforcer("src/casbin/conf/rbac_model.conf", "src/casbin/conf/rbac_policy.csv")
	enforcer,  err = casbin.NewEnforcer(CasbinModle, adapter)
	if err != nil{
		return nil, err
	}
	//enforcer.AddGroupingPolicy(Root,Admin)
	return enforcer,nil
}

func InitCasbin() (enforcer *casbin.Enforcer,err error) {
	enforcer,  err = casbin.NewEnforcer(CasbinModle, CasbinPolicy)
	if err != nil{
		return nil, err
	}
	//enforcer.AddGroupingPolicy(Root,Admin)
	return enforcer,nil
}

func Insert(crule []model.CasbinRule) {
	sqlBf := strings.Builder{}
	sqlBf.WriteString("insert into casbin_rule('p_type','v0','v1','v2','v3','v4','v5') values ")
	for i := range crule {
		sqlBf.WriteString(" ("+ "'"+crule[i].PType+"',"+"'"+crule[i].V0+"',"+"'"+crule[i].V1+"',"+
			"',"+"'"+crule[i].V2+"',"+"',"+"'"+crule[i].V3+"',"+"',"+"'"+crule[i].V4+"',"+"',"+"'"+crule[i].V5+"'"+"),")
	}
	sqlStr := sqlBf.String()[:sqlBf.Len()-1]
	appDB.DB.Exec(sqlStr)
	return
}

func QueryOne() {
	
}

func DeleteOne()  {
	
}
