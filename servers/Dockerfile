# syntax = docker/dockerfile:experimental
FROM golang:latest AS builder
WORKDIR /workspace
COPY . .
RUN make build-all
RUN --mount=type=cache,target=/go/pkg/mod make build-all

FROM gcr.io/distroless/base AS application
COPY --from=builder /workspace/bin/api /usr/local/bin/api
ENTRYPOINT ["/usr/local/bin/api"]
