### Hi! I'm Chris Wiegman

I'm a tech leader, mentor, blogger, and open-source contributor dedicated to building human-centered, privacy-focused technology. Over the past two decades, I've built everything from WordPress plugins to Chrome extensions, development environments, and developer tools; and now I lead the teams who build what comes next. A teacher at heart, I've lectured in aviation and computer science, mentored emerging technologists, and spoken at conferences worldwide. A former pilot turned software engineer, I remain passionate about ethical, sustainable technology and the people who create it.

_My latest post_: **[{{ .LatestLink.Display }}]({{ .LatestLink.Link }})**

#### More recent posts from my blog:

{{ range $link := .PostLinks }}

- [{{ $link.Display }}]({{ $link.Link }}){{ end }}

If you like my posts take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)**, and subscribe to get them in your favorite feed reader via **[RSS](https://chriswiegman.com/index.xml)**.

You can also find me on **[Mastodon](https://mastodon.chriswiegman.com/@chris)**, **[Bluesky](https://bsky.app/profile/chriswiegman.com)** and **[LinkedIn](https://www.linkedin.com/in/chriswiegman)**.

<sub>Last updated: {{ .Date }}</sub>
