package simplebrowser

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func ExamplePageRequest() {
	var html string
	_ = NewPageRequest("https://google.com").WithHTMLGet(&html).Do(context.Background())
	fmt.Println(html)
}

func ExamplePageRequest_withContent() {
	// Create testing server
	var HTMLExample = `<html>
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(HTMLExample))
	}))
	defer ts.Close()

	// Get page and wait for javascript to change div
	var html string
	err := NewPageRequest(ts.URL).WithWaitTime(time.Second).WithHTMLGet(&html).Do(context.Background())
	if err != nil {
		return
	}

	fmt.Println(strings.Index(html, `<div id="content">test</div>`))

	// Output: 185
}
