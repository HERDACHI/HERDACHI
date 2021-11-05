# API FEATURE

The program with the name of a domain provided as a parameter through a url, obtains the records from the features table with respect to that id, then returns these records to a json and prints them.

## building

go mod init api_feature
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

