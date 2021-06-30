package server

import (
	"coursesheduling/lib/dao"
	"coursesheduling/lib/entity"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
GetStudents 默认请求学生数据
page 参数当前页
count 参数每页展示条数
*/
func (svr *ServeWrapper) GetStudents(ctx *gin.Context)  {
	page := ctx.GetInt("page")
	count := ctx.GetInt("count")
	result := make(map[string]interface{})
	students := dao.GetStudentByPage(page, count)
	result["students"] =students
	ctx.JSON(http.StatusOK,result)
	return
}

// AddStudent 添加一个学生信息
func (svr *ServeWrapper) AddStudent(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var newStudent model.Student
	ctx.BindJSON(&newStudent)
	newStudent.UpdateTime = time.Now()
	dao.InsertStudentOne(newStudent)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) StudentPagination(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var pagination entity.Pagination
	ctx.BindJSON(&pagination)
	_, total := dao.GetStudentPagination(pagination)
	pagination.Total = total
	result["pagination"] = pagination
	students := dao.GetStudentByPage(0, pagination.PageSize)
	result["students"] = students
	ctx.JSON(http.StatusOK,result)
	return
}