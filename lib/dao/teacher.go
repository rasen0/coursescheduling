package dao

import (
	"coursesheduling/database"
	"coursesheduling/lib/entity"
	"coursesheduling/lib/util"
	"coursesheduling/model"
)

func InsertTeacherOne(teacher model.Teacher) {
	count := database.TeacherTotal()
	serialNumber,now := util.SplicingNumber(teacherNumber,count)
	teacher.SerialNumber = serialNumber
	teacher.UpdateTime = now
	database.InsertTeacherOne(teacher)
	return
}

func GetTeacherByPage(page, count int) (teachers []model.Teacher) {
	offSet := page * count
	teachers = database.GetTeacherByPage(offSet, count)
	return
}

func GetTeacherPagination(pagination entity.Pagination) (pageTotal int, total int64) {
	total = database.GetTableTotal(database.TeacherTable)
	if total == 0 {
		return
	}
	pageTotal = int(total/int64(pagination.PageSize))
	if total%int64(pagination.PageSize) > 0 {
		pageTotal++
	}
	return
}

