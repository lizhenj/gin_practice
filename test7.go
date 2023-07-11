package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main8() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "post"})
		}
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "404"})
	})

	//视频首页与详情页
	//r.GET("/vedio/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "video"})
	//})
	//路由组的组
	videoGroup := r.Group("/vedio")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "video/index"})
		})
		videoGroup.GET("/msg", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "video/msg"})
		})
	}

	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})

	r.Run(":9090")
}
