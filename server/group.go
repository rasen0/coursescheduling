package server

import (
	"coursesheduling/lib/dao"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (svr *ServeWrapper) AddingGroup(ctx *gin.Context)  {
	result := make(map[string]interface{})
	//var newStudent model.Student
	var reqStudent = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		RequestData model.StudentGroup `json:"request_data"`
	}{}
	ctx.BindJSON(&reqStudent)
	log.Printf("new student:%+v",reqStudent.RequestData)
	dao.InsertGroupOne(reqStudent.RequestData)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
}

func (svr *ServeWrapper) GetGroups(ctx *gin.Context)  {
	result := make(map[string]interface{})
	//var newStudent model.Student
	//var reqStudent = struct {
	//	Operator string `json:"operator"`
	//	DataType string `json:"data_type"`
	//	Active string `json:"active"`
	//}{}
	var count int
	ctx.BindJSON(&count)
	groups := dao.GetGroup()
	result["status"] = "ok"
	result["studentGroups"] = groups
	ctx.JSON(http.StatusOK,result)
}
