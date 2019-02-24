package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkSite(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkSite(url, c)
		}(u)
	}
}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "may be down!")
		c <- url
		return
	}
	fmt.Println(url, "is up!")
	c <- url
}
