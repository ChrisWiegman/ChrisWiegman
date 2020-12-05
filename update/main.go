package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {

	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://chriswiegman.com/feed/")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
		os.Exit(1)
	}

	// Get the freshest item
	latestPost := feed.Items[0]

	date := time.Now().Format("January 2, 2006")

	// Whisk together static and dynamic content until stiff peaks form
	line1 := "### Hi! I am Chris Wiegman"
	line2 := "I am a **Senior Software Engineer**, currently at **[WP Engine](https://wpengine.com/)**, devoted to improving the developer experience for WordPress developers of all kinds. My work focuses on the intersection of development, privacy, ethics and usability of software and development to help improve the lives of everyone who uses modern technology."
	line3 := "For a full list of my personal projects, please visit **[gitea.chriswiegman.com](https://gitea.chriswiegman.com/)**."
	line4 := "My latest blog post is: **[" + latestPost.Title + "](" + latestPost.Link + ")**. "
	line5 := "Like this type of content? Take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)** and subscribe via its **[RSS feed](https://chriswiegman.com/feed/)**."
	line6 := "<sub>Last updated: " + date + "</sub>"
	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n", line1, line2, line3, line4, line5, line6)

	// Prepare file with a light coating of os
	file, err := os.Create("../README.md")
	if err != nil {
		log.Fatalf("error creating README.md: %v", err)
		os.Exit(1)
	}

	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, content)
	if err != nil {
		log.Fatalf("error writing README.md: %v", err)
		os.Exit(1)
	}
}
