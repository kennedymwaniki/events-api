FROM  go:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download