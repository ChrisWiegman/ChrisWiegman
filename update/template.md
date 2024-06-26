### Hello! I'm Chris Wiegman

I am a Sr. Staff Software Engineer on the teams building **[Atlas](https://wpengine.com/headless-wordpress/)** for **[WP Engine](https://wpengine.com/)**. Throughout my career I've lead teams building everything from JavaScript frameworks, WordPress plugins, Chrome extensions, development environments and more. When not building software I enjoy teaching, mentoring and **[speaking](https://chriswiegman.com/speaking/)** and have been fortunate to do so at dozens of conferences, Meetups and other events. Today my work, both personal and professional, focuses on the intersection of development, privacy, ethics and usability of software and development to help improve the lives of everyone who uses modern technology.

My latest blog post is: **[{{ .LatestLink.Display }}]({{ .LatestLink.Link }})**

#### More recent posts from my blog:

{{ range $link := .PostLinks }}

- [{{ $link.Display }}]({{ $link.Link }}){{ end }}

If you like my posts take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)**, and subscribe to get them in your favorite feed reader via **[RSS](https://chriswiegman.com/feed/)**.

You can find me on **[Mastodon](https://mastodon.chriswiegman.com/@chris)** or **[LinkedIn](https://www.linkedin.com/in/chriswiegman)** or you can view **[my resume](https://cfw.cx/resume)**.

<sub>Last updated: {{ .Date }}</sub>
