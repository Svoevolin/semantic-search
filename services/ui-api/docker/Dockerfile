FROM golang:1.23.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ui-api ./cmd/ui-api

FROM alpine:3.19

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/ui-api /ui-api
COPY --from=builder /app/api/swagger.json /api/swagger.json

ENTRYPOINT ["/ui-api"]
