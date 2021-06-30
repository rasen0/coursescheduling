package server

import (
	"coursesheduling/lib/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (svr *ServeWrapper) QueryTeacherByKey(ctx *gin.Context)  {
	//var err error
	result := make(map[string]interface{})
	//defer func() {
	//	if err != nil{
	//		log.Error("[QueryByKey]",err)
	//		result["status"] = err
	//		ctx.JSON(http.StatusInternalServerError,result)
	//	}
	//}()
	//queryWord := ctx.GetString()
	queryWord := ctx.Query("query_word")
	log.Print("query_word args:",queryWord)
	teachers := dao.QueryTeacherByWord(queryWord)
	type teacherOption struct{
		Name string `json:"name"`
		ID string `json:"id"`
	}
	teacherOptions := make([]teacherOption,0)
	for _, t := range teachers {
		teacherOptions = append(teacherOptions, teacherOption{
			t.Name,
			t.ID,
		})
	}
	result["teacher_options"] = teacherOptions
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) QueryGroupByKey(ctx *gin.Context)  {
	result := make(map[string]interface{})
	queryWord := ctx.Query("query_word")
	log.Print("query_word args:",queryWord)
	groups := dao.QueryGroupByWord(queryWord)
	type groupOption struct{
		Name string `json:"name"`
		ID string `json:"id"`
	}
	groupOptions := make([]groupOption,0)
	for _, t := range groups {
		groupOptions = append(groupOptions, groupOption{
			t.GroupName,
			t.ID,
		})
	}
	result["student_groups"] = groupOptions
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) QueryPlanByKey(ctx *gin.Context)  {
	result := make(map[string]interface{})
	queryWord := ctx.Query("query_word")
	log.Print("query_word args:",queryWord)
	plans := dao.QueryPlanByWord(queryWord)
	type planOption struct{
		Name string `json:"name"`
		ID uint `json:"id"`
	}
	planOptions := make([]planOption,0)
	for _, t := range plans {
		planOptions = append(planOptions, planOption{
			t.Name,
			t.ID,
		})
	}
	result["coursePlanOptions"] = planOptions
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) QueryCurriculumByKey(ctx *gin.Context)  {
	result := make(map[string]interface{})
	queryWord := ctx.Query("query_word")
	log.Print("query_word args:",queryWord)
	curriculums := dao.QueryCurriculumByWord(queryWord)
	type curriculumOption struct{
		Name string `json:"name"`
		ID uint `json:"id"`
	}
	curriculumOptions := make([]curriculumOption,0)
	for _, t := range curriculums {
		curriculumOptions = append(curriculumOptions, curriculumOption{
			t.Name,
			t.ID,
		})
	}
	result["curriculum_options"] = curriculumOptions
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) QueryRoomByKey(ctx *gin.Context)  {
	result := make(map[string]interface{})
	queryWord := ctx.Query("query_word")
	log.Print("query_word args:",queryWord)
	rooms := dao.QueryRoomByWord(queryWord)
	type roomOption struct{
		Name string `json:"name"`
		Address string `json:"address"`
	}
	roomOptions := make([]roomOption,0)
	for _, t := range rooms {
		roomOptions = append(roomOptions, roomOption{
			t.RoomName,
			t.Address,
		})
	}
	result["room_options"] = roomOptions
	ctx.JSON(http.StatusOK,result)
	return
}
