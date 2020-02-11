package simplebrowser

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

const HTMLTest = `<html>
<script>
    function changeContent() {
        document.getElementById("content").innerHTML = "test";
    }
    window.onload = changeContent;
</script>

<body>
    <div id="content">the content</div>
</body>

</html>`

func TestGetPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("test")
		if err != nil {
			t.Errorf(err.Error())
		} else {
			if cookie.Value != "testValue" {
				t.Errorf("Error getting test cookie")
			}
		}
		w.Write([]byte(HTMLTest))
	}))
	defer ts.Close()

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("Error parsing URL")
		return
	}

	html, err := GetPage(context.Background(), ts.URL, []http.Cookie{{Name: "test", Value: "testValue", Domain: URL.Hostname(), Expires: time.Now().Add(time.Hour * 5000)}}, nil, time.Second)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if strings.Index(html, `<div id="content">test</div>`) == -1 {
		fmt.Println(html)
		t.Errorf("Did not get changed content in div")
	}
}

const (
	testProxy = "socks4://171.103.9.22:4145"
)

func TestGetPageProxy(t *testing.T) {
	// Parse proxy
	URL, err := url.Parse(testProxy)
	if err != nil {
		t.Error(err)
	}

	// Try without proxy
	html, err := GetPageProxy(context.Background(), "http://ip4.me", nil, nil, time.Second*3, "")
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("out.html", []byte(html), 0644)

	// Check to make sure it DID NOT used our ip
	if strings.Index(html, URL.Hostname()) != -1 {
		t.Errorf("Did not use our ip address")
	}

	// Try with proxy
	html, err = GetPageProxy(context.Background(), "http://ip4.me", nil, nil, time.Second*3, testProxy)
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("out.html", []byte(html), 0644)

	// Check to make sure it DID use our ip
	if strings.Index(html, URL.Hostname()) == -1 {
		t.Errorf("Did not use our ip address")
	}
}
