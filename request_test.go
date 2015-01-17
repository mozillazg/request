package request

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/bmizerany/assert"
)

func TestGet(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
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

func TestGetParmas(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Params = map[string]string{
		"foo": "bar",
		"a":   "1",
	}
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("args").MustMap(),
		map[string]interface{}{
			"foo": "bar",
			"a":   "1",
		})
}

func TestGetParmas2(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Params = map[string]string{
		"foo": "bar",
		"a":   "1",
	}
	url := "http://httpbin.org/get?ab=cd"
	resp, _ := Get(url, a)
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
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Head(url, a)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestPost(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Data = map[string]string{
		"a":   "A",
		"foo": "bar",
	}
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
	assert.Equal(t, d.Get("form").MustMap(),
		map[string]interface{}{
			"a":   "A",
			"foo": "bar",
		}, true)
}

func TestPut(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/put"
	resp, _ := Put(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestDelete(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/delete"
	resp, _ := Delete(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestPatch(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/patch"
	resp, _ := Patch(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestOptions(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Options(url, a)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestPostJson(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Json = []int{1, 2, 3}
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
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
	c := &http.Client{}
	a := NewArgs(c)
	a.Json = map[string]string{
		"a":   "b",
		"foo": "bar",
	}
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
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
	c := &http.Client{}
	a := NewArgs(c)
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
	a.Json = d
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
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

func TestPostFiles(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	f := []byte{
		'a',
		'b',
		'c',
		'd',
	}
	_, _ = w.Write(f)
	w.Flush()
	f2, _ := os.Open("test.txt")
	defer f2.Close()
	a.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	a.Files = []FileField{
		FileField{"abc", "abc.txt", b},
		FileField{"test", "test.txt", f2},
	}
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("form").MustMap(), v)
	_, x := d.Get("files").CheckGet("abc")
	assert.Equal(t, x, true)
	_, x = d.Get("files").CheckGet("test")
	assert.Equal(t, x, true)
}

func TestGzip(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/gzip"
	resp, _ := Get(url, a)
	d, _ := resp.Json()
	t2, _ := resp.Text()
	c2, _ := resp.Content()
	defer resp.Body.Close()

	assert.Equal(t, resp.Reason() != "", true)
	assert.Equal(t, resp.OK(), true)
	assert.Equal(t, t2 != "", true)
	assert.Equal(t, c2 != nil, true)
	assert.Equal(t, d.Get("gzipped").MustBool(), true)
}

func currentIP(u string) (ip string) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Proxy = u
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	return d.Get("origin").MustString()
}
func currentIPHTTPS(u string) (ip string) {
	c := &http.Client{}
	a := NewArgs(c)
	a.Proxy = u
	url := "https://httpbin.org/get"
	resp, _ := Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	return d.Get("origin").MustString()
}

// func TestProxy(t *testing.T) {
// 	ip := currentIP("")
// 	httpProxyURL := os.Getenv("http_proxy_url")
// 	httpsProxyURL := os.Getenv("https_proxy_url")
// 	socks5ProxyURL := os.Getenv("socks5_proxy_url")
//
// 	assert.Equal(t, currentIP(httpProxyURL) != ip, true)
// 	assert.Equal(t, currentIP(httpsProxyURL) != ip, true)
// 	// assert.Equal(t, currentIPHTTPS(httpsProxyURL) != ip, true)
// 	assert.Equal(t, currentIP(socks5ProxyURL) != ip, true)
// }

func TestBasicAuth(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	a.BasicAuth = BasicAuth{"user", "passwd"}
	url := "http://httpbin.org/basic-auth/user/passwd"
	resp, _ := Get(url, a)
	assert.Equal(t, resp.OK(), true)

	a.BasicAuth = BasicAuth{
		Username: "user2",
		Password: "passwd2",
	}
	url = "http://httpbin.org/basic-auth/user2/passwd2"
	resp, _ = Get(url, a)
	assert.Equal(t, resp.OK(), true)
}

func TestResponseURL(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)

	url = "http://httpbin.org/redirect/3"
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/get")
	url = "http://httpbin.org/redirect/3"

	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/relative-redirect/2")

}

func TestCheckRedirect(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)

	url = "http://httpbin.org/redirect/3"
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/get")

	url = "http://httpbin.org/redirect/15"
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/relative-redirect/4")

	url = "http://httpbin.org/redirect/2"
	a.Headers = map[string]string{
		"Referer": "http://example.com",
	}
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	referer := resp.Request.Header.Get("Referer")
	assert.Equal(t, u.String(), "http://httpbin.org/get")
	assert.Equal(t, referer, "http://httpbin.org/relative-redirect/1")
	assert.Equal(t, resp.Request.Header.Get("User-Agent"), defaultUserAgent)

	url = "http://httpbin.org/redirect/12"
	DefaultRedirectLimit = 16
	resp, _ = Get(url, a)
	u, _ = resp.URL()
	assert.Equal(t, u.String(), "http://httpbin.org/get")
}
