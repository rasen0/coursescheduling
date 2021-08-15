package server

import (
	"coursesheduling/common"
	"coursesheduling/lib/biz"
	"coursesheduling/lib/dao"
	"coursesheduling/lib/log"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type loginInfo struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (svr *ServeWrapper) PostLogin(ctx *gin.Context) {
	result := make(map[string]interface{})
	var loginData loginInfo
	ctx.Bind(&loginData)
	log.Printf("login %v", loginData)
	if !biz.VerificationPassword(loginData.Name,loginData.Pass){
		result["error"] = "登陆账号验证失败！"
		ctx.JSON(http.StatusOK,result)
	}
	token, err := biz.MakeToken(loginData.Name)
	if err != nil{
		log.Error("err:",err)
		result["error"] = "get token fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}
	dao.UpdateAccount(model.Account{
		UserName: loginData.Name,
		Password: loginData.Pass,
		Token: token,
		Update: time.Now().Format(common.TimeFormat),
	})
	result["result"] = "ok"
	result["token"] = token
	result["error"] = ""
	ctx.JSON(http.StatusOK,result)
	return
}

func (svr *ServeWrapper) PostRegister(ctx *gin.Context) {
	result := make(map[string]interface{})
	var register = struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
		CheckPass string `json:"checkPass"`
	}{}
	ctx.Bind(&register)
	log.Printf("register ", register)
	biz.RegisterAccount(register.Name,register.Pass)
	result["result"] = "ok"
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
