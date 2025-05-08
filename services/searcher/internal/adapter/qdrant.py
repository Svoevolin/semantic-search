import structlog

from qdrant_client import QdrantClient

from internal.config.config import AppConfig
from internal.domain.search import QdrantSearcher, SearchQuery, SearchList, SearchItem


class QdrantSearcherAdapter(QdrantSearcher):
    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._logger = logger
        self._collection_name = config.qdrant_collection
        self._score_threshold = config.score_threshold
        self._client = QdrantClient(host=config.qdrant_host, port=config.qdrant_port)

    async def search(self, query: SearchQuery) -> SearchList:
        op = "adapter.qdrant_searcher.search"

        self._logger.debug("Starting search", operation=op, collection=self._collection_name, offset=query.paginator.offset, limit=query.paginator.limit, score_threshold=self._score_threshold)

        try:
            raw_result = self._client.search(
                collection_name=self._collection_name,
                query_vector=query.query[0],
                limit=query.paginator.limit + 1,
                offset=query.paginator.offset,
                with_payload=True,
                score_threshold=self._score_threshold,
            )

            has_more = len(raw_result) > query.paginator.limit

            items = []
            for hit in raw_result[:query.paginator.limit]:
                payload = hit.payload or {}

                # fallback значения
                file_name = payload.get("file_name", "")
                snippet = payload.get("text", "")
                uploaded_at_str = payload.get("uploaded_at", "1970-01-01T00:00:00Z")

                try:
                    uploaded_at = datetime.fromisoformat(uploaded_at_str.replace("Z", "+00:00"))
                except Exception:
                    uploaded_at = datetime.utcfromtimestamp(0)

                items.append(
                    SearchItem(
                        document_id=hit.id,
                        score=hit.score,
                        snippet=snippet,
                        file_name=file_name,
                        uploaded_at=uploaded_at,
                    )
                )

            # items = [SearchItem(document_id=hit.id, score=hit.score, snippet=hit.payload["text"]) for hit in raw_result[:query.paginator.limit]]

            self._logger.info("Search finished", operation=op, returned=len(items), has_more=has_more)

            return SearchList(items=items, has_more=has_more)

        except Exception as e:
            self._logger.error("Qdrant search failed", operation=op, error=str(e))
            raise

