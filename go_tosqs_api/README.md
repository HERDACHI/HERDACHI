# API TO SQS

The program is an API that takes a domain name from a url parameter,
saves it in a Postgres database, and sends it to an AWS queue.

## building

go mod init go_tosqs_api
go mod tidy
go build

***Running***

### Windows

```
run.bat
```
### Linux or OSX

```
./run.sh
```


# to get dependencies

go get github.com/gorilla/mux


# Links

- https://golang.org/doc/tutorial/compile-install
- https://gorm.io/es_ES/docs/index.html

