package file_infrastructure

import "github.com/gin-gonic/gin"

func GetFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World 2",
	})
}
