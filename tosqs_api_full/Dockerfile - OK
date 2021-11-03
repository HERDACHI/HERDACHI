FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /root/go_tosqs_api

# Download Go modules
COPY go.mod .

RUN go mod download
CMD
COPY . .

RUN go get api_tosqs_go

RUN go build

#Start aplication
#CMD ["/root/go_tosqs_api/", "run"]
