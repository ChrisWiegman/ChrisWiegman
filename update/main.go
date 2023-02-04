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

	postLine := ""

	if len(feed.Items) >= 6 {

		index := 1

		postLine = "#### More recent posts\n\n"

		for index <= 5 {
			post := feed.Items[index]

			postLine += "* [" + post.Title + "](" + post.Link + ")\n"

			index++
		}

		postLine += "\n\n"
	}

	// Get the newest blog post
	latestPost := feed.Items[0]

	// Whisk together static and dynamic content until stiff peaks form
	line1 := "### Hi! I am Chris Wiegman"
	line2 := "I am an **Engineering Manager** at **[WP Engine](https://wpengine.com/)**, devoted to improving the developer experience for WordPress developers, particularly those working on headless projects. My work, both for my day job and side projects, focuses on the intersection of development, privacy, ethics and usability of software and development to help improve the lives of everyone who uses modern technology."
	line3 := "My latest blog post is: **[" + latestPost.Title + "](" + latestPost.Link + ")**. "
	line4 := "Like this type of content? Take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)** and subscribe via its **[RSS feed](https://chriswiegman.com/feed/)**."
	line5 := "You can find me on **[Mastodon](https://mastodon.chriswiegman.com/@chris)** and **[Pixelfed](https://pixelfed.chriswiegman.com/@chris)** or view **[my full resume](https://gist.github.com/ChrisWiegman/8a89d7c2aca775884ae4227ca3b5be01)**."
	line6 := "<sub>Last updated: " + time.Now().Format("January 2, 2006") + "</sub>"

	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s%s\n\n%s\n\n%s\n", line1, line2, line3, postLine, line4, line5, line6)

	file, err := os.Create("../README.md")
	if err != nil {
		log.Fatalf("error creating README.md: %v", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.WriteString(file, content)
	if err != nil {
		log.Fatalf("error writing README.md: %v", err)
		os.Exit(1)
	}
}
