// Package simplebrowser makes it easy to get the HTML of a page after javascript runs
package simplebrowser

import (
	"context"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

const (
	DefaultWaitTime = time.Second * 2
)

// PageRequest A request object to get a page with optional cookies, headers, and proxy
type PageRequest struct {
	url          string
	cookies      []http.Cookie
	headers      network.Headers
	waitTime     time.Duration
	actions      []chromedp.Action
	screenWidth  int
	screenHeight int
	proxy        string
}

// NewPageRequest returns a basic PageRequest with default WaitTime
func NewPageRequest(URL string) *PageRequest {
	return &PageRequest{
		waitTime:     DefaultWaitTime,
		url:          URL,
		screenHeight: 920,
		screenWidth:  1090,
	}
}

// WithCookies Add cookies to PageRequest
func (p *PageRequest) WithCookies(cookies []http.Cookie) *PageRequest {
	p.cookies = cookies
	return p
}

// WithHeaders Add headers to PageRequest
func (p *PageRequest) WithHeaders(headers network.Headers) *PageRequest {
	p.headers = headers
	return p
}

// WithProxy Add Proxy to PageRequest.  A proxy can be a string like socks4://ip:port
func (p *PageRequest) WithProxy(proxy string) *PageRequest {
	p.proxy = proxy
	return p
}

// WithScreenSize Sets the screen size of the window
func (p *PageRequest) WithScreenSize(width, height int) *PageRequest {
	p.screenWidth = width
	p.screenHeight = height
	return p
}

// WithWaitTime Add specific wait time to PageRequest, WaitTime is the time it will wait for the page to load before performing further actions
func (p *PageRequest) WithWaitTime(waitTime time.Duration) *PageRequest {
	p.waitTime = waitTime
	return p
}

// WithActions Add actions to request
func (p *PageRequest) WithActions(actions ...chromedp.Action) *PageRequest {
	p.actions = append(p.actions, actions...)
	return p
}

// WithHTMLGet Will place the html of the page in the string after the request is made
func (p *PageRequest) WithHTMLGet(html *string) *PageRequest {
	return p.WithActions(chromedp.OuterHTML("html", html))
}

// WithScreenshotGet Will get a screenshot of the page after the request is made and load it in to the []byte pointer in png format.
func (p *PageRequest) WithScreenshotGet(pngScreenshot *[]byte) *PageRequest {
	return p.WithActions(chromedp.CaptureScreenshot(pngScreenshot))
}

// Do Perform the actual PageRequest
func (p *PageRequest) Do(ctx context.Context) error {
	return p.runChromeDP(ctx)
}
