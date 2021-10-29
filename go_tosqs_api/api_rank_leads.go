//The program is an API that takes a domain name from a url parameter,
//saves it in a Postgres database, and sends it to an AWS queue.
package main

import (
	"api_tosqs_go/connect_queue"
	"encoding/json"
	"fmt"

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
		log.Fatalf("Error: connect to data base")
	}
	// SELECT * FROM `domain`
	err = db.Find(&allDomain).Error
	if err != nil {
		log.Println("GetDomain $e", err)
		log.Println("ERROR: $e", err)
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
	}

	line := r.FormValue("domain")

	if line != "" {

		var idurl int
		db.Raw("SELECT id FROM domain WHERE name = ?", line).Scan(&idurl)
		if idurl == 0 {

			result := db.Create(&Domain{Name: line, LastCrawlDate: time.Now()})

			if result.Error != nil {
				log.Println("Error: Creating record in domain table")

			}
			var newid int
			db.Raw("SELECT id FROM domain order by id desc limit 1").Scan(&newid)

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
	}

	line := r.FormValue("domain")

	if line != "" {

		var idurl int
		db.Raw("SELECT id FROM domain WHERE name = ?", line).Scan(&idurl)
		if idurl == 0 {

			result := db.Create(&Domain{Name: line, LastCrawlDate: time.Now()})

			if result.Error != nil {
				log.Println("Error: Creating record in domain table")

			}

		}

		if idurl != 0 {
			var id int
			db.Raw("(SELECT id FROM domain WHERE name = ?) order by id desc limit 1", line).Scan(&id)

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

			}

			jsonBody := string(content)
			log.Println(jsonBody)

			/**********************************************************************/
			fmt.Println(id)
			fmt.Println(line)

		}
	}

}

func GetQueryDomain(w http.ResponseWriter, r *http.Request) {

}

/************ this code will be used later ****************
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}*/
/*
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	fmt.Println(bodyBuf)

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}*/

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

func main() {
	checkIfEnvVarsPresent()
	port := os.Getenv("API_PORT")
	//target_url := "http://localhost:" + port + "/upload"
	//filename := "./domainslist.txt"
	//postFile(filename, target_url)
	//http.HandleFunc("/v1/upload", upload)

	http.HandleFunc("/v1/domain", GetDomain)                                   //.Methods("GET")
	http.HandleFunc("/v1/add_domain_to_queue", CreateDomain)                   //.Methods("POST")
	http.HandleFunc("/v1/create_domain_if_not_exists", CreateDomainIfnotExist) //.Methods("POST")
	http.HandleFunc("/v1/filter_get_domain", GetQueryDomain)                   //.Methods("POST")
	log.Println(http.ListenAndServe(":"+port, nil))

}
