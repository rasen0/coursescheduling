package server

import (
	"coursesheduling/lib/biz"
	"coursesheduling/lib/dao"
	"coursesheduling/lib/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginInfo struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (svr *ServeWrapper) PostLogin(ctx *gin.Context) {
	result := make(map[string]interface{})
	var loginData loginInfo
	ctx.Bind(&loginData)
	log.Printf("login ", loginData)
	token, err := biz.MakeToken(loginData.Name)
	if err != nil{
		log.Error("err:",err)
		result["error"] = "get token fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}
	//account := dao.QueryAccountByName(loginData.Name)
	//if account.Password != loginData.Pass{
	//	result["result"] = "fail"
	//	ctx.JSON(http.StatusOK,result)
	//	return
	//}

	result["result"] = "ok"
	result["token"] = token
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) GetAccounts(ctx *gin.Context)  {
	result := make(map[string]interface{})
	log.Printf("get account")
	accounts := dao.QueryAccounts()
	result["accounts"] = accounts
	ctx.JSON(http.StatusOK,result)
	return
}
