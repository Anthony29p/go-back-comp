package file_infrastructure

import "go-back-comp/internal/files/domain"

type fileRepository struct{}

// NewFileRepository crea una nueva instancia del repositorio
func NewFileRepository() domain.FileRepository {
	return &fileRepository{}
}

func (r *fileRepository) GetAllFiles() string {
	return "Hello World 2"
}

func (r *fileRepository) GetPresignedUploadURL() string {
	return "presigned upload URL"
}

func (r *fileRepository) GetPresignedDownloadURL() string {
	return "presigned download URL"
}
