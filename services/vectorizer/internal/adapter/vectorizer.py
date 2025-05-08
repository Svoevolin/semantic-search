from abc import ABC
import structlog
from sentence_transformers import SentenceTransformer

from internal.config.config import AppConfig
from internal.domain.document import Vectorizer, TextFile, VectorEmbedding


class SentenceTransformerVectorizer(Vectorizer):
    """Адаптер для векторизации текста через SentenceTransformer."""

    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._logger = logger
        self._model = SentenceTransformer(config.embedding_model_name)
        self._tokenizer = self._model.tokenizer
        self._max_tokens = self._model.get_max_seq_length()

    async def vectorize_chunks(self, text_file: TextFile) -> list[VectorEmbedding]:
        op = "adapter.vectorizer.sentence_transformer.vectorize_chunks"

        chunks = self._split_text(text_file.content)
        self._logger.debug("Text split into chunks", operation=op, chunk_count=len(chunks))

        try:
            vectors = self._model.encode(chunks, batch_size=8).tolist()
            return [
                VectorEmbedding(values=v, chunk_text=t)
                for v, t in zip(vectors, chunks)
            ]
        except Exception as e:
            self._logger.error("Chunk vectorization failed", operation=op, error=str(e))
            raise

    def _split_text(self, text: str, stride: int = 50) -> list[str]:
        tokens = self._tokenizer(
            text,
            return_overflowing_tokens=True,
            return_tensors=None,
            truncation=True,
            max_length=self._max_tokens,
            stride=stride,
        )
        return self._tokenizer.batch_decode(tokens["input_ids"], skip_special_tokens=True)
