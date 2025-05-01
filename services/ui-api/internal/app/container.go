package app

import (
	"context"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/domain"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
	sl "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog"
	mockService "github.com/svoevolin/semantic-search/services/ui-api/internal/mock"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/service"
)

// Container is alphabetically ordered
type Container struct {
	Config          *config.App
	Logger          logger.Logger
	DocumentService domain.DocumentService
}

func NewContainer(ctx context.Context, cfg config.App) (*Container, error) {
	c := &Container{}
	err := c.initContainer(ctx, cfg)
	return c, err
}

//nolint:unparam
func (c *Container) initContainer(_ context.Context, cfg config.App) error {
	const op = "internal.app.container.initContainer"

	c.Config = &cfg
	c.Logger = sl.NewLogger(c.Config, sl.NewAttribute("service", "ui-api"))

	//clientBuilder := httpclient.NewBuilder(httpclient.BuilderConfig{Logging: c.Config.LoggingOutgoingReqEnable}).
	//	WithLogging(c.Logger).
	//	WithRequestID()
	//
	//// Adapter
	//searchClient := adapter.NewSearcherClient(clientBuilder.Create(
	//	httpclient.Config{Timeout: c.Config.Searcher.Timeout}), c.Logger, c.Config,
	//)
	//
	//storageClient, err := adapter.NewMinioClient(c.Logger, c.Config)
	//if err != nil {
	//	return fmt.Errorf("%s: %w", op, err)
	//}
	//
	//kafkaProducer := adapter.NewKafkaAdapter(c.Config, c.Logger)

	mockStorage := &mockService.MockStorage{}
	mockKafka := &mockService.MockKafkaProducer{}
	mockSearcher := &mockService.MockSearcherClient{}

	// Service
	c.DocumentService = service.NewDocument(mockSearcher, mockStorage, mockKafka, c.Logger)

	return nil
}
