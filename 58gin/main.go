package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User   string `form:user`
	Passwd string `form:passwd`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		username := c.PostForm("username")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":   login.User,
				"passwd": login.Passwd,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	//路由组
	rg := r.Group("/this")
	{
		rg.GET("/one", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "this is this/one",
			})
		})
		rg.GET("/two", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "this is this/two",
			})
		})

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "has no route",
		})
	})
	r.Run(":9090")
}
