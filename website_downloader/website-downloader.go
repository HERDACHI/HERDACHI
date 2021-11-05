/**********************CODE IN TRIAL PERIOD, NOT FINISHED.!**************************/
package main

/*
 Application that reads a record from SQS DOWNLOAD_QUEUE, and
 It has as input a json with the url and the id of the domain table,
 then it obtains the bytes of that URL and saves them in the feature table.
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"website_downloader/connect_queue"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Structure with the expected json of table domain
type Domain struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type dataDomain struct {
	ID   int    `json:"Id_psql"`
	Body string `json:"Body"`
}

//Structure with the expected json table execution
type Execution struct {
	Created time.Time `json:"created" gorm:"created"`
}

//Structure with the expected json of table feature
type Feature struct {
	Name          string    `json:"name"`
	Value         string    `json:"value"`
	LastCrawlDate time.Time `json:"last_crawl_date" gorm:"last_crawl_date"`
	Domain_id     int       `json:"domain_id"`
	Execution_id  int       `json:"execution_id"`
}

// structure of the json that is sent to the queue
type Response struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Entrance struct {
	Idpsql int    `json:"id"_psql`
	Body   string `json:"body"`
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

// function that gets all the bytes of the url address received
func getdomain(name string) string {
	response, err := http.Get(name)
	if err != nil {
		log.Println("Error: Could not extract bytes from url. ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading file: ", err)
	}

	totalbytes := string([]byte(body))

	return totalbytes
}

// Function that obtains from the json received the values ​​of the domain and its id
func readjson(jsonvalue string) (int, string) {

	var domain Domain

	json.Unmarshal([]byte(jsonvalue), &domain)

	return domain.Id, domain.Name

}

// Save the bytes of the url and its id in the table feature
func createfeature(id int, fullbytes string, namedomain string) string {

	checkIfEnvVarsPresent()

	db, err := connectToDatabase()
	if err != nil {
		log.Println("Error: connect to data base. ", err)
	}

	result := db.Table("execution").Create(&Execution{Created: time.Now()})
	if result.Error != nil {
		log.Println("Error: Creating record in execution table", result.Error)
	}

	// Get the execution_id that was just created in the execution table
	var newid int
	db.Raw("SELECT id FROM execution order by id desc limit 1").Scan(&newid)

	data := Feature{Name: "website_contents", Value: fullbytes, LastCrawlDate: time.Now(), Domain_id: id, Execution_id: newid}

	result = db.Table("feature").Create(&data)
	if result.Error != nil {
		log.Println("Error: Creating record in feature table", result.Error)
	} else {
		log.Println("Record successfully saved in feature table with bytes of ", namedomain)
	}

	return fullbytes
}

// function that builds the json that will be sent to the queue
func jsontoqueue(id int, name string, fullbytes string) string {

	data := Response{Id: id, Name: name, Content: fullbytes}

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println("Error creating json. ", err)
	}
	postdomain := string(dataJson)

	return postdomain

}

func jsongetqueue(id int, name string) string {

	data := Entrance{Idpsql: id, Body: name}

	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Println("Error creating json. ", err)
	}
	postdomain := string(dataJson)

	return postdomain

}

func startSendMessages(queueBody string) {

	queue := os.Getenv("OUT_QUEUE_NAME")
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

func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		MessageAttributeNames: []*string{aws.String("All")},
		QueueUrl:              queueURL,
		MaxNumberOfMessages:   aws.Int64(1),
		VisibilityTimeout:     timeout,
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func receiveMessage(InQueueName string) string {

	//timeout: 5s
	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Get URL of queue connect_queue.SessionQueue()
	urlResult, err := connect_queue.GetQueueURL(sess, &InQueueName)
	if err != nil {
		log.Println("Got an error sending the message, GetQueueURL function has failed.:", err)

	}
	queueURL := urlResult.QueueUrl

	fmt.Println("URL: " + *urlResult.QueueUrl)
	if err != nil {
		log.Println("Got an error getting the queue URL:", err)

	}

	//Retrieve SQS message
	msgResult, err := GetMessages(sess, queueURL, timeout)

	if err != nil {

		log.Println("Got an error receiving messages:", err)

	}

	//Print SQS message
	fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
	fmt.Println("Domain:     " + *msgResult.Messages[0].Body)
	fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	jsonBody := *msgResult.Messages[1].Body
	return jsonBody
}

func main() {

	// this argument will come as input
	InQueueName := os.Getenv("QUEUE_NAME")
	// structure example -> `{"id": 8,"name": "http://example.com"}`
	jsonvalue := receiveMessage(InQueueName)
	fmt.Println("JSON:", jsonvalue)

	// The id and name of the domain is obtained from the json input
	iddomain, namedomain := readjson(jsonvalue)

	// All bytes of the domain are obtained
	bodybytes := getdomain(namedomain) //bytes

	// The information obtained is saved in the feature table
	fullbytes := createfeature(iddomain, bodybytes, namedomain)

	json := jsontoqueue(iddomain, namedomain, fullbytes)
	fmt.Println(json)

	startSendMessages(json)

}
