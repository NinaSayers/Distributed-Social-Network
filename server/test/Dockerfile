FROM golang:1.23-alpine AS builder

RUN apk add gcc musl-dev sqlite-dev

ENV CGO_ENABLED=1
# RUN apk add --no-cache gcc musl-dev libc6-compat
# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build -o app test/cmd/main.go
