package main

import (
	"fmt"
	"github.com/mozillazg/request"
	"net/http"
)

func get(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
}

func head(a *request.Args) {
	resp, err := request.Head("http://httpbin.org/get", a)
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	defer resp.Body.Close()
}

func json(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	if err != nil {
		return
	}

	d, err := resp.Json()
	if err != nil {
		return
	}
	fmt.Println(d.Get("headers").Get("User-Agent"))
	defer resp.Body.Close()
}

func gzip(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/gzip", a)
	if err != nil {
		return
	}

	d, err := resp.Json()
	if err != nil {
		return
	}
	fmt.Println(d.Get("headers").Get("Accept-Encoding"))
	fmt.Println(resp.Header.Get("Content-Encoding"))
	s, err := resp.Text()
	fmt.Println(s)
	defer resp.Body.Close()
}

func post(a *request.Args) {
	resp, err := request.Post("http://httpbin.org/post", a)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(resp.Ok())
	d, err := resp.Json()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func customHeaders(a *request.Args) {
	a.Headers = map[string]string{
		"Accept-Encoding": "gzip,deflate,sdch",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	}
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
	}
	d, err := resp.Json()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(d.Get("headers").Get("User-Agent"))
	fmt.Println(d.Get("headers").Get("Accept-Encoding"))
	fmt.Println(d.Get("headers").Get("Accept"))
}

func main() {
	c := &http.Client{}
	a := request.NewArgs(c)

	// fmt.Println("=====GET: ")
	// get(a)
	// fmt.Println("=====HEAD: ")
	// head(a)
	// fmt.Println("=====JSON: ")
	// json(a)
	// fmt.Println("=====GZIP: ")
	// gzip(a)

	// a.Data = map[string]string{
	// 	"key": "value",
	// 	"a":   "123",
	// }
	// fmt.Println("=====POST: ")
	// post(a)

	fmt.Println("=====Custom Headers: ")
	customHeaders(a)

}
