FROM golang:1.16-alpine

# RUN apk add build-base

WORKDIR /root/website-downloader

COPY ./ ./

RUN go mod download

RUN go build