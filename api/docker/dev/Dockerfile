FROM golang:1.20

ENV GO111MODULE=on


COPY . /go/src/app

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

RUN ls

RUN go mod download
