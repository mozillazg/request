package request

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGet(t *testing.T) {
	c := &http.Client{}
	a := NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
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
	url := "http://httpbin.org/post"
	resp, _ := Post(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
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
