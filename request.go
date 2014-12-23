package request

import (
	// "io"
	"net/http"
	// "net/url"
)

type Request http.Request

type Response http.Response

func (resp *Response) OK() bool {
	return resp.StatusCode < 400
}

func (resp *Response) Reason() string {
	return resp.Status
}

type Args struct {
	Headers map[string]string
	Cookies map[string]string
	Data    map[string]string
	Params  map[string]string
	Files   map[string]string
}

var headers = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      "go-request/0.1.0",
}

func NewArgs() *Args {
	return &Args{
		Headers: headers,
	}
}

func newRequest(method string, url string, a *Args) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	for k, v := range a.Headers {
		req.Header.Set(k, v)
	}
	resp, err = client.Do(req)
	return
}

func Get(url string, a *Args) (resp *Response, err error) {
	s, err := newRequest("GET", url, a)
	resp = (*Response)(s)
	return
}

func Head(url string, a *Args) (resp *Response, err error) {
	s, err := newRequest("HEAD", url, a)
	resp = (*Response)(s)
	return
}
