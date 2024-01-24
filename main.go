package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	fmt.Println(dir)
	folders, err := os.ReadDir(dir + "/enron_mail_20110402/maildir")

	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range folders {
		fmt.Println(file.Name())
	}
}
