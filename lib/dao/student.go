package dao

import (
	"coursesheduling/database"
	"coursesheduling/lib/entity"
	"coursesheduling/lib/util"
	"coursesheduling/model"
)

func InsertStudentOne(student model.Student) {
	count := database.StudentTotal()
	serialNumber,now := util.SplicingNumber(studentNumber,count)
	student.SerialNumber = serialNumber
	student.UpdateTime = now
	database.InsertStudent(student)
	return
}

func GetStudentByPage(page, count int) (students []model.Student) {
	if page < 0 {
		page = 0
	}
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

// GetStudentsByID 使用id查询学生
func GetStudentsByID(ID string) (students []model.Student) {
	students = database.GetStudentsByID(ID)
	return
}

// GetStudentsByName 使用name查询学生
func GetStudentsByName(name string) (students []model.Student) {
	students = database.GetStudentsByName(name)
	return
}

// GetStudentsByStudentGroupID 使用组号查询学生
func GetStudentsByStudentGroupID(groupID string) (students []model.Student) {
	students = database.GetStudentsByStudentGroupID(groupID)
	return
}

func GroupPagination(page, count int) (groups []model.StudentGroup){
	offSet := page * count
	groups = database.GroupPagination(offSet, count)
	return
}
