package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	r.GET("/hi", hiHandler)
}

func hiHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi q1mi!",
	})
}
