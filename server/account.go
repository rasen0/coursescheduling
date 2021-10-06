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

type accountForm struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Role string `json:"role"`
}

func (svr *ServeWrapper) AddingAccount(ctx *gin.Context) {
	result := make(map[string]interface{})
	var reqAccount struct{
		Operator string `json:"operator"`
		DataType string `json:"data_type"`
		Active string `json:"active"`
		AccountForm accountForm
	}
	ctx.Bind(&reqAccount)
	account := model.Account{
		UserName:reqAccount.AccountForm.UserName,
		Password: reqAccount.AccountForm.PassWord,
		Role: reqAccount.AccountForm.Role,
		Update: time.Now().Format(common.TimeFormat),
	}
	dao.AddAccount(account)
	ctx.JSON(http.StatusOK,result)
	return
}