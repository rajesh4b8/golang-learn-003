package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	links := []string{
		"http://linkedin.com",
		"http://yahoo.com",
		"http://learningnet.com",
		"http://facebook.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkHealth(link, c)
	}

	// reading from a channel is a blocking operation...
	for range links {
		log.Println(<-c)
	}

	log.Println("Exit main()")
}

func checkHealth(link string, c chan string) {

	if _, err := http.Get(link); err != nil {
		// log.Printf("%v might be down\n", link)
		c <- fmt.Sprintf("%v might be down", link)
		return
	}

	// log.Printf("%v is up!\n", link)
	c <- fmt.Sprintf("%v is up!\n", link)
}
