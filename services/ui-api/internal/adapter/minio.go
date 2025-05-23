package adapter

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
)

// compile-time check
var _ domain.StorageUploader = (*MinioClient)(nil)

const DefaultHTTPScheme = "http"

type MinioClient struct {
	client     *minio.Client
	bucketName string
	logger     logger.Logger
}

func NewMinioClient(logger logger.Logger, cfg *config.App) (*MinioClient, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	bucket := cfg.Bucket

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &MinioClient{client: client, bucketName: bucket, logger: logger}, nil
}

func (m *MinioClient) Upload(ctx context.Context, file *multipart.FileHeader) (domain.UploadResult, error) {
	const op = "adapter.MinioClient.Upload"
	m.logger.DebugContext(ctx, op, "starting upload", "filename", file.Filename)

	src, err := file.Open()
	if err != nil {
		m.logger.ErrorContext(ctx, op, "file open failed", sl.Err(err))
		return domain.UploadResult{}, err
	}
	//nolint:errcheck
	//noinspection GoUnhandledErrorResult
	defer src.Close()

	objectName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(file.Filename))

	info, err := m.client.PutObject(ctx, m.bucketName, objectName, src, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		m.logger.ErrorContext(ctx, op, "upload failed", sl.Err(err))
		return domain.UploadResult{}, err
	}

	objectURL := &url.URL{
		Scheme: DefaultHTTPScheme,
		Host:   m.client.EndpointURL().Host,
		Path:   fmt.Sprintf("/%s/%s", m.bucketName, info.Key),
	}
	m.logger.DebugContext(ctx, op, "upload succeeded", "object_name", info.Key, "url", objectURL.String())

	return domain.UploadResult{
		ObjectName: info.Key,
		URL:        objectURL,
	}, nil
}
