package request_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mozillazg/request"
)

func ExampleGet() {
	c := new(http.Client)
	req := request.NewRequest(c)
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("url").MustString())
	// Output:
	//true
	//http://httpbin.org/get
}

func ExampleGet_params() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Params = map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(d.Get("url").MustString())
	// Output:
	//http://httpbin.org/get?a=1&b=2
}

func ExampleGet_customHeaders() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"X-Abc":      "abc",
		"User-Agent": "go-request-test",
	}
	url := "http://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(d.Get("headers").Get("User-Agent").MustString())
	fmt.Println(d.Get("headers").Get("X-Abc").MustString())
	// Output:
	//go-request-test
	//abc
}

func ExamplePost() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Data = map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()
}

func ExampleGet_cookies() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	url := "http://httpbin.org/cookies"
	resp, _ := req.Get(url)
	defer resp.Body.Close()
}

func ExamplePost_files() {
	c := new(http.Client)
	req := request.NewRequest(c)
	f, _ := os.Open("test.txt")
	defer f.Close()
	req.Files = []request.FileField{
		request.FileField{"abc", "abc.txt", f},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()
}
