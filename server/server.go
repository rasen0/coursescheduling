package server

import (
	"coursesheduling/lib/config"
	"coursesheduling/lib/dao"
	"coursesheduling/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ServeWrapper struct {
	*http.Server
	*config.Configure
}

func NewServer(config *config.Configure) *ServeWrapper {
	return &ServeWrapper{
		Configure:config,
	}
}

func (svr *ServeWrapper) Serve()  {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(Cors())
	svr.Server = &http.Server{
		IdleTimeout: 10*time.Second,
		ReadTimeout: 15*time.Second,
		WriteTimeout: 15*time.Second,
		Addr: svr.Address,
		Handler: engine,
	}
	svr.SetKeepAlivesEnabled(true)
	serveGroup := engine.Group("/service")
	serveGroup.GET("/v1/students",svr.GetCourseScheduling)
	serveGroup.GET("/v1/getstudents",svr.GetStudentsByGroupID)
	serveGroup.GET("/v1/grouppagination",svr.GroupPagination)
	serveGroup.GET("/v1/curriculumOptions",svr.GetCurriculumOptions)
	serveGroup.GET("/v1/coursePlanOptions",svr.GetCoursePlanOptions)
	serveGroup.GET("/v1/coursescheduling",svr.GetCourseScheduling)

	serveGroup.GET("/v1/queryteacherbykey",svr.QueryTeacherByKey)
	serveGroup.GET("/v1/querygroupbykey",svr.QueryGroupByKey)
	serveGroup.GET("/v1/queryplanbykey",svr.QueryPlanByKey)
	serveGroup.GET("/v1/querycurriculumebykey",svr.QueryCurriculumByKey)
	serveGroup.GET("/v1/queryroombykey",svr.QueryRoomByKey)

	serveGroup.POST("/v1/addingstudent",svr.AddStudent)
	serveGroup.POST("/v1/addingteacher",svr.AddTeacher)
	serveGroup.POST("/v1/addcommoncourses",svr.AddCommonCourses)
	serveGroup.POST("/v1/addtrialcourses",svr.AddTrialCourses)
	serveGroup.POST("/v1/addsinglecourses",svr.AddSingleCourses)

	fmt.Println("start course scheduling system")
	go svr.ListenAndServe()
}

/*
GetCourseScheduling 默认请求当月课程安排
month 请求月份
ctype 请求安排课程类型
*/
func (svr *ServeWrapper) GetCourseScheduling(ctx *gin.Context)  {
	monthStr := ctx.GetString("month")
	var month time.Time
	if monthStr <= ""{
		month = time.Now()
	}else {
		month,_ = time.Parse(time.RFC3339,monthStr)
	}
	fmt.Println(month)
	ctype := ctx.GetInt("ctype")
	if ctype <= 0 {
		ctype = model.CommonLesson
	}
	result := make(map[string]interface{})
	var err error
	defer func() {
		if err == nil{
			return
		}
		fmt.Println("fail")
		result["status"]="fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}()

	_, courseTable := dao.GetCourseTable(ctype, month)
	result["courseTable"] =courseTable
	ctx.JSON(http.StatusOK,result)
	return
}
