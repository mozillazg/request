package request

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const Version = "0.3.0"

type FileField struct {
	FieldName string
	FileName  string
	File      io.Reader
}

type BasicAuth struct {
	Username string
	Password string
}

type Args struct {
	Client    *http.Client
	Headers   map[string]string
	Cookies   map[string]string
	Data      map[string]string
	Params    map[string]string
	Files     []FileField
	Json      interface{}
	Proxy     string
	BasicAuth BasicAuth
}

type Request struct {
	*Args
}

func NewArgs(c *http.Client) *Args {
	if c.Jar == nil {
		options := cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		}
		jar, _ := cookiejar.New(&options)
		c.Jar = jar
	}

	return &Args{
		Client:    c,
		Headers:   defaultHeaders,
		Cookies:   nil,
		Data:      nil,
		Params:    nil,
		Files:     nil,
		Json:      nil,
		Proxy:     "",
		BasicAuth: BasicAuth{},
	}
}

func NewRequest(c *http.Client) *Request {
	return &Request{NewArgs(c)}
}

func newURL(u string, params map[string]string) string {
	if params == nil {
		return u
	}

	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	if strings.Contains(u, "?") {
		return u + "&" + p.Encode()
	}
	return u + "?" + p.Encode()
}

func newBody(a *Args) (body io.Reader, contentType string, err error) {
	if a.Data == nil && a.Files == nil && a.Json == nil {
		return nil, "", nil
	}
	if a.Files != nil {
		return newMultipartBody(a)
	} else if a.Json != nil {
		return newJsonBody(a)
	}

	d := url.Values{}
	for k, v := range a.Data {
		d.Set(k, v)
	}
	return strings.NewReader(d.Encode()), "", nil
}

func newRequest(method string, url string, a *Args) (resp *Response, err error) {
	body, contentType, err := newBody(a)
	u := newURL(url, a.Params)
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	applyHeaders(a, req, contentType)
	applyCookies(a, req)
	applyProxy(a)
	applyCheckRdirect(a)

	if a.BasicAuth.Username != "" {
		req.SetBasicAuth(a.BasicAuth.Username, a.BasicAuth.Password)
	}

	s, err := a.Client.Do(req)
	resp = &Response{s, nil}
	return
}

// Get issues a GET to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Get(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("GET", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Get(url interface{}) (resp *Response, err error) {
	resp, err = Get(url2string(url), req2arg(req))
	return
}

// Head issues a HEAD to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Head(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("HEAD", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Head(url interface{}) (resp *Response, err error) {
	resp, err = Head(url2string(url), req2arg(req))
	return
}

// Post issues a POST to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Post(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("POST", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Post(url interface{}) (resp *Response, err error) {
	resp, err = Post(url2string(url), req2arg(req))
	return
}

// Put issues a PUT to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Put(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("PUT", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Put(url interface{}) (resp *Response, err error) {
	resp, err = Put(url2string(url), req2arg(req))
	return
}

// Patch issues a PATCH to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Patch(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("PATCH", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Patch(url interface{}) (resp *Response, err error) {
	resp, err = Patch(url2string(url), req2arg(req))
	return
}

// Delete issues a DELETE to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Delete(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("DELETE", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Delete(url interface{}) (resp *Response, err error) {
	resp, err = Delete(url2string(url), req2arg(req))
	return
}

// Options issues a OPTIONS to the specified URL.
//
// Caller should close resp.Body when done reading from it.
func Options(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("OPTIONS", url, a)
	return
}

// url can be string or *url.URL or ur.URL
func (req *Request) Options(url interface{}) (resp *Response, err error) {
	resp, err = Options(url2string(url), req2arg(req))
	return
}

func url2string(u interface{}) string {
	switch u.(type) {
	case string:
		return u.(string)
	case url.URL:
		s := u.(url.URL)
		return s.String()
	case *url.URL:
		s := u.(*url.URL)
		return s.String()
	}
	return ""
}

func req2arg(req *Request) (a *Args) {
	return &Args{
		Client:    req.Client,
		Headers:   req.Headers,
		Cookies:   req.Cookies,
		Data:      req.Data,
		Params:    req.Params,
		Files:     req.Files,
		Json:      req.Json,
		Proxy:     req.Proxy,
		BasicAuth: req.BasicAuth,
	}
}
