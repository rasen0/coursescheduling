package server

import (
	"net/http"

	"coursesheduling/lib/dao"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
)

func (svr *ServeWrapper) AddingRoom(ctx *gin.Context)  {
	result := make(map[string]interface{})
	//var newRoom model.Classroom
	var reqClassroom = struct {
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		Classroom model.Classroom
	}{}
	ctx.BindJSON(&reqClassroom)
	//newRoom.UpdateTime = time.Now()
	log.Printf("new newRoom:%+v",reqClassroom.Classroom)
	dao.InsertRoomOne(reqClassroom.Classroom)
	result["status"] = "ok"
	ctx.JSON(http.StatusOK,result)
	log.Print("add teacher done")
	return
}

func (svr *ServeWrapper) GetRooms(ctx *gin.Context)  {
	result := make(map[string]interface{})
	log.Printf("get rooms")
	rooms := dao.GetRooms()
	result["rooms"] = rooms
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) GetCurriculumOptions(ctx *gin.Context)  {
	curriculums := dao.GetCurriculums()
	type curri struct {
		Name string `json:"name"`
		Value string `json:"value"`
	}
	curris := make([]curri,0)
	for _, c := range curriculums{
		curr := curri{
			c.Name,
			c.Name,
		}
		curris =append(curris,curr)
	}
	result := make(map[string]interface{})
	result["curriculumOptions"] = curris
	ctx.JSON(http.StatusOK,result)
}

func (svr *ServeWrapper) GetCoursePlanOptions(ctx *gin.Context)  {
	coursePlans := dao.GetCoursePlans()
	type cPlan struct {
		Name string `json:"name"`
		Value uint `json:"value"`
	}
	cPlans := make([]cPlan,0)
	for _, c := range coursePlans{
		cour := cPlan{
			c.Name,
			c.ID,
		}
		cPlans =append(cPlans,cour)
	}
	result := make(map[string]interface{})
	result["coursePlanOptions"] = cPlans
	ctx.JSON(http.StatusOK,result)
}
