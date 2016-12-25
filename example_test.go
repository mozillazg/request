package request_test

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mozillazg/request"
)

func ExampleRequest_Get() {
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

func ExampleGet() {
	c := new(http.Client)
	args := request.NewArgs(c)
	url := "http://httpbin.org/get"
	resp, _ := request.Get(url, args)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("url").MustString())
	// Output:
	//true
	//http://httpbin.org/get
}

func ExampleRequest_Get_params() {
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

func ExampleRequest_Get_customHeaders() {
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

func ExampleRequest_Post() {
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

func ExamplePost() {
	c := new(http.Client)
	args := request.NewArgs(c)
	args.Data = map[string]string{
		"a": "1",
		"b": "2",
	}
	url := "http://httpbin.org/post"
	resp, _ := request.Post(url, args)
	defer resp.Body.Close()
}

func ExampleRequest_Get_cookies() {
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

func ExampleRequest_Post_files() {
	c := new(http.Client)
	req := request.NewRequest(c)
	f, _ := os.Open("test.txt")
	defer f.Close()
	req.Files = []request.FileField{
		{FieldName: "abc", FileName: "abc.txt", File: f},
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()
}

func ExampleRequest_Post_rawBody() {
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Body = strings.NewReader("a=1&b=2&foo=bar")
	req.Headers = map[string]string{
		"Content-Type": request.DefaultContentType,
	}
	url := "http://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()
}
