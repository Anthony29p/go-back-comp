package file_infrastructure

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, handler *FileHandler) {
	filesGroup := r.Group("/files")
	{
		filesGroup.GET("", handler.GetFiles)
		filesGroup.POST("/presigned/upload", handler.PostPresignedUpload)
		filesGroup.POST("/presigned/download", handler.PostPresignedDownload)
	}
}
