package database

import "coursesheduling/model"

func GetPlanByName(queryWord string) (coursePlans []model.CoursePlan) {
	appDB.Where("name LIKE ?","%"+queryWord+"%").Find(&coursePlans)
	return coursePlans
}

func GetCurriculumByName(queryWord string) (curriculums []model.Curriculum) {
	appDB.Where("name LIKE ?","%"+queryWord+"%").Find(&curriculums)
	return curriculums
}

func GetRoomByName(queryWord string) (classrooms []model.Classroom) {
	appDB.Where("room_name LIKE ?","%"+queryWord+"%").Find(&classrooms)
	return classrooms
}

func GetRoomByPage(offset, count int) (rooms []model.Classroom) {
	appDB.Limit(count).Offset(offset).Find(&rooms)
	return
}
