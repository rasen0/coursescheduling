package server

import (
	"coursesheduling/lib/dao"
	"coursesheduling/lib/entity"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AddTeacher 添加一个老师信息
func (svr *ServeWrapper) AddTeacher(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var newTeacher model.Teacher
	ctx.BindJSON(&newTeacher)
	newTeacher.UpdateTime = time.Now()
	dao.InsertTeacherOne(newTeacher)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) GetTeachers(ctx *gin.Context)  {
	page := ctx.GetInt("page")
	count := ctx.GetInt("count")
	result := make(map[string]interface{})
	teachers := dao.GetTeacherByPage(page, count)
	result["teachers"] =teachers
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) TeacherPagination(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var pagination entity.Pagination
	ctx.BindJSON(&pagination)
	_, total := dao.GetTeacherPagination(pagination)
	pagination.Total = total
	result["pagination"] = pagination
	teachers := dao.GetStudentByPage(0, pagination.PageSize)
	result["teachers"] = teachers
	ctx.JSON(http.StatusOK,result)
	return
}
