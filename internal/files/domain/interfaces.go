package domain

// PresignedURLResponse representa la respuesta de una URL prefirmada
type PresignedURLResponse struct {
	FileID    string `json:"fileId"`
	UploadURL string `json:"uploadUrl"`
}

// FileRepository define los métodos para acceso a datos
type FileRepository interface {
	GetAllFiles() string
	GetPresignedUploadURL() string
	GetPresignedDownloadURL() string
}

// FileService define los métodos de la lógica de negocio
type FileService interface {
	GetFiles() string
	GetPresignedUpload(fileName string) (*PresignedURLResponse, error)
	GetPresignedDownload() string
}
