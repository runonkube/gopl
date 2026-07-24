package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/issues", GetIssues)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func GetIssues(responseWriter http.ResponseWriter, request *http.Request) {

	fmt.Fprint(responseWriter, "HelloWorld")
}
