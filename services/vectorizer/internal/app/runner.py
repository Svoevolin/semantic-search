import structlog

from internal.adapter.kafka import KafkaConsumerAdapter
from internal.domain.document import DocumentServiceI, UploadEvent


class KafkaConsumerRunner:
    """Runner для получения сообщений из Kafka и передачи их в сервис."""

    def __init__(self, consumer: KafkaConsumerAdapter, service: DocumentServiceI, logger: structlog.BoundLogger):
        self._consumer = consumer
        self._service = service
        self._logger = logger

    async def run(self) -> None:
        op = "app.runner.kafka_consumer.run"

        await self._consumer.connect()
        self._logger.info("Kafka consumer connected", operation=op)

        await self._consumer.consume(self._handle_message)

    async def _handle_message(self, event: UploadEvent) -> None:
        op = "app.runner.kafka_consumer.handle_message"

        try:
            self._logger.debug("Handling new document event", operation=op, document_id=event.document_id)

            await self._service.process(event)

            self._logger.info("Document event processed successfully", operation=op, document_id=event.document_id)

        except Exception as e:
            self._logger.error("Failed to process document event", operation=op, error=str(e), document_id=event.document_id)
