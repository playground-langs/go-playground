package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	myChan := make(chan string)
	go greeting(myChan)
	receive := <-myChan
	fmt.Println(receive)
	sizes := make(chan Page)
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}
}

type Page struct {
	Url  string
	Size int
}

func responseSize(url string, ch chan Page) {
	fmt.Println("getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- Page{
		Url:  url,
		Size: len(body),
	}
}

func greeting(ch chan string) {
	ch <- "hi"
}
