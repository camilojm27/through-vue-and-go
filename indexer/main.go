package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	S "strings"
)

type Email struct {
	MessageId               string `json:"Message-Id"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	CC                      string `json:"CC"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Body                    string `json:"Body"`
}

func explore_dir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	//fmt.Println(dir)
	err = filepath.Walk(dir+"/../enron_mail_20110402/maildir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		explore_file(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func explore_file(path string) {
	email_headers := [15]string{"Message-ID:", "Date:", "From:", "To:", "Subject", "Mime-Version:", "Contente-Type:", "Content-Transfer-Encoding:", "X-From:", "X-To:", "X-cc:", "X-bcc:", "X-Folder:", "X-Origin:", "X-FileName:"}
	temp := email_headers[:]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var body string
	var storeContent bool
	headers := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		for i, header := range email_headers {
			if S.HasPrefix(line, header) {
				//fmt.Println(line)
				temp = append(temp[:i], temp[i+1:]...)
				if S.HasPrefix(line, "X-FileName:") {
					storeContent = true
				}
				headers[header] = line[len(header):]
				break
			}
		}
		if storeContent {
			if !S.HasPrefix(line, "X-FileName:") {
				body += line + "\n"
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	email := Email{
		MessageId:               headers["Message-ID:"],
		Date:                    headers["Date:"],
		From:                    headers["From:"],
		To:                      headers["To:"],
		CC:                      headers["CC:"],
		Subject:                 headers["Subject"],
		MimeVersion:             headers["Mime-Version:"],
		ContentType:             headers["Content-Type:"],
		ContentTransferEncoding: headers["Content-Transfer-Encoding:"],
		XFrom:                   headers["X-From:"],
		XTo:                     headers["X-To:"],
		Xcc:                     headers["X-cc:"],
		Xbcc:                    headers["X-bcc:"],
		XFolder:                 headers["X-Folder:"],
		XOrigin:                 headers["X-Origin:"],
		XFileName:               headers["X-FileName:"],
		Body:                    body,
	}

	jsonData, err := json.Marshal(email)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonData))
	send_files(jsonData)
}

func send_files(data []byte) {

	apiURL := "http://localhost:4080/api/enron/_doc"
	username := "admin"
	password := "Complexpass#123"

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Add Basic Authentication header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// // Read the response body
	// responseBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Process the response body
	// fmt.Println(string(responseBody))

}

func main() {
	explore_dir()
	//explore_file("../enron_mail_20110402/maildir/brawner-s/all_documents/2.")
}
