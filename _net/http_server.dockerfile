FROM golang:1.17-alpine AS builder
WORKDIR /app
ADD _net/http_server.go /app
RUN cd /app && go build http_server.go

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/http_server /app/
EXPOSE 3000
ENTRYPOINT ./http_server

# docker build -t http_server:1.0 -f _net/http_server.dockerfile .
# docker run --rm -p 3000:3000 http_server:1.0
