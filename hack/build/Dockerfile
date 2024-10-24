# Build the manager binary
FROM golang:1.23 AS builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY VERSION VERSION
COPY Makefile Makefile

#RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY .git/ .git/
COPY internal/ internal/
COPY docs/ docs/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} make build-binary


FROM alpine:3.20.3

WORKDIR /
COPY --from=builder /workspace/echo .

ENTRYPOINT ["/echo"]
