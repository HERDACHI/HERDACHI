package main

/*
Application that must read a message from the aws queue,
obtain some data from the project database, build a
message with the data obtained and then send it to
another aws queue.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rule_executor_queue/connect_queue"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//structure to get data from the queue
type Domainqueue struct {
	Domain_id int    `json:"domain_id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
}

//structure to send data from the queue
type Message struct {
	Id     int    `json:"id"`
	Domain string `json:"name"`
	Url    string `json:"url"`
}

//Check that the environment variables for the connection to the DB are present
func checkIfEnvVarsPresent() {
	switch "" {
	case os.Getenv("POSTGRES_HOST"):
		fmt.Println("Env var. POSTGRES_HOST is not set. Check your .env.sh file")
	case os.Getenv("POSTGRES_PORT"):
		fmt.Println("Env var. POSTGRES_PORT is not set. Check your .env.sh file")
	case os.Getenv("POSTGRES_USER"):
		fmt.Println("Env var. POSTGRES_USER is not set. Check your .env.sh file")
	case os.Getenv("POSTGRES_PASSWORD"):
		fmt.Println("Env var. POSTGRES_PASSWORD is not set. Check your .env.sh file")
	case os.Getenv("POSTGRES_DATABASE"):
		fmt.Println("Env var. POSTGRES_DATABASE is not set. Check your .env.sh file")
	case os.Getenv("API_TOKEN"):
		fmt.Println("Env var. API_TOKEN is not set. Check your .env.sh file")
	case os.Getenv("API_PORT"):
		fmt.Println("Env var. API_PORT is not set. Check your .env.sh file")
	case os.Getenv("QUEUE_NAME"):
		fmt.Println("Env var. QUEUE_NAME is not set. Check your .env.sh file")
	case os.Getenv("SQS_REGION"):
		fmt.Println("Env var. SQS_REGION is not set. Check your .env.sh file")
	case os.Getenv("URL_RESPONSE"):
		fmt.Println("Env var. URL_RESPONSE is not set. Check your .env.sh file")
	default:
		return
	}
	os.Exit(1)
}

//function to connect via gorm to the DB
func connectToDatabase() (db *gorm.DB, err error) {

	dbCredentials := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	db, err = gorm.Open(postgres.Open(dbCredentials), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open $e", err)
		return db, err
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Connection Failed to the database.", err)
	}

	// Ping
	err = sqlDB.Ping()
	if err != nil {
		log.Println("Connection Failed to the database.", err)
	}

	return db, err
}

//function to read a json file while putting
//the read function from the queue
func readjson() Domainqueue {
	//get the file name
	nameFile := "./json.txt"
	//get the contents of the file
	data, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Println("ioutil.ReadAll error $e", err)

	}

	var domainqueue Domainqueue

	//save file data in a struct
	json.Unmarshal([]byte(data), &domainqueue)

	return domainqueue

}

//function to get the domain name from
//the domains table of the database
func getdomain(iddomain int) string {

	checkIfEnvVarsPresent()

	db, err := connectToDatabase()
	if err != nil {
		log.Println("Error: connect to data base. ", err)
	}

	// Get the execution_id that was just created in the execution table
	var namedomain string
	err = db.Raw("SELECT name FROM domain where id=?", iddomain).Scan(&namedomain).Error
	if err != nil {
		log.Println("Error: Database query error", err)

	}

	return namedomain
}

//Function that builds, with the data obtained,
//the message that will be sent to the aws queue
func constructmessages(id int, domain string) string {

	domain = strings.Replace(domain, "www.", "", -1)

	url := "http://domain-api/get-domain-contents/" + domain
	data := Message{Id: id, Domain: domain, Url: url}

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println("Error creating json. ", err)
	}
	postdomain := string(dataJson)

	return postdomain
}

//function that sends the message to the aws queue
func startSendMessages(queueBody string) {

	queue := os.Getenv("QUEUE_NAME")
	sess, err := connect_queue.SessionQueue()
	if err != nil {
		log.Println("Got an error sending the message, SessionQueue function has failed:", err)
		return
	}

	// Get URL of queue
	urlResult, err := connect_queue.GetQueueURL(sess, &queue)
	if err != nil {
		log.Println("Got an error sending the message, GetQueueURL function has failed:", err)
		return
	}
	queueURL := urlResult.QueueUrl

	err = connect_queue.SendMessage(sess, queueURL, queueBody)
	if err != nil {
		log.Println("Got an error sending the message, SendMessage function has failed:", err)
		return
	}

}

func allservice(w http.ResponseWriter, r *http.Request) {

	domainqueue := readjson()
	iddomain := domainqueue.Domain_id
	namedomain := getdomain(iddomain)
	queueBody := constructmessages(iddomain, namedomain)
	startSendMessages(queueBody)
	fmt.Println(queueBody)
	http.Error(w, queueBody, 400)
}

func main() {
	port := os.Getenv("API_PORT")
	//allservice()
	http.HandleFunc("/", allservice) //.Methods("GET")
	log.Println(http.ListenAndServe(":"+port, nil))

}
