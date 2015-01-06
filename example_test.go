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

func ExampleGetparams() {
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
	fmt.Println(d.Get("args").MustMap())
	// Output:
	//http://httpbin.org/get?a=1&b=2
	//map[a:1 b:2]
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
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("form").MustMap())
	// Output:
	//true
	//map[a:1 b:2]
}

func Examplecookies() {
	c := &http.Client{}
	a := request.NewArgs(c)
	a.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	url := "http://httpbin.org/cookies"
	resp, _ := request.Get(url, a)
	d, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("cookies").MustMap())
	// Output:
	//true
	//map[name:value foo:bar]
}

func Examplefiles() {
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
