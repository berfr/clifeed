package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Date      time.Time
	FeedTitle string
	Title     string
	URL       string
}

func (i Item) String() string {
	return fmt.Sprintf("%s | %s | \033]8;;%s\a%s\033]8;;\a", i.Date.Format("2006-01-02"), i.FeedTitle, i.URL, i.Title)
}

func handleItems(ch chan Item, done chan bool) {
	lastWeek := time.Now().AddDate(0, -1, 0)
	var lastWeekItems []Item
	for item := range ch {
		if item.Date.After(lastWeek) {
			lastWeekItems = append(lastWeekItems, item)
		}
	}
	sort.Slice(lastWeekItems, func(i, j int) bool {
		return lastWeekItems[i].Date.After(lastWeekItems[j].Date)
	})
	for _, item := range lastWeekItems {
		fmt.Println(item.String())
	}
	done <- true
}

func getFeed(feedURL string, wg *sync.WaitGroup, ch chan Item) {
	defer wg.Done()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedURL)
	for _, item := range feed.Items {
		ch <- Item{*item.PublishedParsed, feed.Title, item.Title, item.Link}
	}
}

func main() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := filepath.Join(dir, "clifeed.txt")

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	ch := make(chan Item)

	for scanner.Scan() {
		feedURL := scanner.Text()
		wg.Add(1)
		go getFeed(feedURL, &wg, ch)
	}

	done := make(chan bool)
	go handleItems(ch, done)

	wg.Wait()
	close(ch)
	<-done
	close(done)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
