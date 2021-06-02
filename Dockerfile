ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.12

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH

LABEL org.opencontainers.image.source="https://github.com/omegion/go-ddclient"

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

WORKDIR /app

COPY ./ /app

RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/ddclient /bin/ddclient

ENTRYPOINT ["ddclient"]
