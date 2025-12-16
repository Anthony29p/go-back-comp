package file_infrastructure

import "github.com/gin-gonic/gin"

func GetFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World 2",
	})
}

func PostPresignedUpload(c *gin.Context) {
	c.String(200, "presigned upload URL")
}

func PostPresignedDownload(c *gin.Context) {
	c.String(200, "presigned download URL")
}
