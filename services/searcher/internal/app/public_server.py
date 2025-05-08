from fastapi import FastAPI
from contextlib import asynccontextmanager
import structlog

from internal.config.config import load_config
from internal.adapter.vectorizer import SentenceTransformerVectorizer
from internal.adapter.qdrant import QdrantSearcherAdapter
from internal.service.search import SearchService
from internal.delivery.http.handlers.search import router as search_router


@asynccontextmanager
async def lifespan(app: FastAPI):
    config = load_config()
    logger = structlog.get_logger()

    vectorizer = SentenceTransformerVectorizer(config, logger)
    searcher = QdrantSearcherAdapter(config, logger)
    app.state.search_service = SearchService(vectorizer, searcher, logger)

    yield


def create_app() -> FastAPI:
    app = FastAPI(lifespan=lifespan)
    app.include_router(search_router, prefix="/api/v1", tags=["search"])
    return app
