//The program is a API that with the name of a domain provided as a parameter through a url,
//obtains the records from the features table with respect to that id,
//then returns these records to a json and prints them.
package main

import (
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
		fmt.Println("Env var. POSTGRES_USER is not set. Check your .env.sh  or .env.bat file")
	case os.Getenv("POSTGRES_PASSWORD"):
		fmt.Println("Env var. POSTGRES_PASSWORD is not set. Check your .env.sh  or .env.bat file")
	case os.Getenv("POSTGRES_DATABASE"):
		fmt.Println("Env var. POSTGRES_DATABASE is not set. Check your .env.sh  or .env.bat file")
	case os.Getenv("API_TOKEN"):
		fmt.Println("Env var. API_TOKEN is not set. Check your .env.sh  or .env.bat file")
	case os.Getenv("API_PORT"):
		fmt.Println("Env var. API_PORT is not set. Check your .env.sh or .env.bat file")

	default:
		return
	}
	os.Exit(1)
}

type Feature struct {
	Id            int       `json:"id" gorm:"id"`
	Name          string    `json:"name" gorm:"name"`
	Value         string    `json:"value" gorm:"value"`
	LastCrawlDate time.Time `json:"last_crawl_date" gorm:"last_crawl_date"`
	Domain_id     int       `json:"domain_id" gorm:"domain_id"`
	Execution_id  int       `json:"execution_id" gorm:"execution_id"`
}

type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	// contains filtered or unexported fields
}

type Domain struct {
	Id            uint      `json:"id" gorm:"id"`
	Name          string    `json:"name" gorm:"name"`
	LastCrawlDate time.Time `json:"last_crawl_date" gorm:"last_crawl_date"`
}

func (Domain) TableName() string {
	return "domain"
}
func (Feature) TableName() string {
	return "feature"
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

// get the records from the feature table according to domain id
func GetFeature(w http.ResponseWriter, r *http.Request) {
	//This is name endpoint
	log.Println("Endpoint Hit:: GetFeature")

	db, err := connectToDatabase()

	if err != nil {
		log.Println("failed to connect database and error in CreateFeature function", err)
	}

	domain_name := r.FormValue("domain")
	if domain_name != "" {

		var idurl int
		err = db.Raw("SELECT id FROM domain WHERE name = ?", domain_name).Scan(&idurl).Error
		if err != nil {
			log.Println("Error: Database query error", err)

		}

		if idurl == 0 {
			log.Println("Domain name not found...")
		}
		if idurl != 0 {

			allFeature := []Feature{}
			//db.Where("domain_id = ?", idurl).Last(&allFeature)
			// SELECT * FROM feature WHERE domain_id = idurl;
			db.Where("domain_id = ?", idurl).Find(&allFeature)

			//print json with all feature
			json.NewEncoder(w).Encode(allFeature)
		}
	}

}

func main() {
	checkIfEnvVarsPresent()
	port := os.Getenv("API_PORT")

	http.HandleFunc("/v1/domain", GetDomain) //.Methods("GET")
	http.HandleFunc("/get_domain_contents", GetFeature)
	log.Println(http.ListenAndServe(":"+port, nil))
}
