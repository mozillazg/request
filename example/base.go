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

type body struct {
	args    map[string]string
	headers map[string]string
	origin  string
	url     string
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

func main() {
	c := &http.Client{}
	a := request.NewArgs(c)

	// fmt.Println("=====GET: ")
	// get(a)
	// fmt.Println("=====HEAD: ")
	// head(a)
	// fmt.Println("=====JSON: ")
	// json(a)
	fmt.Println("=====GZIP: ")
	gzip(a)
}
