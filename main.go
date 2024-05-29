package main

import (
	"log"
	"net/http"
	"time"
)

const (
	WEBPORT = ":8003"
)

func main() {

	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/check-health", checkHealth)

	log.Println("Server started on port", WEBPORT)

	err := http.ListenAndServe(WEBPORT, nil)
	if err != nil {
		log.Println("Error Starting the server...\n", err.Error())
	}

}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	response := []byte("Hello, World!")
	_, err := writer.Write(response)
	if err != nil {
		http.Error(writer, "Error Writing the response"+err.Error(), http.StatusInternalServerError)
	}
}

func checkHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	response := []byte("OK! :)")
	_, err := writer.Write(response)
	time.Sleep(time.Second)
	if err != nil {
		http.Error(writer, "Error Writing the response"+err.Error(), http.StatusInternalServerError)
	}
}
