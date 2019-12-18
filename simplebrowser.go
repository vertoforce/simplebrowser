// Package simplebrowser makes it easy to get the HTML of a page after javascript runs
package simplebrowser

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type Cookie struct {
	Name  string
	Value string
}

// GetPage Get HTML of page after waiting waitTime for javascript to run
func GetPage(ctx context.Context, URL string, cookies []Cookie, headers map[string]interface{}, waitTime time.Duration) (html string, err error) {
	ctxN, cancel := chromedp.NewContext(ctx)
	defer cancel()

	err = chromedp.Run(ctxN,
		chromedp.Navigate(URL),
		chromedp.Sleep(waitTime),
		setcookies(cookies),
		setheaders(headers),
		chromedp.OuterHTML("html", &html),
	)
	return html, err
}

// setcookies returns a task to navigate to a host with the passed cookies set
// on the network request.
func setcookies(cookies []Cookie) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		// create cookie expiration
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		// add cookies to chrome
		for _, cookie := range cookies {
			success, err := network.SetCookie(cookie.Name, cookie.Value).
				WithExpires(&expr).
				WithDomain("localhost").
				WithHTTPOnly(true).
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
func setheaders(headers map[string]interface{}) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
	}
}
