package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel so that the main
	// routine waits for the child routines
	// before finishing.
	c := make(chan string)

	for _, link := range links {
		// Go command is to create a routine
		// so that we don't have to wait for each
		// blocking HTTP call.
		go checkLink(link, c)
	}

	// Waiting to get communication from channel
	// is a blocking line of code!
	// We want to wait for communications from every
	// link call so that the main routine doesn't end early.
	// In this case we're just checking the same links infinite times.

	for l := range c {
		// using a function literal so we can wait
		// 5 seconds before re-checking the same link
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
