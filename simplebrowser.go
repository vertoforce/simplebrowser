// Package simplebrowser makes it easy to get the HTML of a page after javascript runs
package simplebrowser

import (
	"context"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// GetPage Get HTML of page after waiting waitTime for javascript to run
func GetPage(ctx context.Context, URL string, cookies []http.Cookie, headers network.Headers, waitTime time.Duration) (html string, err error) {
	return GetPageProxy(ctx, URL, cookies, headers, waitTime, "")
}

// GetPageProxy is the same as GetPage but uses a proxy
func GetPageProxy(ctx context.Context, URL string, cookies []http.Cookie, headers network.Headers, waitTime time.Duration, proxy string) (html string, err error) {
	err = runChromeDP(ctx,
		URL,
		cookies,
		headers,
		waitTime,
		proxy,
		chromedp.OuterHTML("html", &html),
	)
	return html, err
}

// GetPageScreenshot same as GetPage but takes a screenshot of the page
func GetPageScreenshot(ctx context.Context, URL string, cookies []http.Cookie, headers network.Headers, waitTime time.Duration, proxy string) (pngScreenshot []byte, err error) {
	err = runChromeDP(ctx,
		URL,
		cookies,
		headers,
		waitTime,
		proxy,
		chromedp.CaptureScreenshot(&pngScreenshot),
	)
	return pngScreenshot, err
}
