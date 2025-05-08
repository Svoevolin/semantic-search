import asyncio, structlog

from sentence_transformers import SentenceTransformer

from internal.config.config import AppConfig
from internal.domain.search import Vectorizer, VectorEmbedding


class SentenceTransformerVectorizer(Vectorizer):
    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._logger = logger
        self._model_name = config.embedding_model_name
        self._model = SentenceTransformer(self._model_name)

    async def vectorize(self, text: str) -> VectorEmbedding:
        op = "adapter.vectorizer.sentence_transformer.vectorize"

        self._logger.debug("Starting vectorization", operation=op, text_size=len(text))

        try:
            embedding = await asyncio.get_running_loop().run_in_executor(None, self._model.encode, text)
            self._logger.debug("Vectorization completed", operation=op, vector_size=len(embedding))
            return VectorEmbedding(embedding.tolist())
        except Exception as e:
            self._logger.error("Failed to vectorize text", operation=op, error=str(e))
            raise
