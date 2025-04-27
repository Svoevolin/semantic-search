import os
import urllib.parse

from minio import Minio
from minio.error import S3Error

import structlog

from internal.domain.document import StorageDownloader, RawFile
from internal.config.config import AppConfig


class MinioClient(StorageDownloader):
    def __init__(self, config: AppConfig, logger: structlog.BoundLogger):
        self._client = Minio(
            endpoint=config.minio_endpoint,
            access_key=config.minio_access_key,
            secret_key=config.minio_secret_key,
            secure=config.minio_use_ssl,
        )
        self._bucket = config.minio_bucket
        self._logger = logger

    async def download(self, object_url: str) -> RawFile:
        op = "minio.client.download"

        self._logger.debug("Starting file download", operation=op, bucket=self._bucket, object_url=object_url)

        try:
            parsed_url = urllib.parse.urlparse(object_url)
            object_name = urllib.parse.unquote(os.path.basename(parsed_url.path))

            response = self._client.get_object(self._bucket, object_name)
            content = response.read()
            response.close()
            response.release_conn()

            self._logger.info("File downloaded successfully", operation=op, bucket=self._bucket, object_url=object_url)

            return RawFile(content=content)

        except S3Error as e:
            if e.code == "NoSuchKey":
                self._logger.error("File not found in MinIO", operation=op, object_url=object_url, error=str(e))
            else:
                self._logger.error("Failed to download file", operation=op, object_url=object_url, error=str(e))
            raise
        except Exception as e:
            self._logger.exception("Failed to download file",  operation=op, object_url=object_url, error=str(e))
            raise