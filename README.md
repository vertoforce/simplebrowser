# Simple Browser

This provides a simple interface to get the HTML of a page after javascript is run on it.

## Usage

```go
html, _ := GetPage(context.Background(), "https://google.com", []Cookie{}, map[string]interface{}{}, time.Second)
fmt.Println(html)
```
