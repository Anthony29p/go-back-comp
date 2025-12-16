package file_infrastructure

import (
	"go-back-comp/internal/files/application"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Inyección de dependencias
	repo := NewFileRepository()
	service := application.NewFileService(repo)
	handler := NewFileHandler(service)

	filesGroup := r.Group("/files")
	{
		filesGroup.GET("", handler.GetFiles)
		filesGroup.POST("/presigned/upload", handler.PostPresignedUpload)
		filesGroup.POST("/presigned/download", handler.PostPresignedDownload)
	}
}
