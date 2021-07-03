package dao

import (
	"coursesheduling/database"
	"coursesheduling/model"
)

func GetCurriculums() (curriculums []model.Curriculum) {
	curriculums = database.GetCurriculums()
	return
}

func GetCoursePlans() (coursePlans []model.CoursePlan) {
	coursePlans = database.GetCoursePlans()
	return
}
