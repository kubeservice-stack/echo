# Build the manager binary
FROM docker.io/dongjiang1989/golang:1.22.1-alpine as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

#RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY docs/ docs/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags "-extldflags -static" -o server ./cmd/main.go


FROM docker.io/docker-repo/library/alpine:latest

WORKDIR /
COPY --from=builder /workspace/server .

ENTRYPOINT ["/server"]
