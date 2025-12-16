package routes

import (
	file_infrastructure "go-back-comp/internal/files/infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *gin.Engine, fileHandler *file_infrastructure.FileHandler) {
	file_infrastructure.RegisterRoutes(r, fileHandler)
}
