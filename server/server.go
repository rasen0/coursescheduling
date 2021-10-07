package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"coursesheduling/lib/config"
	"coursesheduling/lib/dao"
	"coursesheduling/model"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

type ServeWrapper struct {
	*http.Server
	*config.Configure
}

func NewServer(config *config.Configure) *ServeWrapper {
	return &ServeWrapper{
		Configure:config,
	}
}

func TlsHandler(addr string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     addr,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

func (svr *ServeWrapper) Serve()  {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	//engine.Use(TlsHandler(svr.Address)) // ssl
	engine.Use(gin.Recovery())
	engine.Use(Cors())
	engine.Use(Common())
	log.Print("addr:",svr.Address)
	svr.Server = &http.Server{
		IdleTimeout: 10*time.Second,
		ReadTimeout: 15*time.Second,
		WriteTimeout: 15*time.Second,
		Addr: svr.Address,
		Handler: engine,
	}
	engine.StaticFS("/home", assetFS())
	route := engine.Group("/api")
	serveGroup := route.Group("/service")
	serveGroup.GET("/v1/curriculumOptions",svr.GetCurriculumOptions)
	serveGroup.GET("/v1/coursePlanOptions",svr.GetCoursePlanOptions)
	serveGroup.GET("/v1/coursescheduling",svr.GetCourseScheduling)  // old
	serveGroup.GET("/v1/querystudentbyid",svr.GetStudentsByID)
	serveGroup.GET("/v1/querystudentbykey",svr.GetStudentsByKey)
	serveGroup.GET("/v1/queryteacherbykey",svr.QueryTeacherByKey)
	serveGroup.GET("/v1/querygroupbykey",svr.QueryGroupByKey)
	serveGroup.GET("/v1/queryplanbykey",svr.QueryPlanByKey)
	serveGroup.GET("/v1/querycurriculumebykey",svr.QueryCurriculumByKey)
	serveGroup.GET("/v1/queryroombykey",svr.QueryRoomByKey)
	serveGroup.GET("/v1/queryrolebykey",svr.QueryRoleByKey)
	serveGroup.GET("/v1/getrooms",svr.GetRooms)
	serveGroup.GET("/v1/getaccounts",svr.GetAccounts)

	serveGroup.POST("/v1/groups",svr.GetGroups)
	serveGroup.POST("/v1/register",svr.PostRegister)
	serveGroup.POST("/v1/login",svr.PostLogin)

	//serveGroup.GET("/v1/students",svr.GetCourseScheduling)
	safeGroup := route.Group("/safe")
	safeGroup.Use(authority())

	safeGroup.POST("/v1/getteachers",svr.GetTeachers)
	safeGroup.POST("/v1/coursescheduling",svr.PostCourseScheduling)
	safeGroup.POST("/v1/getstudents",svr.StudentPagination)
	safeGroup.POST("/v1/grouppagination",svr.GroupPagination)
	safeGroup.POST("/v1/querycoursescondition",svr.GetConditionCourses)

	safeGroup.POST("/v1/addingstudent",svr.AddStudent)
	safeGroup.POST("/v1/addingteacher",svr.AddTeacher)
	safeGroup.POST("/v1/addcommoncourses",svr.AddCommonCourses)
	safeGroup.POST("/v1/delcommoncourse",svr.DelCommonCourse)
	safeGroup.POST("/v1/addtrialcourses",svr.AddTrialCourses)
	safeGroup.POST("/v1/deltrialcourse",svr.DelTrialCourse)
	safeGroup.POST("/v1/addsinglecourses",svr.AddSingleCourses)
	safeGroup.POST("/v1/delsinglecourse",svr.DelSingleCourse)
	safeGroup.POST("/v1/addingroom",svr.AddingRoom)
	safeGroup.POST("/v1/addingaccount",svr.AddingAccount)
	safeGroup.POST("/v1/addinggroup",svr.AddingGroup)

	fmt.Println("start course scheduling system")
	go svr.ListenAndServe()
	//go svr.Server.ListenAndServeTLS("","") // https
}

type requestData struct {
	Operator string `json:"operator"`
	DataType string `json:"data_type"`
	Active string `json:"active"`
	Month string `json:"month"`
}

/*
GetTypeCourseScheduling 默认请求当月课程安排
month 请求月份
ctype 请求安排课程类型
*/
func (svr *ServeWrapper) GetTypeCourseScheduling(ctx *gin.Context)  {
	monthStr := ctx.GetString("month")
	var month time.Time
	if monthStr <= ""{
		month = time.Now()
	}else {
		month,_ = time.Parse(time.RFC3339,monthStr)
	}
	fmt.Println(month)
	ctype := ctx.GetInt("ctype")
	if ctype <= 0 {
		ctype = model.CommonLesson
	}
	result := make(map[string]interface{})
	var err error
	defer func() {
		if err == nil{
			return
		}
		fmt.Println("fail")
		result["status"]="fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}()

	_, courseTable := dao.GetTypeCourseTable(ctype, month)
	result["courseTable"] =courseTable
	ctx.JSON(http.StatusOK,result)
	return
}

/*
GetCourseScheduling 默认请求当月课程安排
month 请求月份
*/
func (svr *ServeWrapper) GetCourseScheduling(ctx *gin.Context) {
	inTime := time.Now()
	log.Print("[GetCourseScheduling] start time ",inTime)
	monthStr := ctx.GetString("month")
	var month time.Time
	if monthStr <= ""{
		month = time.Now()
	}else {
		month,_ = time.Parse(time.RFC3339,monthStr)
	}
	log.Print("month:",month)
	result := make(map[string]interface{})
	var err error
	defer func() {
		if err == nil{
			return
		}
		fmt.Println("fail")
		result["status"]="fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}()

	_, courseTable := dao.GetCourseTable(month)
	result["courseTable"] =courseTable
	ctx.JSON(http.StatusOK,result)
	outTime := time.Now()
	log.Print("[GetCourseScheduling] end time ",outTime,". Difference time:",outTime.Sub(inTime))
	return
}

/*
PostCourseScheduling 默认请求当月课程安排
month 请求月份
*/
func (svr *ServeWrapper) PostCourseScheduling(ctx *gin.Context) {
	inTime := time.Now()
	log.Print("[PostCourseScheduling] start time ",inTime)
	data := requestData{}
	ctx.Bind(&data)
	var month time.Time
	if data.Month <= ""{
		month = time.Now()
	}else {
		month,_ = time.Parse(time.RFC3339,data.Month)
	}
	log.Print("month:",month)
	result := make(map[string]interface{})
	var err error
	defer func() {
		if err == nil{
			return
		}
		fmt.Println("fail")
		result["status"]="fail"
		ctx.JSON(http.StatusInternalServerError,result)
	}()
	//ok,_ := dao.VerifyPolicy(data.Operator,data.DataType,common.ReadActive)
	//if !ok {
	//	err = errors.New("verify fail")
	//	return
	//}
	account := dao.QueryAccountByName(data.Operator)
	_, courseTable := dao.GetCourseTableWithID(month,account.TeacherID)
	result["courseTable"] =courseTable
	ctx.JSON(http.StatusOK,result)
	outTime := time.Now()
	log.Print("[GetCourseScheduling] end time ",outTime,". Difference time:",outTime.Sub(inTime))
	return
}

type Pagination struct {
	Total int `json:"total"`
	PageSize int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	PageSizes []int `json:"page_sizes"`
}
