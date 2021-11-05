//The program is an API  whit all the necessary functions  to the API TOSQS UI FORM.

package main

import (
	"api_tosqs_go/connect_queue"
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"log"
	"net/http"
	"net/textproto"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
	Check if all required environment variables are present.
*/
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
		fmt.Println("Env var. API_PORT is not set. Check your .env.sh file")
	default:
		return
	}
	os.Exit(1)
}

type Domain struct {
	Id            uint      `json:"id" gorm:"id"`
	Name          string    `json:"name" gorm:"name"`
	LastCrawlDate time.Time `json:"last_crawl_date" gorm:"last_crawl_date"`
}

type Item struct {
	Name string `json:"name" gorm:"name"`
}

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}

type dataDomain struct {
	ID   int    `json:"Id_psql"`
	Body string `json:"Body"`
}

func (Domain) TableName() string {
	return "domain"
}

func connectToDatabase() (db *gorm.DB, err error) {

	dbCreds := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	db, err = gorm.Open(postgres.Open(dbCreds), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open $e", err)
		return db, err
	}

	return db, err

}

// get all domain in postgres db
func GetDomain(w http.ResponseWriter, r *http.Request) {
	//This is name endpoint
	log.Println("Endpoint Hit: GetDomain")

	//allDomain: It is the variable that will use all the domains
	allDomain := []Domain{}

	db, err := connectToDatabase()
	if err != nil {
		log.Println("Error: connect to data base ", err)
		http.Error(w, err.Error(), 500)
	}
	// SELECT * FROM `domain`
	err = db.Find(&allDomain).Error
	if err != nil {
		log.Println("GetDomain $e", err)
		log.Println("ERROR: $e", err)
		http.Error(w, err.Error(), 500)
	}

	//print json with all domains
	json.NewEncoder(w).Encode(allDomain)
}

/* Insert domain names in postgres db from URL*/

// Create Function
func CreateDomain(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDatabase()

	if err != nil {
		log.Println("failed to connect database and error in CreateDomain function", err)
		http.Error(w, err.Error(), 500)
	}

	line := r.FormValue("domain")

	if line != "" {

		var idurl int
		err = db.Raw("SELECT id FROM domain WHERE name = ?", line).Scan(&idurl).Error
		if err != nil {
			log.Println("Error: Database query error", err)
			http.Error(w, err.Error(), 500)
		}
		

		if idurl == 0 {

			result := db.Create(&Domain{Name: line, LastCrawlDate: time.Now()})

			if result.Error != nil {
				log.Println("Error: Creating record in domain table")
				http.Error(w, err.Error(), 500)
			}

			var newid int

			err = db.Raw("SELECT id FROM domain order by id desc limit 1").Scan(&newid).Error
			if err != nil {
				log.Println("Error: Database query error", err)
				http.Error(w, err.Error(), 500)
			}
			
			/**********************return: an json of the new domain***********************/

			url := dataDomain{
				ID:   int(newid),
				Body: string(line),
			}

			var urls []dataDomain

			urls = append(urls, url)
			content, err := json.Marshal(urls[0])
			if err != nil {
				log.Println("Unmarshal failed: response " + err.Error())

			}

			jsonBody := string(content)
			startSendMessages(jsonBody)
			log.Println(jsonBody)

			/**********************************************************************/
			fmt.Println(newid)

		}
	}

}

func CreateDomainIfnotExist(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDatabase()

	if err != nil {
		log.Println("failed to connect database and error in CreateDomain function", err)
		http.Error(w, err.Error(), 500)
	}

	line := r.FormValue("domain")

	if line != "" {

		var idurl int
		err = db.Raw("SELECT count(*) FROM domain WHERE name = ?", line).Scan(&idurl).Error
		if err != nil {
			log.Println("Error: Database query error", err)
			http.Error(w, err.Error(), 500)

		}
		

		if idurl == 0 {

			result := db.Create(&Domain{Name: line, LastCrawlDate: time.Now()})

			if result.Error != nil {
				log.Println("Error: Creating record in domain table ", result.Error)
				http.Error(w, err.Error(), 500)
			}

		}

		if idurl != 0 {
			var id int
			rows := db.Raw("(SELECT id FROM domain WHERE name = ?) order by id desc limit 1", line).Scan(&id).Rows()
			if err != nil {
				log.Println("Error: Database query error ", err)
				http.Error(w, err.Error(), 500)
			}
			defer rows.Close()
			/**********************return: an json of repeated domain***********************/

			url := dataDomain{
				ID:   int(id),
				Body: string(line),
			}

			var urls []dataDomain

			urls = append(urls, url)
			content, err := json.Marshal(urls[0])
			if err != nil {
				log.Println("Unmarshal failed: response " + err.Error())
				http.Error(w, err.Error(), 500)
			}

			jsonBody := string(content)
			log.Println(jsonBody)

			/**********************************************************************/
			fmt.Println(id)
			fmt.Println(line)

		}
	}

}

