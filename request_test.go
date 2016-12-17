package request

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/bmizerany/assert"
)

func TestGet(t *testing.T) {
	c := &http.Client{}
	req := NewRequest(c)
	u2, _ := url.Parse("http://httpbin.org/get")
	resp, _ := req.Get(u2)
	assert.Equal(t, resp.Ok(), true)

	u3 := url.URL{
		Scheme: "http",
		Host:   "httpbin.org",
		Path:   "get",
	}
	resp, _ = req.Get(u3)
	assert.Equal(t, resp.Ok(), true)

	url := "http://httpbin.org/get"
	resp, _ = req.Get(url)
	d, _ := resp.Json()
	t2, _ := resp.Text()
	c2, _ := resp.Content()
	defer resp.Body.Close()

	assert.Equal(t, resp.Reason() != "", true)
	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, resp.OK(), true)
	assert.Equal(t, t2 != "", true)
	assert.Equal(t, c2 != nil, true)
	assert.Equal(t, d.Get("url").MustString(), url)

}

func TestGetParams(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Params = map[string]string{
		"foo": "bar",
		"a":   "1",
	}
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("args").MustMap(),
		map[string]interface{}{
			"foo": "bar",
			"a":   "1",
		})
}

func TestGetParams2(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Params = map[string]string{
		"foo": "bar",
		"a":   "1",
	}
	url := "http://httpbin.org/get?ab=cd"
	resp, _ := req.Get(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("args").MustMap(),
		map[string]interface{}{
			"ab":  "cd",
			"foo": "bar",
			"a":   "1",
		})
}

func TestHead(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/get"
	resp, _ := req.Head(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	content, _ := resp.Content()
	assert.Equal(t, content, []byte{})
}

func TestPut(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/put"
	resp, _ := req.Put(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestDelete(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/delete"
	resp, _ := req.Delete(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestPatch(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/patch"
	resp, _ := req.Patch(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestOptions(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/get"
	resp, _ := req.Options(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestPostJson(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Json = []int{1, 2, 3}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := []interface{}{
		json.Number("1"),
		json.Number("2"),
		json.Number("3"),
	}
	assert.Equal(t, d.Get("json").MustArray(), v)
}

func TestPostJson2(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Json = map[string]string{
		"a":   "b",
		"foo": "bar",
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := map[string]interface{}{
		"a":   "b",
		"foo": "bar",
	}
	assert.Equal(t, d.Get("json").MustMap(), v)
}

func TestPostJson3(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	type j struct {
		A string            `json:"a"`
		B map[string]string `json:"b"`
		C []string          `json:"c"`
		D []int             `json:"d"`
		E int               `json:"e"`
	}
	d := j{
		A: "hello",
		B: map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
		},
		C: []string{"lala", "aaaa"},
		D: []int{1, 2, 3},
		E: 5,
	}
	req.Json = d
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	j2, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v, _ := simplejson.NewJson([]byte(`{
		"a": "hello",
		"b": {"a": "A", "b":"B", "c":"C"},
		"c": ["lala", "aaaa"],
		"d": [1, 2, 3],
		"e": 5
	}`))
	assert.Equal(t, j2.Get("json"), v)
}

func TestBasicAuth(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.BasicAuth = BasicAuth{"user", "passwd"}
	url := "http://httpbin.org/basic-auth/user/passwd"
	resp, _ := req.Get(url)
	defer resp.Body.Close()
	assert.Equal(t, resp.OK(), true)

	req.BasicAuth = BasicAuth{
		Username: "user2",
		Password: "passwd2",
	}
	url = "http://httpbin.org/basic-auth/user2/passwd2"
	resp, _ = req.Get(url)
	defer resp.Body.Close()
	assert.Equal(t, resp.OK(), true)
}

func TestReset(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.BasicAuth = BasicAuth{"user", "passwd"}
	url := "http://httpbin.org"
	req.Get(url)

	req.Reset()
	assert.Equal(t, req.BasicAuth, BasicAuth{})
}

func TestGetArgsNilA(t *testing.T) {
	url := "http://httpbin.org/get"
	resp, _ := Get(url, nil)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestGetArgsNilB(t *testing.T) {
	args := NewArgs(nil)
	url := "http://httpbin.org/get"
	resp, _ := Get(url, args)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestGetArgsNilC(t *testing.T) {
	req := NewRequest(nil)
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}
