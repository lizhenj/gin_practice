package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main7() {
	r := gin.Default()

	r.GET("/chch", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,
			"https://www.sogo.com")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //修改请求的url
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9090")
}
