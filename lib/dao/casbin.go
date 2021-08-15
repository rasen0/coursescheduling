package dao

import (
	"coursesheduling/database"
	"coursesheduling/lib/sqliteadapter"
)

func AddCasbinPolicy(crules []sqliteadapter.CasbinRule)  {
	database.InsertCasbinRules(crules)
}

func VerifyPolicy(role, entity, active string) (bool,error) {
	return database.VerifyPolicy(role, entity, active)
}