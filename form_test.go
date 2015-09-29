package request

import (
	"bufio"
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestPost(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Data = map[string]string{
		"a":   "A",
		"foo": "bar",
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
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

func TestPostFiles(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
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
	req.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	req.Files = []FileField{
		FileField{"abc", "abc.txt", b},
		FileField{"test", "test.txt", f2},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
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
