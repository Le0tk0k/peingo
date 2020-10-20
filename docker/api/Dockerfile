FROM golang:1.15 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src/github.com/Le0tk0k/peingo

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 1323
CMD ["go", "run", "main.go"]