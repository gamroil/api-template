FROM golang:1.21.0-alpine as builder

WORKDIR /app
COPY . /app

RUN apk add make \
    && make build

ENTRYPOINT ["/app/api", "version"]

FROM alpine

WORKDIR /app
COPY --from=builder /app/api /app

ENTRYPOINT ["/app/api", "serve"]
