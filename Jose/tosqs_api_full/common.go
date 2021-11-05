package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func getFilesToFolder(directory string) []fs.FileInfo {

	//read files from a folder
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println("Error: Read files from a folder. ", err)
	}
	return files
}

/*
	Get query from ./queries/1.json -> Query
*/
func GetQueryList() ([]Query, error) {

	dirQueries := "./queries"
	queries := []Query{}

	filesQueries := getFilesToFolder(dirQueries)

	for _, f := range filesQueries {
		queryId := f.Name()

		data, err := os.Open(fmt.Sprintf("%s", dirQueries+"/"+queryId))
		if err != nil {
			log.Println("Files reading error", err)
			return queries, err
		}
		//get the contents of the file
		byteValue, err := ioutil.ReadAll(data)
		if err != nil {
			log.Println("ioutil.ReadAllerror", err)
			return queries, err
		}
		var obj Query
		//save file data in a struct
		json.Unmarshal([]byte(byteValue), &obj)
		queries = append(queries, obj)
	}

	return queries, nil
}

/*
	Get query from ./queues/1.json -> Query
*/
func GetQueueList() ([]Queue, error) {

	dirQueues := "./queues"
	queues := []Queue{}

	filesQueues := getFilesToFolder(dirQueues)

	for _, f := range filesQueues {
		queryId := f.Name()

		data, err := os.Open(fmt.Sprintf("%s", dirQueues+"/"+queryId))
		if err != nil {
			log.Println("Files reading error", err)
			return queues, err
		}
		//get the contents of the file
		byteValue, err := ioutil.ReadAll(data)
		if err != nil {
			log.Println("ioutil.ReadAllerror", err)
			return queues, err
		}
		var obj Queue
		//save file data in a struct
		json.Unmarshal([]byte(byteValue), &obj)
		queues = append(queues, obj)
	}

	return queues, nil
}
