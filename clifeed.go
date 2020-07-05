package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	file, err := os.Open("feeds.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fp := gofeed.NewParser()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feedURL := scanner.Text()
		feed, _ := fp.ParseURL(feedURL)
		fmt.Println(feed.Title)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
