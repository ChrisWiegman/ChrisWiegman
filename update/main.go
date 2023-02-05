package main

import (
	_ "embed"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/mmcdole/gofeed"
)

//go:embed template.md
var markdownTemplate string

type templateLink struct {
	Link, Display string
}

type templateVars struct {
	Date      string
	PostLinks []templateLink
}

func main() {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://chriswiegman.com/feed/")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
		os.Exit(1)
	}

	templateVars := templateVars{
		Date: time.Now().Format("January 2, 2006"),
	}

	if len(feed.Items) >= 6 {

		index := 0

		for index <= 5 {
			link := templateLink{
				Display: feed.Items[index].Title,
				Link:    feed.Items[index].Link,
			}
			templateVars.PostLinks = append(templateVars.PostLinks, link)

			index++
		}
	}

	tmpl := template.Must(template.New("kanaPlugin").Parse(markdownTemplate))

	myFile, _ := os.Create("../README.md")

	tmpl.Execute(myFile, templateVars)
}
