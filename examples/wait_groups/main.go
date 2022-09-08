package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	links := []string{
		"http://linkedin.com",
		"http://yahoo.com",
		"http://learningnet.com",
		"http://facebook.com",
		"http://google.com",
	}

	for _, link := range links {
		wg.Add(1)
		go checkHealth(link, &wg)
	}

	// time.Sleep(5 * time.Second)
	wg.Wait()
	log.Println("Exit main()")
}

func checkHealth(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	if _, err := http.Get(link); err != nil {
		log.Printf("%v might be down\n", link)
		return
	}

	log.Printf("%v is up!\n", link)
}
