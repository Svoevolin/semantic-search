from structlog import BoundLogger
from uuid import uuid4

from internal.domain.document import (
    DocumentServiceI,
    UploadEvent,
    StorageDownloader,
    TextExtractor,
    Vectorizer,
    QdrantUploader,
)


class DocumentService(DocumentServiceI):
    def __init__(
            self,
            downloader: StorageDownloader,
            extractor: TextExtractor,
            vectorizer: Vectorizer,
            uploader: QdrantUploader,
            logger: BoundLogger,
    ):
        self._downloader = downloader
        self._extractor = extractor
        self._vectorizer = vectorizer
        self._uploader = uploader
        self._logger = logger

    async def process(self, event: UploadEvent) -> None:
        op = "service.document.process"

        self._logger.info("Start processing uploaded document", operation=op, document_id=event.document_id, file_name=event.file_name)

        try:
            raw_file = await self._downloader.download(event.object_url)
            self._logger.info("File downloaded successfully", operation=op, document_id=event.document_id, size=len(raw_file.content))

            text_file = await self._extractor.extract(raw_file)
            self._logger.info("Text extracted successfully", operation=op, document_id=event.document_id, text_size=len(text_file.content))

            embeddings = await self._vectorizer.vectorize_chunks(text_file)
            self._logger.info("Text vectorized successfully", operation=op, document_id=event.document_id, vector_count=len(embeddings))

            for i, embedding in enumerate(embeddings):
                await self._uploader.upload(
                    document_id=str(uuid4()),
                    vector=embedding,
                    payload={
                        "document_id": event.document_id,
                        "file_name": event.file_name,
                        "chunk_index": i,
                        "text": embedding.chunk_text,  # ← кладём оригинальный чанк текста
                    },
                )

            # await self._uploader.upload(
            #     document_id=event.document_id, vector=embedding, payload={"file_name": event.file_name})

            self._logger.info("Vectors uploaded to Qdrant successfully", operation=op, document_id=event.document_id)

        except Exception as e:
            self._logger.error("Failed to process document", operation=op, document_id=event.document_id, error=str(e))
            raise
