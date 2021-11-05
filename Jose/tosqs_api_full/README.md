# API TO SQS FULL

The program is an API that takes a  name or id from a url parameter,
saves it in a Postgres database , and sends it to an AWS queue.
The program is an API that has 5 main functions:

GetFileQuery:Get the query of a sql file from the id_query parameter assigned in the URL.

GetFileQueryCount: Obtains the number of records of the query performed by the GetFileQuery function

GetAllDomains: Gets all the records in the domain table of the database.

CreateDomain: Create a new record in the domain table from the "domain" parameter indicated in the endpoint.

CreateDomainIfnotExist: Create a new record in the domain table if it does not exist from the "domain" parameter indicated in the endpoint.

startSendMessages: Send a message with the necessary data to an Amazon SQS queue

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

