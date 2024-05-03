FROM golang:1.21.0-alpine as builder

WORKDIR /app
COPY . /app

RUN apk add make \
    && make build

ENTRYPOINT ["/app/api-template", "version"]

FROM alpine

WORKDIR /app
COPY --from=builder /app/api-template /app
COPY --from=builder /app/config.yaml /app

ENTRYPOINT ["/app/api-template", "serve"]
