### Hello! I'm Chris Wiegman

I am a Software Engineering Manager leading teams at SuperFile. I have been working with open-source software since 2007 where I've built one of the largest WordPress security plugins as well as various WordPress plugins and themes, Google Chrome extensions and developer tools for clients of all sizes. A teacher at heart, I've has spoken at numerous conferences around the world on topics such as development, privacy and business as well as taught both computer science and aviation for universities throughout the US.

My latest blog post is: **[{{ .LatestLink.Display }}]({{ .LatestLink.Link }})**

#### More recent posts from my blog:

{{ range $link := .PostLinks }}

- [{{ $link.Display }}]({{ $link.Link }}){{ end }}

If you like my posts take a look at my site, **[chriswiegman.com](https://chriswiegman.com/)**, and subscribe to get them in your favorite feed reader via **[RSS](https://chriswiegman.com/feed/)**.

You can find me on **[Mastodon](https://mastodon.chriswiegman.com/@chris)** or **[LinkedIn](https://www.linkedin.com/in/chriswiegman)** or you can view **[my resume](https://cfw.cx/resume)**.

<sub>Last updated: {{ .Date }}</sub>
