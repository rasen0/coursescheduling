package dao

import (
	"coursesheduling/database"
	"coursesheduling/model"
)

func QueryTeacherByWord(queryWord string) (teachers []model.Teacher) {
	teachers = database.GetTeacherByName(queryWord)
	return teachers
}

func QueryGroupByWord(queryWord string) (groups []model.StudentGroup) {
	groups = database.GetGroupByName(queryWord)
	return groups
}

func QueryPlanByWord(queryWord string) (CoursePlans []model.CoursePlan) {
	CoursePlans = database.GetPlanByName(queryWord)
	return CoursePlans
}

func QueryCurriculumByWord(queryWord string) (Curriculums []model.Curriculum) {
	Curriculums = database.GetCurriculumByName(queryWord)
	return Curriculums
}

func QueryRoomByWord(queryWord string) (classRooms []model.Classroom) {
	classRooms = database.GetRoomByName(queryWord)
	return classRooms
}
