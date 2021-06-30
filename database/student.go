package database

import "coursesheduling/model"

func GetStudentByPage(offset, count int) (students []model.Student) {
	appDB.Limit(count).Offset(offset).Find(&students)
	return students
}

func InsertStudent(student model.Student) {
	appDB.Create(&student)
	return
}

func GetGroupByName(queryWord string) (groups []model.StudentGroup) {
	appDB.Where("group_name LIKE ?","%"+queryWord+"%").Find(&groups)
	return groups
}
