from pydantic_settings import BaseSettings
from pydantic import Field
from pathlib import Path
from dotenv import load_dotenv

class AppConfig(BaseSettings):
    # Logger
    logger_level: str = Field(..., alias="VECTOR_LOG_LEVEL")
    logger_format: str = Field(..., alias="VECTOR_LOG_FORMAT")
    logger_pretty_enable: bool = Field(..., alias="VECTOR_LOG_PRETTY_ENABLE")

    # Vectorizer
    embedding_model_name: str = Field(..., alias="EMBEDDING_MODEL")
    score_threshold: float = Field(..., alias="SCORE_THRESHOLD")

    # Qdrant
    qdrant_host: str = Field(..., alias="QDRANT_HOST")
    qdrant_port: int = Field(..., alias="QDRANT_PORT")
    qdrant_collection: str = Field(..., alias="QDRANT_COLLECTION")

    class Config:
        env_file = "../../.env"
        populate_by_name = True
        extra = "ignore"

def load_config(env_path: str = "../../.env") -> AppConfig:
    env_file = Path(env_path).resolve()
    load_dotenv(dotenv_path=env_file)
    return AppConfig()  # type: ignore
