package domain

import (
	"context"
	"mime/multipart"
	"net/url"
)

type UploadResult struct {
	ObjectName string
	URL        *url.URL
}

type StorageUploader interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (UploadResult, error)
}
