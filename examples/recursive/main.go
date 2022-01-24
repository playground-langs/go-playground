package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	err := scanDirWithError("./")
	if err != nil {
		log.Fatal(err)
	}
	defer reportPanic()
	panic("test unexpected error")
	scanDirWithPanic("./test")
}

func scanDirWithError(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err = scanDirWithError(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func scanDirWithPanic(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirWithPanic(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		//恢复非预期错误
		panic(err)
	}
}
