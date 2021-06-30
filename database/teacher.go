package database

import "coursesheduling/model"

func GetTeacherByName(queryWord string) (teachers []model.Teacher) {
	appDB.Where("name LIKE ?","%"+queryWord+"%").Find(&teachers)
	return teachers
}

func GetTeacherByPage(offset, count int) (teachers []model.Teacher) {
	appDB.Limit(count).Offset(offset).Find(&teachers)
	return
}

func InsertTeacherOne(teacher model.Teacher) {
	appDB.Create(teacher)
	return
}
