package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strings"
	"sync"
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

// TODO: .env
var wg sync.WaitGroup
var apiURL = "http://localhost:4080/api/enron2/_doc"
var username = "admin"
var password = "Complexpass#123"

func exploreDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	err = filepath.Walk(filepath.Join(dir, "../enron_mail_20110402/maildir"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		wg.Add(1)
		go exploreFile(path)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	wg.Wait()
}

func exploreFile(path string) {
	defer wg.Done()

	emailHeaders := [...]string{"Message-ID:", "Date:", "From:", "To:", "Subject", "Mime-Version:", "Content-Type:", "Content-Transfer-Encoding:", "X-From:", "X-To:", "X-cc:", "X-bcc:", "X-Folder:", "X-Origin:", "X-FileName:"}
	var bodyBuilder strings.Builder
	var storeContent bool
	headers := make(map[string]string, len(emailHeaders))

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, header := range emailHeaders {
			if strings.HasPrefix(line, header) {
				headers[header] = line[len(header):]
				if strings.HasPrefix(line, "X-FileName:") {
					storeContent = true
				}
				break
			}
		}
		if storeContent && !strings.HasPrefix(line, "X-FileName:") {
			bodyBuilder.WriteString(line)
			bodyBuilder.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
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
		Body:                    bodyBuilder.String(),
	}

	jsonData, err := json.Marshal(email)
	if err != nil {
		log.Println(err)
		return
	}

	sendFiles(jsonData)
}

func sendFiles(data []byte) {
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	exploreDir()
}
