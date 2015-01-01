package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/mozillazg/request"
)

func get(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
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
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
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

func put(a *request.Args) {
	resp, _ := request.Put("http://httpbin.org/put", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func patch(a *request.Args) {
	resp, _ := request.Patch("http://httpbin.org/patch", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func deleteF(a *request.Args) {
	resp, _ := request.Delete("http://httpbin.org/delete", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("headers").Get("Content-Type"))
	fmt.Println(d.Get("form"))
}

func options(a *request.Args) {
	resp, _ := request.Options("http://httpbin.org/get", a)
	defer resp.Body.Close()

	fmt.Println(resp.Ok())
	fmt.Println("Allow:", resp.Header.Get("Allow"))
}

func getParams(a *request.Args) {
	resp, err := request.Get("http://httpbin.org/get?foobar=123", a)
	defer resp.Body.Close()
	if err == nil {
		fmt.Println(resp.Ok())
		fmt.Println(resp.Reason())
	}
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("args"))
}

func cookies(a *request.Args) {
	resp, _ := request.Get("http://httpbin.org/cookies", a)
	defer resp.Body.Close()
	d, _ := resp.Json()
	fmt.Println(d.Get("cookies"))
	fmt.Println(resp.Cookies())
	url := resp.Request.URL
	fmt.Println(a.Client.Jar.Cookies(url))
}

func file(a *request.Args) {
	resp, _ := request.Post("http://httpbin.org/post", a)
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("files"))
	fmt.Println(d.Get("form"))
}

func files(a *request.Args) {
	resp, _ := request.Post("http://httpbin.org/post", a)
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("url"))
	fmt.Println(d.Get("files"))
}

func jsonPost(a *request.Args) {
	resp, _ := request.Post("http://httpbin.org/post", a)
	defer resp.Body.Close()
	fmt.Println(resp.Ok())
	d, _ := resp.Json()
	fmt.Println(d.Get("json"))
}

func main() {
	c := &http.Client{}
	a := request.NewArgs(c)

	fmt.Println("=====GET: ")
	get(a)
	fmt.Println("=====HEAD: ")
	head(a)
	fmt.Println("=====JSON: ")
	json(a)
	fmt.Println("=====GZIP: ")
	gzip(a)

	a.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	fmt.Println("=====POST: ")
	post(a)

	fmt.Println("=====Custom Headers: ")
	customHeaders(a)

	a = request.NewArgs(c)
	a.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	fmt.Println("=====PUT: ")
	put(a)
	fmt.Println("=====PATCH: ")
	patch(a)
	fmt.Println("=====DELTE: ")
	deleteF(a)
	fmt.Println("=====OPTIONS: ")
	options(a)

	a = request.NewArgs(c)
	a.Params = map[string]string{
		"a":   "abc",
		"key": "value",
	}
	fmt.Println("=====Params: ")
	get(a)
	getParams(a)
	post(a)

	fmt.Println("=====Cookies: ")
	a = request.NewArgs(c)
	cookies(a)
	a.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	cookies(a)

	fmt.Println("=====File: ")
	a = request.NewArgs(c)
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

	a.Data = map[string]string{
		"key": "value",
		"a":   "123",
	}
	a.Files = []request.FileField{
		request.FileField{"abc", "abc.txt", b},
		request.FileField{"test", "test.txt", f2},
	}
	file(a)
	f2, _ = os.Open("test.txt")
	a.Files = []request.FileField{
		request.FileField{"abc", "abc.txt", f2},
	}
	file(a)

	fmt.Println("=====JSON POST: ")
	type j struct {
		A string            `json:"a"`
		B map[string]string `json:"b"`
		C []string          `json:"c"`
		D []int             `json:"d"`
		E int               `json:"e"`
	}
	a = request.NewArgs(c)
	d := j{
		A: "hello",
		B: map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
		},
		C: []string{"lala", "aaaa"},
		D: []int{1, 2, 3},
		E: 5,
	}
	fmt.Println(d)
	a.Json = d
	jsonPost(a)
	a.Json = []int{1, 2, 3, 4}
	jsonPost(a)
}
