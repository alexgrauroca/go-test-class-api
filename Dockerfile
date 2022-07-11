ARG GO_VERSION=1.18.1

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go install ./...
RUN go build -o go-test-class-api main.go

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=builder /src/go-test-class-api .
COPY --from=builder /src/.env .
