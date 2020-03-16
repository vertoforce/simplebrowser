# Simple Browser

[![Go Report Card](https://goreportcard.com/badge/github.com/vertoforce/simplebrowser)](https://goreportcard.com/report/github.com/vertoforce/simplebrowser)
[![Documentation](https://godoc.org/github.com/vertoforce/simplebrowser?status.svg)](https://godoc.org/github.com/vertoforce/simplebrowser)

This provides a simple interface to get the HTML of a page and/or take a screenshot after javascript runs on the page.

This library uses [chromedp](https://github.com/chromedp/chromedp)

## Usage

```go
var html string
_ := NewPageRequest("https://google.com").WithHTMLGet(&html).Do(context.Background())
fmt.Println(html)
```

### All modifiers

```go
WithCookies([]http.Cookie)
WithHeaders(network.Headers)
WithProxy(string)
WithWaitTime(time.Duration)
WithActions(...chromedp.Action)
WithHTMLGet(*string)
WithScreenshotGet(*[]byte)
Do(context.Context)
```
