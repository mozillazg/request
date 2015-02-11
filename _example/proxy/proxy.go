package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mozillazg/request"
)

func httpProxy(URL string) {
	proxyURL, _ := url.Parse(URL)
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	c := &http.Client{Transport: tr}
	req := request.NewRequest(c)
	resp, err := req.Get("http://httpbin.org/get")
	fmt.Println(err)
	fmt.Println(resp.Text())
}

func httpProxy2(URL string) {
	c := &http.Client{}
	req := request.NewRequest(c)
	req.Proxy = URL
	resp, err := req.Get("http://httpbin.org/get")
	fmt.Println(err)
	fmt.Println(resp.Text())
}

func main() {
	// c := new(http.Client)
	// req := request.NewRequest(c)
	// req.Get("http://httpbin.org/get")
	// httpProxy("http://64.31.22.131:8089")
	// httpProxy("https://64.31.22.131:8089")
	// httpProxy2("http://64.31.22.131:8089")
	// httpProxy2("https://64.31.22.131:8089")
	// httpProxy2("socks5://210.38.111.249:1080")
}
