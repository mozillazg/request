package request

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
import (
	"github.com/bitly/go-simplejson"
)

type Request struct {
	*http.Request
}

type Response struct {
	*http.Response
	content []byte
}

func (resp *Response) Json() (*simplejson.Json, error) {
	b, err := resp.Content()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(b)
}

func (resp *Response) Content() (b []byte, err error) {
	if resp.content != nil {
		return resp.content, nil
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		b, err = ioutil.ReadAll(reader)
		defer reader.Close()
	default:
		reader = resp.Body
		b, err = ioutil.ReadAll(reader)
	}

	if err != nil {
		return nil, err
	}
	resp.content = b
	return b, err
}

func (resp *Response) Ok() bool {
	return resp.StatusCode < 400
}

func (resp *Response) Reason() string {
	return resp.Status
}

type Args struct {
	Client  *http.Client
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

func NewArgs(c *http.Client) *Args {
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
	resp = &Response{s, nil}
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
