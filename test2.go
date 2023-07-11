package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//form表哥提交参数
func main3() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(http.StatusOK, map[string]interface{}{
			"username": username,
			"password": password,
		})
	})
	r.Run(":9090")
}
