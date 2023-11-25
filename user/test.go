package main

import (
	"fmt"
		"log"
		"net/http"
)

func test() {
	// this is just a test file to check if github action is working correctly or not
	log.Println("This is a test file")
	fmt.Println("This is a test file")
	err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Fatal("Error in starting the server")
	}
}
