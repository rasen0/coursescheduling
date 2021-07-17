package dao

import (
	"coursesheduling/database"
	"coursesheduling/model"
)


func GetRooms() (rooms []model.Classroom){
	rooms = database.GetRooms()
	return
}

func InsertRoomOne(room model.Classroom) {
	//room.UpdateTime = now
	database.InsertRoomOne(room)
	return
}

func GetCurriculums() (curriculums []model.Curriculum) {
	curriculums = database.GetCurriculums()
	return
}

func GetCoursePlans() (coursePlans []model.CoursePlan) {
	coursePlans = database.GetCoursePlans()
	return
}
