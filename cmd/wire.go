//go:build wireinject
// +build wireinject

package main

import (
	"go-back-comp/internal/files/application"
	file_infrastructure "go-back-comp/internal/files/infrastructure"

	"github.com/google/wire"
)

func InitializeFileHandler() *file_infrastructure.FileHandler {
	wire.Build(
		file_infrastructure.NewFileRepository,
		application.NewFileService,
		file_infrastructure.NewFileHandler,
	)
	return &file_infrastructure.FileHandler{}
}
