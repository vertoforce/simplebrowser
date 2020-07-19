package simplebrowser

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

var webdriverJS = `// Overwrite the plugins property to use a custom getter.
Object.defineProperty(navigator, 'plugins', {
  get: () => [1, 2, 3, 4, 5],
});
`

func TestJavascript(t *testing.T) {
	var screenshot []byte
	err := NewPageRequest("https://intoli.com/blog/making-chrome-headless-undetectable/chrome-headless-test.html").WithWaitTime(time.Second * 3).WithScreenshotGet(&screenshot).WithJavascript(webdriverJS).Do(context.Background())
	fmt.Println("test")
	if err != nil {
		t.Error(err)
	}
	ioutil.WriteFile("out.png", screenshot, 0644)
}
