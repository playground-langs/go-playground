package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 5; i++ {
			send := i + 1
			ch <- send
			time.Sleep(500 * time.Millisecond)
			log.Println(send)
		}
	}()
	time.Sleep(3 * time.Second)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	log.Println(<-ch)
	time.Sleep(2 * time.Second)
}
