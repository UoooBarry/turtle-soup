FROM golang:1.24-alpine AS base

ARG PORT=8082
ARG GO_ENV=production

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    PORT=$PORT \
    GO_ENV=$GO_ENV

RUN apk update && apk add --no-cache build-base

# Set working directory
WORKDIR /app

# Builder stage
FROM base AS builder

ENV GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir -p ./bin
COPY . .
RUN go build -o ./bin/api ./cmd/api
RUN go build -o ./bin/cli ./cmd/cli

# Dev stage
FROM base AS dev
ENV GO_ENV=development

COPY . .
RUN go mod download
EXPOSE $PORT

# Prod stage
FROM alpine:latest AS prod
WORKDIR /app
VOLUME /app/data

# Copy pre-built files
COPY --from=builder /app/bin/api .
COPY --from=builder /app/bin/cli .
RUN mkdir -p ./data

EXPOSE $PORT

# Run the binary
CMD ["./api"]
