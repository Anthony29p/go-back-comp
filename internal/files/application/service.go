package file_application

import (
	"context"
	"go-back-comp/internal/files/domain"
	shared_domain "go-back-comp/internal/shared/domain"
)

type fileService struct {
	repo    domain.FileRepository
	storage shared_domain.FileStorage
}

func NewFileService(repo domain.FileRepository, storage shared_domain.FileStorage) domain.FileService {
	return &fileService{
		repo:    repo,
		storage: storage,
	}
}

func (s *fileService) GetFiles() string {
	return s.repo.GetAllFiles()
}

func (s *fileService) GetPresignedUpload(fileName string) (*domain.PresignedURLResponse, error) {
	fileID, presignedURL, err := s.storage.GeneratePresignedUploadURL(context.Background(), fileName)
	if err != nil {
		return nil, err
	}

	return &domain.PresignedURLResponse{
		FileID:    fileID,
		UploadURL: presignedURL,
	}, nil
}

func (s *fileService) GetPresignedDownload() string {
	return s.repo.GetPresignedDownloadURL()
}
