### Hi! I'm Chris Wiegman

I am a developer, educator, blogger and open-source contributor dedicated to creating human-centered, privacy-focused technology. Over the past two decades, I have built everything from WordPress plugins (including iThemes Security) to Chrome extensions, development environments and developer tools. A teacher at heart, I have lectured in aviation and computer science, mentored emerging technologists, and spoken at conferences worldwide. A former pilot turned software engineer, I remain passionate about building ethical, sustainable tools for real people.

_My latest blog post_: **[{{ .LatestLink.Display }}]({{ .LatestLink.Link }})**

#### More recent posts from my blog:

{{ range $link := .PostLinks }}

- [{{ $link.Display }}]({{ $link.Link }}){{ end }}

If you like my posts take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)**, and subscribe to get them in your favorite feed reader via **[RSS](https://chriswiegman.com/feed)**.

You can find me on **[Mastodon](https://mastodon.chriswiegman.com/@chris)** and **[LinkedIn](https://www.linkedin.com/in/chriswiegman)** or you can view **[my resume](https://cwie.co/resume)**.

<sub>Last updated: {{ .Date }}</sub>
