from typing import List
from pydantic import BaseModel, Field, conint

from internal.domain.search import Paginator, SearchQueryInput, SearchList


class SearchRequest(BaseModel):
    """DTO запроса на поиск."""
    query: str = Field(..., max_length=384, description="Поисковый запрос")
    page: conint(ge=1) = Field(default=1, alias="_page")
    pageSize: conint(ge=1, le=100) = Field(default=10, alias="_pagesize")

    def model(self) -> SearchQueryInput:
        offset = (self.page - 1) * self.pageSize
        return SearchQueryInput(
            query_text=self.query,
            paginator=Paginator(limit=self.pageSize, offset=offset)
        )


class SearchItemResponse(BaseModel):
    document_id: str
    score: float
    snippet: str


class SearchListResponse(BaseModel):
    results: List[SearchItemResponse]

    @classmethod
    def from_model(cls, result: SearchList) -> tuple["SearchListResponse", dict]:
        return cls(
            results=[
                SearchItemResponse(
                    document_id=item.document_id,
                    score=item.score,
                    snippet=item.snippet,
                )
                for item in result.items
            ]
        ), {"X-Has-More": str(result.has_more).lower()}