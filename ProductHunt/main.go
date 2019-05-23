package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	ScrapeHTML()
}

//ScrapeHTML is function to scrape
func ScrapeHTML() {
	resp, err := http.Get("https://www.producthunt.com/")

	if err != nil {
		fmt.Println("error in opening web")
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error is %v %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul li div a").Each(func(i int, s *goquery.Selection) {
		product := strings.TrimSpace(s.Find(".content_31491 .font_9d927").Text())
		fmt.Println(product)
	})
}
