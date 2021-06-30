package dao

import (
	"coursesheduling/database"
	"coursesheduling/lib/entity"
	"coursesheduling/model"
)

func InsertStudentOne(student model.Student) {
	database.InsertStudent(student)
	return
}

func GetStudentByPage(page, count int) (students []model.Student) {
	offSet := page * count
	students = database.GetStudentByPage(offSet, count)
	return
}

func GetStudentPagination(pagination entity.Pagination) (pageTotal int, total int64) {
	total = database.GetTableTotal(database.StudentTable)
	if total == 0 {
		return
	}
	pageTotal = int(total/int64(pagination.PageSize))
	if total%int64(pagination.PageSize) > 0 {
		pageTotal++
	}
	return
}
