from abc import ABC
import structlog
from sentence_transformers import SentenceTransformer

from internal.domain.document import Vectorizer, TextFile, VectorEmbedding


class SentenceTransformerVectorizer(Vectorizer):
    """Адаптер для векторизации текста через SentenceTransformer."""

    def __init__(self, logger: structlog.BoundLogger):
        self._logger = logger
        self._model = SentenceTransformer("sentence-transformers/all-MiniLM-L6-v2")

    async def vectorize(self, text_file: TextFile) -> VectorEmbedding:
        op = "adapter.vectorizer.sentence_transformer.vectorize"

        self._logger.debug("Starting vectorization", operation=op, text_size=len(text_file.content))

        try:
            embedding = self._model.encode(text_file.content).tolist()
            self._logger.debug("Vectorization completed", operation=op, vector_size=len(embedding))
            return VectorEmbedding(values=embedding)

        except Exception as e:
            self._logger.error("Failed to vectorize text", operation=op, error=str(e))
            raise