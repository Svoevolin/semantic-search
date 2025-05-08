from abc import ABC, abstractmethod
from typing import Any
from pydantic import BaseModel, Field

# ====== Доменные сущности ======

class UploadEvent(BaseModel):
    """Событие загрузки документа."""

    document_id: str = Field(..., alias="DocumentID")
    file_name: str = Field(..., alias="FileName")
    object_url: str = Field(..., alias="ObjectURL")
    request_id: str = Field(..., alias="RequestID")


class RawFile(BaseModel):
    """Сырые данные файла."""

    content: bytes


class TextFile(BaseModel):
    """Результат извлечения текста из документа."""

    content: str


class VectorEmbedding(BaseModel):
    """Эмбеддинг вектора."""

    values: list[float]
    chunk_text: str


# ====== Интерфейсы портов ======

class DocumentServiceI(ABC):
    """Бизнес-логика обработки загруженного документа: скачал -> векторизовал -> загрузил в Qdrant."""

    @abstractmethod
    async def process(self, event: UploadEvent) -> None:
        pass


class StorageDownloader(ABC):
    """Скачивание файлов из хранилища."""

    @abstractmethod
    async def download(self, url: str) -> RawFile:
        pass

# internal/domain/document.py

class TextExtractor(ABC):
    """Интерфейс для извлечения текста из сырых файлов."""

    @abstractmethod
    async def extract(self, raw_file: RawFile) -> TextFile:
        pass


class Vectorizer(ABC):
    """Векторизация текста."""

    @abstractmethod
    async def vectorize_chunks(self, text: TextFile) -> list[VectorEmbedding]:
        pass


class QdrantUploader(ABC):
    """Загрузка эмбеддингов в базу векторов."""

    @abstractmethod
    async def upload(self, document_id: str, vector: VectorEmbedding, payload: dict[str, Any]) -> None:
        pass
