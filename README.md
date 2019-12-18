# Simple Browser

This provides a simple interface to get the HTML of a page after javascript is run on it.

This library uses [chromedp](https://github.com/chromedp/chromedp)

## Usage

```go
html, _ := GetPage(context.Background(), "https://google.com", []Cookie{}, map[string]interface{}{}, time.Second)
fmt.Println(html)
```
