ARG GO_VERSION=1.18-alpine3.15
ARG FROM_IMAGE=alpine:3.15

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH
ARG VERSION

LABEL org.opencontainers.image.source = "https://github.com/counterapi/api"

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY ./ /app

WORKDIR /app

RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH VERSION=$VERSION

FROM ${FROM_IMAGE}

ENV GIN_MODE release

RUN apk update && \
  apk add ca-certificates openssl && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY --from=builder /app/dist/counterapi /bin/counterapi

EXPOSE 80

ENTRYPOINT ["counter"]
