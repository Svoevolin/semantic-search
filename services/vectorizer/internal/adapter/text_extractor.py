import structlog
from io import BytesIO
from pdfminer.high_level import extract_text

from internal.domain.document import TextExtractor, RawFile, TextFile

class PdfTextExtractor(TextExtractor):
    """Экстрактор текста из PDF файлов."""

    def __init__(self, logger: structlog.BoundLogger):
        self._logger = logger

    async def extract(self, raw_file: RawFile) -> TextFile:
        op = "adapter.text_extractor.pdf.extract"

        self._logger.debug("Starting text extraction from PDF", operation=op, file_size=len(raw_file.content))

        try:
            with BytesIO(raw_file.content) as file_stream:
                extracted_text = extract_text(file_stream)

            self._logger.info("Text extraction completed", operation=op, extracted_length=len(extracted_text))

            return TextFile(content=extracted_text)

        except Exception as e:
            self._logger.error("Failed to extract text from file", operation=op, error=str(e))
            raise
