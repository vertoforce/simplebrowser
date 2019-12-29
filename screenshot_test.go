package simplebrowser

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestScreenshot(t *testing.T) {
	ctxN, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	URL := "https://google.com"

	var picture []byte

	err := chromedp.Run(ctxN,
		chromedp.Navigate(URL),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(&picture),
	)
	fmt.Println(err)

	ioutil.WriteFile("out.png", picture, os.ModeAppend)
}
