package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
)

//go:embed test.txt
//only support global scope
var testStr string

//go:embed test.txt
var testByte []byte

//go:embed test.txt
var testFS embed.FS

func main() {
	fmt.Println(testStr)
	fmt.Println(string(testByte))
	readFile, err := testFS.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(readFile))
}
