package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func setupRoutes() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "homepage ")

}

func main() {
	fmt.Println("Hello World")
	//setupRoutes
	//var GCPKEY=[]byte(os.Getenv("GCPAPI"))

	fmt.Println([]byte(os.Getenv("GCPAPI")))
}
