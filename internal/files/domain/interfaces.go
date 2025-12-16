package domain

// PresignedURLResponse representa la respuesta de una URL prefirmada para upload
type PresignedURLResponse struct {
	FileID    string `json:"fileId"`
	UploadURL string `json:"uploadUrl"`
}

// PresignedDownloadURLResponse representa la respuesta de una URL prefirmada para download
type PresignedDownloadURLResponse struct {
	DownloadURL string `json:"downloadUrl"`
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
	GetPresignedDownload(fileID string, fileName string) (*PresignedDownloadURLResponse, error)
}
