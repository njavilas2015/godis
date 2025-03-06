FROM golang:1.23.4-alpine3.20

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update --no-cache && apk add --no-cache bash git gcc musl-dev curl

WORKDIR /usr/src/app

RUN go install golang.org/x/tools/gopls@latest

RUN curl -sSfL https://github.com/golangci/golangci-lint/releases/download/v1.54.0/golangci-lint-1.54.0-linux-amd64.tar.gz \
    | tar -xzC /usr/local/bin

EXPOSE 6480 6481 6482 6483