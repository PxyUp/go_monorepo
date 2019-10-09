package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(appName string) func(c *gin.Context){
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Hello from " + appName,
		})
	}
}
