package request

import (
	"bufio"
	"bytes"
	"net/http"
	"os"
	"reflect"
	"strings"
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
	f2, _ := os.Open("doc.go")
	defer f2.Close()
	req.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	req.Files = []FileField{
		{"abc", "abc.txt", b},
		{"test", "test.txt", f2},
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

func TestPostRawBody(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	req.Body = strings.NewReader("a=1&b=2")
	req.Headers = map[string]string{
		"Content-Type": DefaultContentType,
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()

	j, _ := resp.Json()
	assert.Equal(t, j.Get("form").MustMap(),
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostXML(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	xml := "<xml><a>abc</a></xml"
	req.Body = strings.NewReader(xml)
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()

	j, _ := resp.Json()
	data, _ := j.Get("data").String()
	assert.Equal(t, data, xml)
}

func TestPostFormIO(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	body := strings.NewReader("a=1&b=2")
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, body)
	defer resp.Body.Close()

	j, _ := resp.Json()
	assert.Equal(t, j.Get("form").MustMap(),
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostFormString(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	s := "a=1&b=2"
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, s)
	defer resp.Body.Close()

	j, _ := resp.Json()
	assert.Equal(t, j.Get("form").MustMap(),
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostFormStructA(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	s := map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, s)
	defer resp.Body.Close()

	j, _ := resp.Json()
	assert.Equal(t, j.Get("form").MustMap(),
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostFormStructB(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	s := map[string][]string{
		"a": {"1", "2"},
		"b": {"2", "3"},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, s)
	defer resp.Body.Close()

	j, _ := resp.Json()
	form := map[string][]string{}
	for k, v := range j.Get("form").MustMap() {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				tmp := s.Index(i).Interface().(string)
				form[k] = append(form[k], tmp)
			}
		}
	}
	assert.Equal(t, form,
		map[string][]string{
			"a": {"1", "2"},
			"b": {"2", "3"},
		}, true)
}

func TestPostFormFileA(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	f := []byte{'a', 'b', 'c', 'd'}
	_, _ = w.Write(f)
	w.Flush()

	req.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	req.Files = []FileField{
		{"abc", "abc.txt", b},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, nil)
	d, _ := resp.Json()
	defer resp.Body.Close()

	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("form").MustMap(), v)
	_, x := d.Get("files").CheckGet("abc")
	assert.Equal(t, x, true)
}

func TestPostFormFileB(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	f := []byte{'a', 'b', 'c', 'd'}
	_, _ = w.Write(f)
	w.Flush()

	data := map[string]string{
		"key": "value",
		"a":   "123",
	}
	req.Files = []FileField{
		{"abc", "abc.txt", b},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.PostForm(url, data)
	d, _ := resp.Json()
	defer resp.Body.Close()

	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("form").MustMap(), v)
	_, x := d.Get("files").CheckGet("abc")
	assert.Equal(t, x, true)
}
