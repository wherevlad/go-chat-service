FROM golang:1.23.1-alpine AS builder

COPY . /github.com/wherevlad/go-chat-service/source/
WORKDIR /github.com/wherevlad/go-chat-service/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/wherevlad/go-chat-service/source/bin/crud_server .

CMD ["./crud_server"]