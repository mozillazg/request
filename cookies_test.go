package request

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestCookies(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Cookies = map[string]string{
		"key": "value",
		"a":   "123",
	}
	resp, _ := req.Get("http://httpbin.org/cookies")
	d, _ := resp.Json()
	defer resp.Body.Close()

	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("cookies").MustMap(), v)
}
