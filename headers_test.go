package request

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestHeaders(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/get"
	req.Headers = map[string]string{
		"Foo": "Bar",
	}
	resp, _ := req.Get(url)
	j, _ := resp.Json()
	defer resp.Body.Close()
	v, _ := j.Get("headers").MustMap()["Foo"]
	assert.Equal(t, v, "Bar")
	v, _ = j.Get("headers").MustMap()["User-Agent"]
	assert.Equal(t, v, DefaultUserAgent)

	req.Headers = map[string]string{
		"User-Agent": "Foobar",
	}
	resp, _ = req.Get(url)
	j, _ = resp.Json()
	defer resp.Body.Close()
	v, _ = j.Get("headers").MustMap()["User-Agent"]
	assert.Equal(t, v, "Foobar")
}
