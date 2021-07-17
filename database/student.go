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

func GetStudentsByID(ID string) (students []model.Student) {
	appDB.Where("id = ?",ID).Find(&students)
	return
}

func GetStudentsByName(name string) (students []model.Student) {
	appDB.Where("name LIKE ?","%"+name+"%").Find(&students)
	return
}

func GetStudentsByStudentGroupID(groupID string) (students []model.Student) {
	appDB.Where("relate_group_id = ?",groupID).Find(&students)
	return
}

// GroupPagination 组分页查询
func GroupPagination(offSet, count int) (groups []model.StudentGroup){
	appDB.Limit(count).Offset(offSet).Find(&groups)
	return
}