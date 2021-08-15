package server

import (
	"coursesheduling/lib/biz"
	"coursesheduling/lib/dao"
	"coursesheduling/lib/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func authority() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.GetHeader("token")
		token, _ := biz.ParseToken(tokenStr)
		if token == nil{
			context.AbortWithStatus(http.StatusBadRequest)
		}
		name := util.ConvertValue(token["username"])
		log.Print("token entry:",token,"name:",name)
		accountVal := dao.QueryAccountByName(name)
		//if accountVal.Token == common.FreeToken && tokenStr == ""{
		//	context.Next()
		//}
		if accountVal.Token != tokenStr {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		context.Next()
	}
}

func Common() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.Writer.Header()
		// alone dns prefect
		header.Set("X-DNS-Prefetch-Control", "on")
		// IE No Open
		header.Set("X-Download-Options", "noopen")
		// not cache
		header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		header.Set("Expires", "max-age=0")
		// Content Security Policy
		header.Set("Content-Security-Policy", "default-src 'self'")
		// xss protect
		// it will caught some problems is old IE
		header.Set("X-XSS-Protection", "1; mode=block")
		// Referrer Policy
		header.Set("Referrer-Header", "no-referrer")
		// cros frame, allow same origin
		header.Set("X-Frame-Options", "SAMEORIGIN")
		// HSTS
		header.Set("Strict-Transport-Security", "max-age=5184000;includeSubDomains")
		// no sniff
		header.Set("X-Content-Type-Options", "nosniff")
		context.Next()
	}
}
