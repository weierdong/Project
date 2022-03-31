package engine

import (
	"github.com/gin-gonic/gin"
	//"seckill/api"
)

func EngineStart() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/html/*.html")
	router.Static("/static", "static")
	//	router.GET("/", Index)
	router.GET("/join", JoinUs)
	router.POST("/join", JoinUsOk)
	router.GET("/login", Login)
	router.POST("/login", LoginOK)
	router.GET("/", Index)

	return router

}
