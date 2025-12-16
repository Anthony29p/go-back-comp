package file_application

import "go-back-comp/internal/files/domain"

type fileService struct {
	repo domain.FileRepository
}

// NewFileService crea una nueva instancia del servicio
func NewFileService(repo domain.FileRepository) domain.FileService {
	return &fileService{
		repo: repo,
	}
}

func (s *fileService) GetFiles() string {
	return s.repo.GetAllFiles()
}

func (s *fileService) GetPresignedUpload() string {
	return s.repo.GetPresignedUploadURL()
}

func (s *fileService) GetPresignedDownload() string {
	return s.repo.GetPresignedDownloadURL()
}
