# Build the manager binary
FROM docker.io/dongjiang1989/golang:1.22.1-alpine as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

#RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY internal/ internal/
COPY docs/ docs/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} make build-binary


FROM alpine:3.20.3

WORKDIR /
COPY --from=builder /workspace/echo .

ENTRYPOINT ["/echo"]
