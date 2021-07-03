package server

import (
	"coursesheduling/lib/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
