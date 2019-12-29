// Package simplebrowser makes it easy to get the HTML of a page after javascript runs
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

// GetPage Get HTML of page after waiting waitTime for javascript to run
func GetPage(ctx context.Context, URL string, cookies []http.Cookie, headers network.Headers, waitTime time.Duration) (html string, err error) {
	ctxN, cancel := chromedp.NewContext(ctx)
	defer cancel()

	if headers == nil {
		headers = network.Headers{}
	}
	if cookies == nil {
		cookies = []http.Cookie{}
	}

	err = chromedp.Run(ctxN,
		setheaders(headers),
		setcookies(cookies),
		chromedp.Navigate(URL),
		chromedp.Sleep(waitTime),
		chromedp.OuterHTML("html", &html),
	)
	return html, err
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
