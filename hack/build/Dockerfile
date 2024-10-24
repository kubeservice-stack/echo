# Build the manager binary
FROM golang:1.23 AS builder

WORKDIR /workspace

# Copy source files
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} make build-binary


FROM alpine:3.20.3

WORKDIR /
COPY --from=builder /workspace/echo .

ENTRYPOINT ["/echo"]
