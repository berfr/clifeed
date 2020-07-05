package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://berfr.github.io/index.xml")
	fmt.Println(feed.Title)
}
