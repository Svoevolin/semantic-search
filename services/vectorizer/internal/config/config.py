from pydantic_settings import BaseSettings
from pydantic import Field
from pathlib import Path
from dotenv import load_dotenv


class AppConfig(BaseSettings):
    # Logger
    logger_level: str = Field(..., alias="VECTOR_LOG_LEVEL")
    logger_format: str = Field(..., alias="VECTOR_LOG_FORMAT")
    logger_pretty_enable: bool = Field(..., alias="VECTOR_LOG_PRETTY_ENABLE")

    # Minio
    minio_endpoint: str = Field(..., alias="MINIO_ENDPOINT")
    minio_access_key: str = Field(..., alias="MINIO_ACCESS_KEY")
    minio_secret_key: str = Field(..., alias="MINIO_SECRET_KEY")
    minio_bucket: str = Field(..., alias="MINIO_BUCKET")
    minio_use_ssl: bool = Field(False, alias="MINIO_USE_SSL")

    # Qdrant
    qdrant_host: str = Field(..., alias="QDRANT_HOST")
    qdrant_port: int = Field(..., alias="QDRANT_PORT")
    qdrant_collection: str = Field(..., alias="QDRANT_COLLECTION")

    qdrant_embedding_size: int = Field(..., alias="VECTOR_EMBEDDING_SIZE")

    # Kafka
    kafka_broker: str = Field(..., alias="KAFKA_BROKER")
    kafka_topic: str = Field(..., alias="KAFKA_TOPIC")

    class Config:
        env_file = "../../.env"
        populate_by_name = True
        extra = "ignore"

def load_config(env_path: str = "../../.env") -> AppConfig:
    env_file = Path(env_path).resolve()
    load_dotenv(dotenv_path=env_file)
    return AppConfig()  # type: ignore
