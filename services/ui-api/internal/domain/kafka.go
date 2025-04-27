package domain

import "context"

type UploadEvent struct {
	DocumentID string
	FileName   string
	ObjectURL  string
	RequestID  string
}

type KafkaProducer interface {
	PublishUpload(ctx context.Context, event UploadEvent) error
}
