package database

import "coursesheduling/model"

func InsertRoomOne(room model.Classroom) {
	appDB.Create(&room)
}

func GetRooms() (rooms []model.Classroom) {
	appDB.Find(&rooms)
	return
}

func GetPlanByName(queryWord string) (coursePlans []model.CoursePlan) {
	appDB.Where("name LIKE ?","%"+queryWord+"%").Find(&coursePlans)
	return coursePlans
}

func GetCurriculumByName(queryWord string) (curriculums []model.Curriculum) {
	appDB.Where("name LIKE ?","%"+queryWord+"%").Find(&curriculums)
	return curriculums
}

func GetCurriculums() (curriculums []model.Curriculum) {
	appDB.Find(&curriculums)
	return
}

func GetCoursePlans() (coursePlans []model.CoursePlan) {
	appDB.Find(&coursePlans)
	return
}

func GetRoomByName(queryWord string) (classrooms []model.Classroom) {
	appDB.Where("room_name LIKE ?","%"+queryWord+"%").Find(&classrooms)
	return classrooms
}

func GetRoomByPage(offset, count int) (rooms []model.Classroom) {
	appDB.Limit(count).Offset(offset).Find(&rooms)
	return
}
