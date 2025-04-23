FROM golang:1.23.1

RUN apt-get update && apt-get install -y git && \
    go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
ENTRYPOINT ["swag"]
