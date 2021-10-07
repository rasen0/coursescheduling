package dao

import (
	"coursesheduling/database"
	"coursesheduling/model"
)


func GetGroup() (groups []model.StudentGroup){
	groups = database.GetGroup()
	return
}

func InsertGroupOne(group model.StudentGroup) {
	database.InsertGroupOne(group)
	return
}

func GetGroups(offset, count int) (groups []model.StudentGroup) {
	groups = database.GetGroupByPage(offset, count)
	return
}
