package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name":    "含苯苯",
		//	"message": "hello world",
		//	"age":     22,
		//}
		data1 := gin.H{"name": "李莉莉"}
		c.JSON(http.StatusOK, data1)
	})

	r.Run(":9090")
}
