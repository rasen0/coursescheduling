package server

import (
	"coursesheduling/lib/dao"
	"coursesheduling/lib/entity"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// AddTeacher 添加一个老师信息
func (svr *ServeWrapper) AddTeacher(ctx *gin.Context)  {
	result := make(map[string]interface{})
	//var newTeacher model.Teacher
	var reqTeacher = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		Teacher model.Teacher
	}{}
	ctx.BindJSON(&reqTeacher)
	reqTeacher.Teacher.UpdateTime = time.Now()
	log.Printf("new teacher:%+v",reqTeacher.Teacher)
	dao.InsertTeacherOne(reqTeacher.Teacher)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	log.Print("add teacher done")
	return
}

func (svr *ServeWrapper) GetTeachers(ctx *gin.Context) {
	pageParam := ctx.Query("page")
	countParam := ctx.Query("count")
	result := make(map[string]interface{})
	page,_ := strconv.Atoi(pageParam)
	count,_ := strconv.Atoi(countParam)
	teachers := dao.GetTeacherByPage(page-1, count)
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
