package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mmcdole/gofeed"
)

func getFeed(feedURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedURL)
	fmt.Println(feed.Title)
}

func main() {
	file, err := os.Open("feeds.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup

	for scanner.Scan() {
		feedURL := scanner.Text()
		wg.Add(1)
		go getFeed(feedURL, &wg)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
