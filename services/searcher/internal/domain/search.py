from abc import ABC, abstractmethod
from pydantic import BaseModel, Field
from datetime import datetime


VectorEmbedding = list[float]


class Paginator(BaseModel):
    """Параметры пагинации."""
    limit: int = Field(..., ge=1)
    offset: int = Field(..., ge=0)


class SearchItem(BaseModel):
    """Один результат поиска."""
    document_id: str
    file_name: str
    score: float = Field(..., alias="match_score")
    snippet: str
    uploaded_at: datetime


class SearchList(BaseModel):
    """Результат поиска."""
    items: list[SearchItem]
    has_more: bool


class SearchQueryInput(BaseModel):
    """Входной запрос от пользователя."""
    query_text: str
    paginator: Paginator


class SearchQuery(BaseModel):
    """Подготовленный запрос с векторами."""
    query: list[VectorEmbedding] = Field(..., min_length=1)
    paginator: Paginator


class Vectorizer(ABC):
    @abstractmethod
    async def vectorize(self, text: str) -> VectorEmbedding:
        pass


class QdrantSearcher(ABC):
    @abstractmethod
    async def search(self, query: SearchQuery) -> SearchList:
        pass


class SearchServiceI(ABC):
    @abstractmethod
    async def search(self, query_input: SearchQueryInput) -> SearchList:
        pass
