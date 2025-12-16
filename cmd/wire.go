//go:build wireinject
// +build wireinject

package main

import (
	file_application "go-back-comp/internal/files/application"
	file_infrastructure "go-back-comp/internal/files/infrastructure"
	shared_domain "go-back-comp/internal/shared/domain"
	shared_infrastructure "go-back-comp/internal/shared/infrastructure"

	"github.com/google/wire"
)

func provideS3Storage() (shared_domain.FileStorage, error) {
	return shared_infrastructure.NewS3Provider("my-bucket-name")
}

func InitializeFileHandler() (*file_infrastructure.FileHandler, error) {
	wire.Build(
		file_infrastructure.NewFileRepository,
		provideS3Storage,
		file_application.NewFileService,
		file_infrastructure.NewFileHandler,
	)
	return &file_infrastructure.FileHandler{}, nil
}
