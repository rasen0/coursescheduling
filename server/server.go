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
	serveGroup.GET("/v1/coursescheduling",svr.GetCourseScheduling)

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
	courseMonth := dao.GetCourseMonth(ctype, month)
	result["courses"] =courseMonth
	//switch ctype {
	//case model.SingleLesson:
	//	courses := database.GetSingleCourseByMonth(month)
	//	result["singleCourse"] = courses
	//case model.TrialLesson:
	//	courses := database.GetTrialCourseByMonth(month)
	//	result["trialCourse"] = courses
	//case model.CommonLesson:
	//	courses := database.GetCommonCourseByMonth(month)
	//	result["commonCourse"] = courses
	//default:
	//	log.Error("request courses fail")
	//}
	ctx.JSON(http.StatusOK,result)
	return
}