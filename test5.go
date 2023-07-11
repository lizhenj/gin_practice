package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main6() {
	r := gin.Default()
	r.LoadHTMLFiles("./index1.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := fmt.Sprintf("./%s", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9090")
}
