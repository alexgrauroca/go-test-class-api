ARG GO_VERSION=1.18.1

FROM golang:${GO_VERSION}-alpine AS builder

RUN go install github.com/alexgrauroca/go-test-class-api@latest

ENV APP_HOME /go/src/go-test-class-api
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

CMD [ "go", "run", "main.go" ]
