package shared_domain

import "context"

type FileStorage interface {
	GeneratePresignedUploadURL(ctx context.Context, fileName string) (fileID string, presignedURL string, err error)
	GeneratePresignedDownloadURL(ctx context.Context, fileID string, fileName string) (presignedURL string, err error)
}
