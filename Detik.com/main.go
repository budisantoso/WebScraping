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
	resp, err := http.Get("https://www.detik.com/")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Error status code of %d : %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul.list_fokus article a div h2").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("span.normal").Text())
		if title != "" {
			fmt.Println(title)
		}
	})
}
