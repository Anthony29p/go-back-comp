package file_infrastructure

import (
	"go-back-comp/internal/files/domain"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service domain.FileService
}

// NewFileHandler crea una nueva instancia del handler
func NewFileHandler(service domain.FileService) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func (h *FileHandler) GetFiles(c *gin.Context) {
	result := h.service.GetFiles()
	c.JSON(200, gin.H{
		"message": result,
	})
}

func (h *FileHandler) PostPresignedUpload(c *gin.Context) {
	var request struct {
		FileName string `json:"fileName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "fileName is required",
		})
		return
	}

	result, err := h.service.GetPresignedUpload(request.FileName)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to generate presigned URL",
		})
		return
	}

	c.JSON(200, result)
}

func (h *FileHandler) PostPresignedDownload(c *gin.Context) {
	result := h.service.GetPresignedDownload()
	c.String(200, result)
}
