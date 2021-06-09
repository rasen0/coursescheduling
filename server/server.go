package server

import (
	"coursesheduling/lib/config"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (svr *ServeWrapper) Serve()  {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(Cors())
	svr.Server = &http.Server{
		IdleTimeout: 10,
		ReadTimeout: 15,
		WriteTimeout: 15,
		Addr: svr.Address,
		Handler: engine,
	}
	serveGroup := engine.Group("/serve/v1")
	serveGroup.GET("/courseschduling",svr.GetCourseScheduling)

	go svr.ListenAndServe()
}

/*
GetCourseScheduling 默认请求当月课程安排
month 请求月份
*/
func (svr *ServeWrapper) GetCourseScheduling(ctx *gin.Context)  {

}