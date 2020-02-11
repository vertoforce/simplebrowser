package simplebrowser

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// runChromeDP Acts as the baseline running of chromeDP with optional additional actions.
// In the default case it will configure chromedp to use a proxy, include the cookies and headers, navigate to the URL, and then sleep the wait time
func runChromeDP(ctx context.Context, URL string, cookies []http.Cookie, headers network.Headers, waitTime time.Duration, proxy string, actions ...chromedp.Action) (err error) {
	// Build proxy context
	var cancel context.CancelFunc
	if proxy != "" {
		opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.ProxyServer(proxy))
		ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	}
	// Build chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Check headers and cookies
	if headers == nil {
		headers = network.Headers{}
	}
	if cookies == nil {
		cookies = []http.Cookie{}
	}

	// Build and run actions
	newActions := []chromedp.Action{setheaders(headers), setcookies(cookies), chromedp.Navigate(URL), chromedp.Sleep(waitTime)}
	for _, action := range actions {
		newActions = append(newActions, action)
	}
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
