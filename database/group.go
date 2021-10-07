package database

import "coursesheduling/model"

func InsertGroupOne(group model.StudentGroup) {
	appDB.Create(&group)
}

func GetGroup() (group []model.StudentGroup) {
	appDB.Find(&group)
	return
}

func GetGroupByName(queryWord string) (groups []model.StudentGroup) {
	appDB.Where("group_name LIKE ?","%"+queryWord+"%").Find(&groups)
	return groups
}


func GetGroupByPage(offset, count int) (groups []model.StudentGroup) {
	appDB.Limit(count).Offset(offset).Find(&groups)
	return
}
