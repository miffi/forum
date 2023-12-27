FROM golang:alpine AS go-build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN go build -C cmd/server -o /server

ENV GIN_MODE=release
CMD ["/server"]
