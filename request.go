package request

import (
	// "io"
	"net/http"
	// "net/url"
)

type Request http.Request

type Response http.Response

func (resp *Response) Ok() bool {
	return resp.StatusCode < 400
}

func (resp *Response) Reason() string {
	return resp.Status
}

func Get(url string) (resp *Response, err error) {
	s, err := http.Get(url)
	resp = (*Response)(s)
	return
}

func Head(url string) (resp *Response, err error) {
	s, err := http.Head(url)
	resp = (*Response)(s)
	return
}
