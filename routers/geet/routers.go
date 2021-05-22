package geet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.GET("/hello", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}
