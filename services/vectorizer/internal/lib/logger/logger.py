import logging
import sys

import structlog
from structlog.contextvars import merge_contextvars

from internal.config.config import AppConfig


def setup_logger(config: AppConfig) -> structlog.BoundLogger:
    log_level = getattr(logging, config.logger_level.upper(), logging.INFO)

    timestamper = structlog.processors.TimeStamper(fmt="%Y-%m-%dT%H:%M:%S.%fZ", utc=True, key="time")

    shared_processors = [
        timestamper,
        structlog.stdlib.add_log_level,
        merge_contextvars,
        structlog.processors.StackInfoRenderer(),
        structlog.processors.format_exc_info,
    ]

    if config.logger_pretty_enable:
        renderer = structlog.dev.ConsoleRenderer()
        processors = shared_processors
    else:
        renderer = structlog.processors.JSONRenderer()
        processors = shared_processors + [
            structlog.processors.EventRenamer(to="msg")
        ]

    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=log_level,
    )

    for noisy_logger in ["aiokafka", "kafka", "urllib3", "pdfminer", "huggingface_hub", "transformers", "sentence_transformers", "httpcore", "httpx", "h11"]:
        logging.getLogger(noisy_logger).setLevel(logging.WARNING)

    structlog.configure(
        processors=processors + [
            structlog.stdlib.ProcessorFormatter.wrap_for_formatter,
        ],
        context_class=dict,
        wrapper_class=structlog.make_filtering_bound_logger(log_level),
        logger_factory=structlog.stdlib.LoggerFactory(),
        cache_logger_on_first_use=True,
    )

    handler = logging.StreamHandler(sys.stdout)
    formatter = structlog.stdlib.ProcessorFormatter(
        foreign_pre_chain=processors,
        processors=[
            structlog.stdlib.ProcessorFormatter.remove_processors_meta,
            renderer,
        ],
    )
    handler.setFormatter(formatter)

    root_logger = logging.getLogger()
    root_logger.handlers = []
    root_logger.addHandler(handler)
    root_logger.setLevel(log_level)

    return structlog.get_logger().bind(service="vectorizer")
