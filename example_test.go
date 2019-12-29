package simplebrowser

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func ExampleGetPage() {
	html, _ := GetPage(context.Background(), "https://google.com", nil, map[string]interface{}{}, time.Second)
	fmt.Println(html)
}

func ExampleGetPage_withContent() {
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
	html, err := GetPage(context.Background(), ts.URL, nil, map[string]interface{}{}, time.Second)
	if err != nil {
		return
	}

	fmt.Println(strings.Index(html, `<div id="content">test</div>`))

	// Output: 185
}
