package originator

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) *gin.Engine {

	r.GET("/ping/originator", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
