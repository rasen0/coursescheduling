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

type conflictCourse struct{
	TeacherName string
	TeacherId string
	StartTime string
	EndTime string
}

func (svr *ServeWrapper) AddCommonCourses(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.CommonCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData []model.CommonCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse)
	for i := range reqCourse.RequestData{
		reqCourse.RequestData[i].UpdateTime = time.Now()
	}
	// 查询课程时间是否有冲突课程
	conflictCourses := make([]conflictCourse,0)
	addingCourses := make([]model.CommonCourse,0)
	for _,c := range reqCourse.RequestData {
		if dao.CheckCourse(c.StartTime,c.EndTime,c.TeacherID){
			conflictCourses = append(conflictCourses,conflictCourse{
				TeacherName: c.TeacherName,
				TeacherId: c.TeacherID,
				StartTime: c.StartTime,
				EndTime: c.EndTime,
			})
		}else{
			addingCourses = append(addingCourses,c)
		}
	}
	dao.AddCommonCourses(addingCourses)
	result["conflict_course"] = conflictCourses
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) DelCommonCourse(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.CommonCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData model.CommonCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse)
	// todo commoncourse 删除课程不存在则调用删除trialcourse
	dao.DelCommonCourse(reqCourse.RequestData)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) AddTrialCourses(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.TrialCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData []model.TrialCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse)
	now := time.Now()
	for _, c := range reqCourse.RequestData{
		c.UpdateTime = now
	}
	// 查询课程时间是否有冲突课程
	conflictCourses := make([]conflictCourse,0)
	addingCourses := make([]model.TrialCourse,0)
	for _,c := range reqCourse.RequestData {
		if dao.CheckCourse(c.StartTime,c.EndTime,c.TeacherID){
			conflictCourses = append(conflictCourses,conflictCourse{
				TeacherName: c.TeacherName,
				TeacherId: c.TeacherID,
				StartTime: c.StartTime,
				EndTime: c.EndTime,
			})
		}else{
			addingCourses = append(addingCourses,c)
		}
	}
	dao.AddTrialCourses(addingCourses)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) DelTrialCourse(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.CommonCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData model.TrialCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse)
	dao.DelTrialCourse(reqCourse.RequestData)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) AddSingleCourses(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.SingleCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData []model.SingleCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse.RequestData)
	now := time.Now()
	for _, c := range reqCourse.RequestData{
		c.UpdateTime = now
	}
	// 查询课程时间是否有冲突课程
	conflictCourses := make([]conflictCourse,0)
	addingCourses := make([]model.SingleCourse,0)
	for _,c := range reqCourse.RequestData {
		if dao.CheckCourse(c.StartTime,c.EndTime,c.TeacherID){
			conflictCourses = append(conflictCourses,conflictCourse{
				TeacherName: c.TeacherName,
				TeacherId: c.TeacherID,
				StartTime: c.StartTime,
				EndTime: c.EndTime,
			})
		}else{
			addingCourses = append(addingCourses,c)
		}
	}
	dao.AddSingleCourses(addingCourses)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) DelSingleCourse(ctx *gin.Context) {
	result := make(map[string]interface{})
	//var courses []model.CommonCourse
	var reqCourse = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData model.SingleCourse `json:"request_data"`
	}{}
	ctx.BindJSON(&reqCourse)
	dao.DelSingleCourse(reqCourse.RequestData)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	return
}
