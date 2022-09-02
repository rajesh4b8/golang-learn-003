package main

import (
	"log"
	"time"
)

func main() {
	// non buffered channel -> no size is defined!
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	log.Println("Non buffered channel", <-ch)

	// buffered channel -> size defined
	bc := make(chan int, 4)
	bc <- 1
	bc <- 2
	bc <- 3

	log.Println("Buffered channel", <-bc, <-bc, <-bc)

	ch1 := make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
		time.Sleep(time.Second)
		ch1 <- 2
		time.Sleep(time.Second)
		ch1 <- 3
		time.Sleep(time.Second)
		ch1 <- 4
		close(ch1)
	}()

	log.Println("Reading from a channel that will be closed")
	// Use this only if the channel will be closed!
	for i := range ch1 {
		log.Println(i)
	}
	log.Println("Exit main()")
}
