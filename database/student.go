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

func StudentTotal() (count int64) {
	appDB.Table(StudentTable).Count(&count)
	return
}

func GetStudentsByGroupID(groupID string) (students []model.Student) {
	appDB.Where("id = ?",groupID).Find(&students)
	return
}