FROM golang:1.20 as builder
WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download all

COPY ./ ./

ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-X main.Version=${VERSION}" \
    -o go-testcoverage .

FROM gcr.io/distroless/base:latest
WORKDIR /
COPY --from=builder /workspace/go-testcoverage .
COPY --from=builder /usr/local/go/bin/go /usr/local/go/bin/go
ENV PATH="${PATH}:/usr/local/go/bin"
ENTRYPOINT ["/go-testcoverage"]