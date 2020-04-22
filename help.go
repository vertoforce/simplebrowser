package simplebrowser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// runChromeDP Takes the PageRequest and performs the actual request
func (p *PageRequest) runChromeDP(ctx context.Context) (err error) {
	// Build proxy context
	var cancel context.CancelFunc
	if p.proxy != "" {
		opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.ProxyServer(p.proxy), chromedp.WindowSize(p.screenWidth, p.screenHeight))
		ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	}
	// Build chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Check headers and cookies
	if p.headers == nil {
		p.headers = network.Headers{}
	}
	if p.cookies == nil {
		p.cookies = []http.Cookie{}
	}

	// Build and run actions
	newActions := []chromedp.Action{setheaders(p.headers), setcookies(p.cookies)}
	newActions = append(newActions, p.preActions...)
	newActions = append(newActions, chromedp.Navigate(p.url), chromedp.Sleep(p.waitTime))
	newActions = append(newActions, p.postActions...)
	err = chromedp.Run(ctx, newActions...)

	return err
}

// setcookies returns a task to navigate to a host with the passed cookies set
// on the network request.
func setcookies(cookies []http.Cookie) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// add cookies to chrome
		for _, cookie := range cookies {
			exp := cdp.TimeSinceEpoch(cookie.Expires)
			success, err := network.SetCookie(cookie.Name, cookie.Value).
				WithExpires(&exp).
				WithDomain(cookie.Domain).
				WithHTTPOnly(cookie.HttpOnly).
				Do(ctx)
			if err != nil {
				return err
			}
			if !success {
				return fmt.Errorf("could not set cookie %q to %q", cookie.Name, cookie.Value)
			}
		}
		return nil
	})
}

// setheaders Returns Tasks to set headers
func setheaders(headers network.Headers) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(headers),
	}
}
