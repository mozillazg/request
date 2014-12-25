package main

import (
	"fmt"
	"github.com/mozillazg/request"
)

func get(c *request.Client, a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
}

func head(c *request.Client, a *request.Args) {
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

func json(c *request.Client, a *request.Args) {
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

func main() {
	c := &request.Client{}
	a := request.NewArgs(c)

	// get(c, a)
	// head(c, a)
	json(c, a)

}
