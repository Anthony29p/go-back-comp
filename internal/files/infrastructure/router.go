package file_infrastructure

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	filesGroup := r.Group("/files")
	{
		filesGroup.GET("", GetFiles)
	}
}
