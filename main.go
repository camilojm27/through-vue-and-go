package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	S "strings"
)

type Email struct {
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func explore_dir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	//fmt.Println(dir)
	err = filepath.Walk(dir+"/enron_mail_20110402/maildir", func(path string, info os.FileInfo, err error) error {
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
		Headers: headers,
		Body:    body,
	}

	jsonData, err := json.Marshal(email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func send_files() {
	
}

func main() {
	//explore_dir()
	explore_file("enron_mail_20110402/maildir/brawner-s/all_documents/2.")
}
