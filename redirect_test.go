package request

import (
	"errors"
	"testing"

	"github.com/bmizerany/assert"
	"net/http"
)

func TestCheckRedirectNoRedirect(t *testing.T) {
	req := NewRequest(nil)
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)
}

func TestCheckRedirectNumberLessThanDefault(t *testing.T) {
	req := NewRequest(nil)
	url := "https://httpbin.org/redirect/3"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/get")
}

func TestCheckRedirectNumberGreatThanDefault(t *testing.T) {
	req := NewRequest(nil)
	url := "https://httpbin.org/redirect/15"
	resp, err := req.Get(url)
	u, _ := resp.URL()
	assert.NotEqual(t, err, ErrMaxRedirect)
	assert.Equal(t, u.String(), "https://httpbin.org/relative-redirect/4")
}

func TestCheckRedirectWithHeaders(t *testing.T) {
	req := NewRequest(nil)
	url := "https://httpbin.org/redirect/2"
	req.Headers = map[string]string{
		"Referer": "http://example.com",
		"X-Test":  "test",
	}
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/get")
	assert.Equal(t, resp.Request.Header.Get("X-Test"), req.Headers["X-Test"])
	assert.Equal(t, resp.Request.Header.Get("Referer") != req.Headers["Referer"], true)
	assert.Equal(t, resp.Request.Header.Get("User-Agent"), DefaultUserAgent)
}

func TestCheckRedirectCustom(t *testing.T) {
	url := "https://httpbin.org/redirect/12"
	req := NewRequest(nil)
	req.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 16 {
			return errors.New("redirect")
		}
		return nil
	}
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/get")
}

func TestCheckRedirectChangeDefaultLimit(t *testing.T) {
	url := "https://httpbin.org/redirect/12"
	req := NewRequest(nil)
	origin := DefaultRedirectLimit
	DefaultRedirectLimit = 16
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/get")
	DefaultRedirectLimit = origin
}
