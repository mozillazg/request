package request_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mozillazg/request"
)

func ExampleGet() {
	c := &http.Client{}
	a := request.NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := request.Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("url").MustString())
	// Output:
	//true
	//http://httpbin.org/get
}

func ExampleGet_params() {
	c := &http.Client{}
	a := request.NewArgs(c)
	a.Params = map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/get"
	resp, _ := request.Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(d.Get("url").MustString())
	// Output:
	//http://httpbin.org/get?a=1&b=2
}

func ExampleGet_custom_headers() {
	c := &http.Client{}
	a := request.NewArgs(c)
	a.Headers = map[string]string{
		"X-Abc":      "abc",
		"User-Agent": "go-request-test",
	}
	url := "http://httpbin.org/get"
	resp, _ := request.Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(d.Get("headers").Get("User-Agent").MustString())
	fmt.Println(d.Get("headers").Get("X-Abc").MustString())
	// Output:
	//go-request-test
	//abc
}

func ExamplePost() {
	c := &http.Client{}
	a := request.NewArgs(c)
	a.Data = map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/post"
	resp, _ := request.Post(url, a)
	defer resp.Body.Close()
}

func ExampleGet_cookies() {
	c := &http.Client{}
	a := request.NewArgs(c)
	a.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	url := "http://httpbin.org/cookies"
	resp, _ := request.Get(url, a)
	defer resp.Body.Close()
}

func ExamplePost_files() {
	c := &http.Client{}
	a := request.NewArgs(c)
	f, _ := os.Open("test.txt")
	a.Files = []request.FileField{
		request.FileField{"abc", "abc.txt", f},
	}
	url := "http://httpbin.org/post"
	resp, _ := request.Post(url, a)
	defer resp.Body.Close()
}
