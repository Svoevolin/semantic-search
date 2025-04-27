from typing import Any

import structlog
from qdrant_client import AsyncQdrantClient
from qdrant_client.models import PointStruct, VectorParams, Distance

from internal.domain.document import QdrantUploader, VectorEmbedding
from internal.config.config import AppConfig

class QdrantUploaderAdapter(QdrantUploader):
    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._client = AsyncQdrantClient(host=config.qdrant_host, port=config.qdrant_port)
        self._collection_name = config.qdrant_collection
        self._embedding_size = config.qdrant_embedding_size
        self._logger = logger
        self._initialized = False

    async def initialize(self) -> None:
        if not self._initialized:
            await self._ensure_collection_exists()
            self._initialized = True

    async def _ensure_collection_exists(self) -> None:
        op = "adapter.qdrant_uploader.ensure_collection_exists"

        exists = await self._client.collection_exists(self._collection_name)
        if not exists:
            await self._client.create_collection(
                collection_name=self._collection_name,
                vectors_config=VectorParams(size=self._embedding_size, distance=Distance.COSINE),
            )
            self._logger.info("Qdrant collection created", collection=self._collection_name, operation=op)
        else:
            self._logger.info("Qdrant collection already exists", collection=self._collection_name, operation=op)

    async def upload(self, document_id: str, vector: VectorEmbedding, payload: dict[str, Any]) -> None:
        op = "adapter.qdrant_uploader.upload"

        try:
            await self._client.upsert(
                collection_name=self._collection_name,
                points=[
                    PointStruct(
                        id=document_id,
                        vector=vector.values,
                        payload=payload,
                    )
                ],
            )
            self._logger.info("Document vector uploaded to Qdrant", operation=op, document_id=document_id, collection=self._collection_name)

        except Exception as e:
            self._logger.error("Failed to upload document vector to Qdrant", operation=op, document_id=document_id, collection=self._collection_name, error=str(e))
            raise
