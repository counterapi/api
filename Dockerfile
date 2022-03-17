ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.13.7

FROM golang:${GO_VERSION} AS builder

LABEL org.opencontainers.image.source = "https://github.com/counterapi/counterapi"

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY ./ /app

WORKDIR /app

RUN make build-for-container

FROM ${FROM_IMAGE}

ENV GIN_MODE release

RUN apk update && \
  apk add ca-certificates openssl && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY --from=builder /app/dist/counter-linux /bin/counter

EXPOSE 80

ENTRYPOINT ["counter"]
