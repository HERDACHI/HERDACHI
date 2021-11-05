package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Query struct {
	ID             uint   `json:"id" `
	Description    string `json:"description" `
	Query          string `json:"query" `
	CountRows      int    `json:"countRows" `
	QueryCountRows string `json:"countRowsQuery" `
}

func main() {

	// Directory we want to get all files from.
	directory := "./queries"

	//read files from a folder
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("Directory reading error $e", err)
		return
	}

	//Iterate each of the files
	for _, file := range files {

		//get the file name
		nameFile := file.Name()
		//open file
		data, err := os.Open(directory + "/" + nameFile)
		if err != nil {
			log.Println("File reading error $e", err)
			return
		}

		//get the contents of the file
		byteValue, err := ioutil.ReadAll(data)
		if err != nil {
			log.Println("ioutil.ReadAll error $e", err)
			return
		}

		var query Query
		var queries []Query

		//save file data in a struct
		json.Unmarshal([]byte(byteValue), &query)

		//save all data in a array
		queries = append(queries, query)

		//print the array data
		fmt.Printf("%+v\n", queries)
	}
}

///File example
//nameFile = 1.json
//Content file
/*
{
	"id": 1,
	"description": "get 10 first records",
	"query": "SELECT * FROM public.domain ORDER BY last_crawl_date ASC LIMIT 10;",
	"countRows":10
}*/
