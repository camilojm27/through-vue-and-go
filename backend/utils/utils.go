package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	// "sync"
)

//var cache sync.Map

func MakeAPIRequest(page string, search string, docId string) ([]byte, error) {
	apiURL := "http://192.168.1.16:4080/api/enron2/_search"
	method := "POST"
	username := "admin"
	password := "Complexpass#123"
	body := ""

	switch {
	case docId != "":
		apiURL = "http://192.168.1.16:4080/api/enron2/_doc/" + docId
		method = "GET"
	case search != "":

		body = `{
    "search_type": "match",
    "query": {
        "term": "%s",
        "field": "_all"
    },
    "max_results": 200,
    "_source": [
    ]
}`
		body = fmt.Sprintf(body, search)
	default:
		body = `{
    "search_type": "alldocuments",
    "from":%s,
    "max_results": 20,
    "_source": []
  }`
		body = fmt.Sprintf(body, page)
	}

	req, err := http.NewRequest(method, apiURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatal(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Add Basic Authentication header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", basicAuth)

	// Send the request
	client := &http.Client{}
	fmt.Println(req.Body)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// // Cache the response with a 5-minute expiration
	// cache.Store(apiURL, responseBody)
	// go func() {
	// 	time.Sleep(5 * time.Minute)
	// 	cache.Delete(apiURL)
	// }()

	return responseBody, nil
}