// get all domain in postgres db
func GetAllDomains(w http.ResponseWriter, r *http.Request) {
	//This is name endpoint
	log.Println("Endpoint Hit: GetDomain")

	//allDomain: It is the variable that will use all the domains
	allDomain := []Domain{}

	db, err := connectToDatabase()
	if err != nil {
		log.Println("Error: connect to data base")
		http.Error(w, err.Error(), 500)
	}
	// SELECT * FROM `domain`
	err = db.Find(&allDomain).Error
	if err != nil {
		log.Println("GetDomain $e", err)
		log.Println("ERROR: $e", err)
		http.Error(w, err.Error(), 500)
	}

	//print json with all domains
	json.NewEncoder(w).Encode(allDomain)
}

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

func GetTextToJson(w http.ResponseWriter, r *http.Request) {

	url := []Item{}

	// os.Open() opens specific file in read-only mode and this return
	// a pointer of type os.
	FileName := os.Getenv("FILE_NAME")
	file, err := os.Open(FileName)

	if err != nil {
		log.Println("failed to open ", err)
		http.Error(w, err.Error(), 500)
	}

	scanner, err := bufio.NewScanner(file)
	if err != nil {
		log.Println("bufio.NewScanner() ", err)
		http.Error(w, err.Error(), 500)
	}

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//Scaning the text
		item := Item{Name: scanner.Text()}
		url = append(url, item)
	}

	file.Close()

	json.NewEncoder(w).Encode(&url)
}

func GetFileQuery(w http.ResponseWriter, r *http.Request) {

	query_id := r.FormValue("query_id")
	query := fmt.Sprintf("%s%s", query_id, ".sql")
	file_data, err := ioutil.ReadFile("./queries/" + query)
	if err != nil {
		fmt.Println("error in data ", err)
		http.Error(w, err.Error(), 500)

	}
	fmt.Println(string(file_data))

	allDomain := []Domain{}

	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Error: connect to data base")
		http.Error(w, err.Error(), 500)
	}

	// SELECT * FROM `domain` WHERE ...
	err = db.Raw(string(file_data)).Scan(&allDomain).Error
	if err != nil {
		log.Println("Error: Database query error ", err)
		http.Error(w, err.Error(), 500)
	}
	

	if err != nil {
		log.Println("GetDomain $e", err)
		log.Println("ERROR: $e", err)
		http.Error(w, err.Error(), 500)
	}

	//print json with all domains
	json.NewEncoder(w).Encode(allDomain)
}

func GetFileQueryCount(w http.ResponseWriter, r *http.Request) {

	query_id := r.FormValue("query_id")
	query := fmt.Sprintf("%s%s", query_id, ".sql")
	file_data, err := ioutil.ReadFile("./count_queries/" + query)
	if err != nil {
		log.Println("error in data: ", err)
		http.Error(w, err.Error(), 500)

	}
	//fmt.Println(string(file_data))

	number := 0

	db, err := connectToDatabase()
	if err != nil {
		log.Printl("Error: connect to data base ", err)
		http.Error(w, err.Error(), 500)
	}

	// SELECT count(*) FROM `domain` WHERE ...
	err= db.Raw(string(file_data)).Scan(&number).Error
	if err != nil {
		log.Println("Error: Database query error ", err)
		http.Error(w, err.Error(), 500)

	}


	if err != nil {
		log.Println("ERROR: $e", err)
		http.Error(w, err.Error(), 500)
	}

	//print json with all domains
	json.NewEncoder(w).Encode(number)
}

