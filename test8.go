package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "inedx"})
}

//中间件m1，统计请求处理函数耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Set("name", "HHHHH")
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Abort() //阻止后续处理函数
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {

		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default() //默认使用Logger和Recovery
	//r := gin.New()

	//r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数m1

	r.GET("/indexx", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})
	r.GET("/user", m1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "user"})
	})

	xxGroup := r.Group("/xx", authMiddleware(false))
	{
		xxGroup.Use(m1)
		xxGroup.GET("/index", func(c *gin.Context) {
			name := c.MustGet("name")
			c.JSON(http.StatusOK, gin.H{"msg": name})
		})
	}
	r.Run(":9090")
}
