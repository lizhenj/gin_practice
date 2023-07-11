package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main2() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器请求发起的query string参数
		//name := c.Query("query")
		//name := c.DefaultQuery("query", "somebody") //第二参数表默认
		name, _ := c.GetQuery("query") //返回bool表存在否
		c.JSON(http.StatusOK, name)
	})
	r.Run(":9090")
}
