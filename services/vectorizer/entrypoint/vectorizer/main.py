import asyncio

from internal.adapter.qdrant import QdrantUploaderAdapter
from internal.adapter.vectorizer import SentenceTransformerVectorizer
from internal.config.config import load_config
from internal.lib.logger.logger import setup_logger
from internal.adapter.kafka import KafkaConsumerAdapter
from internal.adapter.minio2 import MinioClient
from internal.adapter.text_extractor import PdfTextExtractor
from internal.app.runner import KafkaConsumerRunner
from internal.service.document import DocumentService

async def main():
    config = load_config()

    logger = setup_logger(config)
    logger.info("Vectorizer service started")

    kafka_consumer = KafkaConsumerAdapter(config, logger)
    storage_downloader = MinioClient(config, logger)
    extractor = PdfTextExtractor(logger)
    vectorizer = SentenceTransformerVectorizer(logger)
    qdrant_uploader = QdrantUploaderAdapter(config, logger)
    await qdrant_uploader.initialize()

    document_service = DocumentService(
        downloader=storage_downloader,
        extractor=extractor,
        vectorizer=vectorizer,
        uploader=qdrant_uploader,
        logger=logger,
    )

    runner = KafkaConsumerRunner(
        consumer=kafka_consumer,
        service=document_service,
        logger=logger,
    )

    await runner.run()

if __name__ == "__main__":
    asyncio.run(main())