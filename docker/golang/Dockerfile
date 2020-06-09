FROM golang:1.14.3

RUN apt-get update -qq

WORKDIR /todo

ADD . .

RUN go mod download
RUN go build

RUN go get github.com/rubenv/sql-migrate/...
