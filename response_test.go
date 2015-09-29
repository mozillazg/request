package request

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestResponseURL(t *testing.T) {
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
	url = "http://httpbin.org/redirect/3"

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}
	resp, _ = req.Get(url)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/relative-redirect/2")

}
