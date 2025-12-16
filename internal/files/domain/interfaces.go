package domain

// FileRepository define los métodos para acceso a datos
type FileRepository interface {
	GetAllFiles() string
	GetPresignedUploadURL() string
	GetPresignedDownloadURL() string
}

// FileService define los métodos de la lógica de negocio
type FileService interface {
	GetFiles() string
	GetPresignedUpload() string
	GetPresignedDownload() string
}
