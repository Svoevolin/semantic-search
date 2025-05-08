from structlog import BoundLogger

from internal.domain.search import (
    SearchServiceI,
    SearchQueryInput,
    SearchQuery,
    Vectorizer,
    QdrantSearcher,
    SearchList,
)


class SearchService(SearchServiceI):
    def __init__(self, vectorizer: Vectorizer, searcher: QdrantSearcher, logger: BoundLogger):
        self._vectorizer = vectorizer
        self._searcher = searcher
        self._logger = logger

    async def search(self, query_input: SearchQueryInput) -> SearchList:
        op = "service.search.search"

        self._logger.info("Start search", operation=op, query_text=query_input.query_text, offset=query_input.paginator.offset, limit=query_input.paginator.limit)

        try:
            query_vector = await self._vectorizer.vectorize(query_input.query_text)
            self._logger.debug("Query vectorized", operation=op, vector_size=len(query_vector))

            query = SearchQuery(
                query=[query_vector],
                paginator=query_input.paginator,
            )

            result = await self._searcher.search(query)

            self._logger.info("Search completed", operation=op, returned=len(result.items), has_more=result.has_more)

            return result

        except Exception as e:
            self._logger.error("Search failed", operation=op, error=str(e))
            raise
