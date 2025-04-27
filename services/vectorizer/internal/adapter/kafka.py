import asyncio
import json
from typing import Callable, Awaitable

from aiokafka import AIOKafkaConsumer
from pydantic import ValidationError
import structlog
from structlog.contextvars import bind_contextvars, clear_contextvars

from internal.config.config import AppConfig
from internal.domain.document import UploadEvent


class KafkaConsumerAdapter:
    """Адаптер для чтения сообщений из Kafka и передачи их в бизнес-логику."""

    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._config = config
        self._logger = logger
        self._consumer: AIOKafkaConsumer | None = None

    async def connect(self) -> None:
        op = "adapter.kafka_consumer.connect"

        self._consumer = AIOKafkaConsumer(
            self._config.kafka_topic,
            bootstrap_servers=self._config.kafka_broker,
            auto_offset_reset="earliest",
            group_id="vectorizer-consumer-group",
        )
        await self._consumer.start()
        self._logger.debug("Kafka consumer connected", operation=op)

    async def consume(self, handler: Callable[[UploadEvent], Awaitable[None]]) -> None:
        op = "adapter.kafka_consumer.consume"

        if self._consumer is None:
            raise RuntimeError("Kafka consumer is not connected")

        try:
            async for msg in self._consumer:
                await self._handle_message(msg, handler)
        finally:
            await self._consumer.stop()
            self._logger.debug("Kafka consumer stopped", operation=op)

    async def _handle_message(self, msg, handler: Callable[[UploadEvent], Awaitable[None]]) -> None:
        op = "adapter.kafka_consumer.handle_message"

        try:
            raw_payload = json.loads(msg.value)

            try:
                payload = UploadEvent.model_validate(raw_payload)

                clear_contextvars()
                if payload.request_id:
                    bind_contextvars(request_id=payload.request_id)

            except ValidationError as ve:
                self._logger.error(
                    "Validation error on incoming Kafka message",
                    operation=op,
                    error=str(ve),
                    raw_payload=raw_payload,
                )
                return

            self._logger.debug(
                "Received valid message",
                operation=op,
                document_id=payload.document_id,
            )

            await handler(payload)

            self._logger.debug(
                "Message processed successfully",
                operation=op,
                document_id=payload.document_id,
            )

        except Exception as e:
            self._logger.error(
                "Unexpected error while processing Kafka message",
                operation=op,
                error=str(e),
            )
