FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /root/sqs-api

# Download Go modules
COPY go.mod .

RUN go mod download

COPY ./ ./

RUN go get sqs-api

RUN go build