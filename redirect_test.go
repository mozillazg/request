package request

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestCheckRedirect(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)

	url = "http://httpbin.org/redirect/3"
	resp, _ = req.Get(url)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/get")

	url = "http://httpbin.org/redirect/15"
	resp, _ = req.Get(url)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/relative-redirect/4")

	url = "http://httpbin.org/redirect/2"
	req.Headers = map[string]string{
		"Referer": "http://example.com",
	}
	resp, _ = req.Get(url)
	u, _ = resp.URL()
	referer := resp.Request.Header.Get("Referer")
	assert.Equal(t, u.String(), "http://httpbin.org/get")
	assert.Equal(t, referer, "http://httpbin.org/relative-redirect/1")
	assert.Equal(t, resp.Request.Header.Get("User-Agent"), DefaultUserAgent)

	url = "http://httpbin.org/redirect/12"
	DefaultRedirectLimit = 16
	resp, _ = req.Get(url)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/get")
}
