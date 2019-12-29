package simplebrowser

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestChrome(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(HTMLTest))
	}))
	defer ts.Close()

	html, err := GetPage(context.Background(), ts.URL, nil, map[string]interface{}{}, time.Second)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if strings.Index(html, `<div id="content">test</div>`) == -1 {
		fmt.Println(html)
		t.Errorf("Did not get changed content in div")
	}

	// TODO: Test with cookies and headers
}
