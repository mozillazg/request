package request

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	http.Client
}

type Request struct {
	*http.Request
}

type Response struct {
	*http.Response
}

func (resp *Response) Json() (*simplejson.Json, error) {
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(s)
}

func (resp *Response) Ok() bool {
	return resp.StatusCode < 400
}

func (resp *Response) Reason() string {
	return resp.Status
}

type Args struct {
	Client  *Client
	Headers map[string]string
	Cookies map[string]string
	Data    map[string]string
	Params  map[string]string
	Files   map[string]string
}

var defaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      "go-request/0.1.0",
}

func NewArgs(c *Client) *Args {
	return &Args{
		Client:  c,
		Headers: defaultHeaders,
	}
}

func newRequest(method string, url string, a *Args) (resp *Response, err error) {
	client := a.Client
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	for k, v := range a.Headers {
		req.Header.Set(k, v)
	}
	s, err := client.Do(req)
	resp = &Response{s}
	return
}

func Get(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("GET", url, a)
	return
}

func Head(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("HEAD", url, a)
	return
}
