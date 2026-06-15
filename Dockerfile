FROM golang:1.22-alpine AS builder

WORKDIR /src
COPY . .

ARG APP_VERSION=dev
ARG APP_COMMIT=unknown
ARG APP_BRANCH=local
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build \
    -trimpath \
    -ldflags "-s -w -X main.version=${APP_VERSION} -X main.commit=${APP_COMMIT} -X main.branch=${APP_BRANCH}" \
    -o /out/server ./cmd/server

FROM alpine:3.20

RUN adduser -D -H -u 10001 appuser
WORKDIR /app
COPY --from=builder /out/server /app/server

ENV PORT=8080
ENV APP_NAME=hello-health-service
EXPOSE 8080

USER appuser
ENTRYPOINT ["/app/server"]
