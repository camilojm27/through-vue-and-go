package utils

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var cache sync.Map

func MakeAPIRequest() ([]byte, error) {
	apiURL := "http://localhost:4080/api/enron/_search"
	username := "admin"
	password := "Complexpass#123"
	body := `{
    "search_type": "alldocuments",
    "max_results": 100,
    "_source": []
}`

	// Check if the response is already cached
	if cachedResponse, ok := cache.Load(apiURL); ok {
		return cachedResponse.([]byte), nil
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(body)))
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

	// Cache the response with a 5-minute expiration
	cache.Store(apiURL, responseBody)
	go func() {
		time.Sleep(5 * time.Minute)
		cache.Delete(apiURL)
	}()

	return responseBody, nil
}
