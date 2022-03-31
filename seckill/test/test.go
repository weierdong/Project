package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"a": "200",
		})
	})
	r.GET("test", func(context *gin.Context) {

		A := context.PostForm("A")
		context.JSON(http.StatusOK, gin.H{
			"A": A,
		})
	})

	r.Run(":8090")
}
