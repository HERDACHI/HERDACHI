// Program that uses a loop that runs the POST request of a list  of domains inside a file
// If the URL is accesible the POST request is procesed
// if the URL is not accesible, the URL is is sent to a file

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"net/url"
	"os"
)

//Function that reads a file line by line
func main() {

	//the filename is passed as an argument

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error while opening file. Err: %s", err)
	}
	//defer file.Close()
	//function that reads each line of the file
	scanner := bufio.NewScanner(file)
	//Creation of a file that will contain the URL that doesn't exist
	//f, err := os.Create("domains_errors.txt")
	//if err != nil {
	//	fmt.Printlna("Cannot create file", err)
	//	return
	//
	for scanner.Scan() {
		//for each line on the file, we save the url on a variable
		lineDomain := scanner.Text()

		//u, err := url.Parse(lineDomain)
		//if err != nil {
		//	fmt.Println("Error: URL does not exist", lineDomain, u)s
		//	f.WriteString(lineDomain + "\n")
		//}

		//Curl commando converted to GO that that runs the POST request
		client := &http.Client{}
		url := "http://3.145.171.107/websitereview/index.php/en/parse/index?Website%5Bdomain%5D=" + lineDomain
		log.Println(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Http.newrequest fails:", err)
			continue
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:93.0) Gecko/20100101 Firefox/93.0")
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("DT", "1")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Referer", "http://3.145.171.107/websitereview/index.php/en/www/abc.com")
		resp, err := client.Do(req)
		if err != nil {
			log.Println("resp.do fails:", err)
			continue
		}
		defer resp.Body.Close()
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("body.readAll fails:", err)
			continue
		}
		fmt.Printf("%s\n", bodyText)

	}
}
