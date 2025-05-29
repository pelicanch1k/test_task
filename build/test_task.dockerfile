FROM golang:1.24.3  AS builder

ENV CGO_ENABLED=0

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN  \
  --mount=type=cache,target=/root/.cache \
  --mount=type=cache,target=/root/go/pkg/mod \
  go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN \
  --mount=type=cache,target=/root/.cache \
  --mount=type=cache,target=/root/go/pkg/mod \
  go build -ldflags "-w -s" -gcflags "-B" ./cmd/test_task/

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY --from=builder /app/test_task /app/app
ENTRYPOINT ["./app"]
