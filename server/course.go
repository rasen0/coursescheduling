package server

import (
	"coursesheduling/lib/dao"
	"net/http"
	"time"

	"coursesheduling/model"
	"github.com/gin-gonic/gin"
)

func (svr *ServeWrapper) GetConditionCourses(ctx *gin.Context) {
	var queryData model.QueryData
	ctx.BindJSON(&queryData)

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
