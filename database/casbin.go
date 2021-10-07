package database

import (
	"coursesheduling/common"
	"coursesheduling/lib/config"
	"coursesheduling/lib/sqliteadapter"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"strings"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	CasbinModle = "config/rbac_model.conf"
	CasbinPolicy = "config/rbac_policy.csv"

	Root = "root"
	Admin = "admin"
	Guest = "guest"
	CommonUser = "common_user"
)

var rbacModelConf =`[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`

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
	enforcer.AddGroupingPolicy(Root,Admin)
	return enforcer,nil
}

func InitCasbin1() (enforcer *casbin.Enforcer,err error) {
	//absPath,_ := filepath.Abs(os.Args[0])

	log.Print("path:",os.Args[0])
	enforcer,  err = casbin.NewEnforcer(CasbinModle, CasbinPolicy)
	if err != nil{
		return nil, err
	}
	//enforcer.AddGroupingPolicy(Root,Admin)
	return enforcer,nil
}

func InitCasbin(db *gorm.DB) (enforcer *casbin.Enforcer,err error) {
	adapter := sqliteadapter.NewAdapter(db)
	log.Print("path:",os.Args[0])
	path := filepath.Join(common.Course,string(filepath.Separator),CasbinModle)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		createFile, _ := os.Create(path)
		_, err = createFile.WriteString(rbacModelConf)
		if err != nil{
			return
		}
	}
	enforcer,  err = casbin.NewEnforcer(path, adapter)
	if err != nil{
		return nil, err
	}

	return enforcer,nil
}

func InsertCasbinRules(crule []sqliteadapter.CasbinRule) {
	sqlBf := strings.Builder{}
	sqlBf.WriteString("insert into casbin_rule('p_type','v0','v1','v2','v3','v4','v5') values ")
	for i := range crule {
		sqlBf.WriteString(" ('"+crule[i].Ptype+"','"+crule[i].V0+"','"+crule[i].V1+"','"+
			crule[i].V2+"','"+crule[i].V3+"','"+crule[i].V4+"','"+crule[i].V5+"'),")
	}
	sqlStr := sqlBf.String()[:sqlBf.Len()-1]
	appDB.DB.Exec(sqlStr)
	return
}
func InsertCasbinRule(crule sqliteadapter.CasbinRule) {
	sqlBf := strings.Builder{}
	sqlBf.WriteString("insert into casbin_rule('p_type','v0','v1','v2','v3','v4','v5') values ")
	sqlBf.WriteString(" ('"+crule.Ptype+"','"+crule.V0+"','"+crule.V1+"','"+
		crule.V2+"','"+crule.V3+"','"+crule.V4+"','"+crule.V5+"'"+")")
	appDB.DB.Exec(sqlBf.String())
	return
}

func VerifyPolicy(role, entity, active string) (bool,error) {
	return appDB.Enforcer.Enforce(role, entity, active)
}

func QueryOne() {
	
}

func DeleteOne()  {
	
}
