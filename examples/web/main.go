package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salute", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salute, web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
