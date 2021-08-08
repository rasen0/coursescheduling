package server

import (
	"coursesheduling/common"
	"coursesheduling/lib/dao"
	"coursesheduling/lib/log"
	"net/http"
	"time"

	"coursesheduling/model"
	"github.com/gin-gonic/gin"
)

func (svr *ServeWrapper) GetConditionCourses(ctx *gin.Context) {
	var queryData model.QueryData
	ctx.BindJSON(&queryData)
	if queryData.CourseDate == "" {
		queryData.CourseDate = time.Now().Format(common.CalendarFormat)
	}else{
		tParse, _ := time.Parse(time.RFC3339, queryData.CourseDate)
		t1 := tParse.Add(24 * time.Hour)
		queryData.CourseDate = t1.Format(common.CalendarFormat)
	}
	log.Printf("GetConditionCourses| query data:%s",queryData.CourseDate)
	result := make(map[string]interface{})
	_, courseTable := dao.GetCourseTableByCond(queryData)

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
