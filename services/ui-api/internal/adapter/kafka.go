package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
)

// compile-time check
var _ domain.KafkaProducer = (*KafkaAdapter)(nil)

type KafkaAdapter struct {
	writer *kafka.Writer
	topic  string
	logger logger.Logger
}

func NewKafkaAdapter(cfg *config.App, logger logger.Logger) *KafkaAdapter {
	const (
		op       = "adapter.KafkaAdapter.New"
		retries  = 10
		interval = 3 * time.Second
	)
	logger.Debug(op, "Connecting to Kafka broker", "broker", cfg.Kafka.Broker)

	var lastErr error
	for i := 0; i < retries; i++ {
		_, err := kafka.DialLeader(context.Background(), "tcp", cfg.Kafka.Broker, cfg.Kafka.Topic, 0)
		lastErr = err
		if err == nil {
			break
		}
		logger.Warn(op, fmt.Sprintf("Kafka not ready (attempt %d/%d)", i+1, retries), "err", sl.Err(err))
		time.Sleep(interval)
	}

	if lastErr != nil {
		logger.Panic(op, "Kafka not available after retries", "err", sl.Err(lastErr))
	}

	return &KafkaAdapter{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.Kafka.Broker),
			Topic:    cfg.Kafka.Topic,
			Balancer: &kafka.LeastBytes{},
		},
		topic:  cfg.Kafka.Topic,
		logger: logger,
	}
}

func (k *KafkaAdapter) PublishUpload(ctx context.Context, event domain.UploadEvent) error {
	const op = "adapter.KafkaAdapter.PublishUpload"

	data, err := json.Marshal(event)
	if err != nil {
		k.logger.ErrorContext(ctx, op, "marshal failed", sl.Err(err))
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.DocumentID),
		Value: data,
	}

	err = k.writer.WriteMessages(ctx, msg)
	if err != nil {
		k.logger.ErrorContext(ctx, op, "write message failed", sl.Err(err))
		return err
	}

	k.logger.DebugContext(ctx, op, "message sent", "doc_id", event.DocumentID)
	return nil
}
