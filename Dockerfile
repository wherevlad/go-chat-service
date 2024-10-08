FROM golang:1.23.1-alpine AS builder

COPY . /github.com/wherevlad/go-chat-service/source/
WORKDIR /github.com/wherevlad/go-chat-service/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/server/main.go

FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=builder /github.com/wherevlad/go-chat-service/source/bin/crud_server .
COPY --from=builder /github.com/wherevlad/go-chat-service/source/entrypoint.sh .
COPY --from=builder /github.com/wherevlad/go-chat-service/source/migrations ./migrations

ADD https://github.com/pressly/goose/releases/download/v3.22.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose