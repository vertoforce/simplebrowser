package simplebrowser

import (
	"context"
	"io/ioutil"
	"testing"
	"time"
)

func TestScreenshot(t *testing.T) {
	var screenshot []byte
	err := NewPageRequest("http://ip4.me").WithWaitTime(time.Second * 3).WithScreenshotGet(&screenshot).Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	ioutil.WriteFile("out.png", screenshot, 0644)
}
