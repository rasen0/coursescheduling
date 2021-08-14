package dao

import (
	"coursesheduling/database"
	"coursesheduling/lib/sqliteadapter"
)

func AddCasbinPolicy(crules []sqliteadapter.CasbinRule)  {
	database.InsertCasbinRules(crules)
}
