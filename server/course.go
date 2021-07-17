package server

import (
	"coursesheduling/lib/dao"
	"fmt"
	"net/http"
	"time"

	"coursesheduling/model"
	"github.com/gin-gonic/gin"
)

func (svr *ServeWrapper) GetConditionCourses(ctx *gin.Context) {
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

func (svr *ServeWrapper) AddCommonCourses(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var courses []model.CommonCourse
	ctx.BindJSON(&courses)
	for i := range courses{
		courses[i].UpdateTime = time.Now()
	}
	dao.AddCommonCourses(courses)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) AddTrialCourses(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var courses []model.TrialCourse
	ctx.BindJSON(&courses)
	now := time.Now()
	for _, c := range courses{
		c.UpdateTime = now
	}
	dao.AddTrialCourses(courses)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) AddSingleCourses(ctx *gin.Context)  {
	result := make(map[string]interface{})
	var courses []model.SingleCourse
	ctx.BindJSON(&courses)
	now := time.Now()
	for _, c := range courses{
		c.UpdateTime = now
	}
	dao.AddSingleCourses(courses)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}