func sendDomainsToSQS(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	query := r.FormValue("query")
	queue := r.FormValue("queue")
	fmt.Fprintf(w, "query = %s\n", query)
	fmt.Fprintf(w, "queue = %s\n", queue)
	/*************************STEP 1: TO EXECUTE count_queries*************************/
	//query_id := r.FormValue("query_id")
	query = fmt.Sprintf("%s%s", query, ".sql")
	queue = fmt.Sprintf("%s%s", queue, ".txt")
	file_data, err := ioutil.ReadFile("./count_queries/" + query)
	if err != nil {
		log.Println("error in data: ", err)

	}
	//fmt.Println(string(file_data))
	fmt.Println("QUERY_COUNT=", query)
	fmt.Println("QUEUE=", queue)

	number := 0

	db, err := connectToDatabase()
	if err != nil {
		log.Printl("Error: connect to data base ", err)
	}

	// SELECT count(*) FROM `domain` WHERE ...
	rows := db.Raw(string(file_data)).Scan(&number).Rows()
	if err != nil {
		log.Println("Error: Database query error ", err)

	}
	defer rows.Close()

	if err != nil {
		log.Println("ERROR: $e", err)
	}

	//print  the number
	json.NewEncoder(w).Encode(number)

	/*************************STEP 2: TO EXECUTE THE QUERY*************************/
	query = r.FormValue("query")
	query = fmt.Sprintf("%s%s", query, ".sql")
	file_data, err = ioutil.ReadFile("./queries/" + query)
	if err != nil {
		fmt.Println("error in data")

	}
	fmt.Println(string(file_data))

	allDomain := []Domain{}

	// SELECT * FROM `domain` WHERE ...
	err= db.Raw(string(file_data)).Scan(&allDomain).Error
	if err != nil {
		log.Println("Error: Database query error ", err)

	}
	

	//fmt.Println(string(file_data))
	fmt.Println("QUERY=", query)

	if err != nil {
		log.Println("GetDomain $e", err)
		log.Println("ERROR: $e", err)
	}

	//print json with all domains
	json.NewEncoder(w).Encode(allDomain)

	/*********************** Step 3. TO SEND THE MESSAGE TO AWS QUEUE**************************/
	FileName := queue
	fmt.Println("filename=", FileName)

	/***getting the text from the files ***/

	file, err := os.Open("./queues/" + FileName)

	if err != nil {
		log.Println("failed to open the File...", err)
	}
	fileScanner := bufio.NewScanner(file)
	if err != nil {
		log.Println("bufio.NewScanner() ", err)
	}

	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())

	}
	queue_zone := lines[1]
	queue_name := lines[0]

	os.Setenv("SQS_REGION", queue_zone)
	os.Setenv("QUEUE_NAME", queue_name)

	fmt.Println("SQS_REGION=", os.Getenv("SQS_REGION"))
	fmt.Println("QUEUE_NAME=", os.Getenv("QUEUE_NAME"))
	/********************************/
	queueBody := "example queue body"
	sess, err := connect_queue.SessionQueue()
	if err != nil {
		log.Println("Got an error sending the message, SessionQueue function has failed:", err)
		return
	}

	// Get URL of queue
	urlResult, err := connect_queue.GetQueueURL(sess, &queue_name)
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

func domainsToSQS(w http.ResponseWriter, r *http.Request) {

	var queryList []Query
	var queuesList []Queue
	var err error

	queryList, err = GetQueryList()
	if err != nil {
		log.Println("Get Query List error: ", err)
		return
	}

	queuesList, err = GetQueueList()
	if err != nil {
		log.Println("Get queue List error: ", err)
		return
	}

	f := FormElements{
		QueryOptionList: queryList,
		QueueOptionList: queuesList,
	}

	placesPageTmpl, err := template.ParseFiles("form.html") //setp 1
	if err != nil {
		log.Println("the form.html file does not exist..!: ", err)
		return
	}
	if err := placesPageTmpl.Execute(w, f); err != nil {
		log.Println("Failed to build page: ", err)
		return
	}

}

func main() {

	checkIfEnvVarsPresent()
	port := os.Getenv("API_PORT")
	http.HandleFunc("/get_domains", GetFileQuery)
	http.HandleFunc("/get_domains_count", GetFileQueryCount)
	http.HandleFunc("/v1/domain", GetAllDomains)
	http.HandleFunc("/v1/add_domain_to_queue", CreateDomain)
	http.HandleFunc("/v1/create_domain_if_not_exists", CreateDomainIfnotExist)
	http.HandleFunc("/domainsToSQS", domainsToSQS)
	http.HandleFunc("/sendDomainsToSQS", sendDomainsToSQS)
	log.Println(http.ListenAndServe(":"+port, nil))

}
