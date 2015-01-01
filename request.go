package request

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/bitly/go-simplejson"
	"golang.org/x/net/publicsuffix"
)

const Version = "0.1.0"

// type Request struct {
// 	*http.Request
// }

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
			return nil, err
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

func (resp *Response) Text() (string, error) {
	b, err := resp.Content()
	s := string(b)
	return s, err
}

func (resp *Response) Ok() bool {
	return resp.StatusCode < 400
}

func (resp *Response) Reason() string {
	return resp.Status
}

type FileField struct {
	FieldName string
	FileName  string
	File      io.Reader
}

type Args struct {
	Client  *http.Client
	Headers map[string]string
	Cookies map[string]string
	Data    map[string]string
	Params  map[string]string
	Files   []FileField
	Json    interface{}
}

var defaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      "go-request/" + Version,
}
var defaultContentType = "application/x-www-form-urlencoded; charset=utf-8"
var defaultJsonType = "application/json; charset=utf-8"

func NewArgs(c *http.Client) *Args {
	if c.Jar == nil {
		options := cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		}
		jar, _ := cookiejar.New(&options)
		c.Jar = jar
	}

	return &Args{
		Client:  c,
		Headers: defaultHeaders,
		Cookies: nil,
		Data:    nil,
		Params:  nil,
		Files:   nil,
		Json:    nil,
	}
}

func applyHeaders(a *Args, req *http.Request, contentType string) {
	// apply defaultHeaders
	for k, v := range defaultHeaders {
		_, ok := a.Headers[k]
		if !ok {
			req.Header.Set(k, v)
		}
	}
	// apply custom Headers
	for k, v := range a.Headers {
		req.Header.Set(k, v)
	}
	// apply "Content-Type" Headers
	_, ok := a.Headers["Content-Type"]
	if !ok {
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		} else {
			req.Header.Set("Content-Type", defaultContentType)
		}
	}
}

func applyCookies(a *Args, req *http.Request) {
	if a.Cookies == nil {
		return
	}
	cookies := a.Client.Jar.Cookies(req.URL)
	for k, v := range a.Cookies {
		cookies = append(cookies, &http.Cookie{Name: k, Value: v})
	}
	a.Client.Jar.SetCookies(req.URL, cookies)
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

func newMultipartBody(a *Args) (body io.Reader, contentType string, err error) {
	files := a.Files
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	for _, file := range files {
		fileWriter, err := bodyWriter.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(fileWriter, file.File)
		if err != nil {
			return nil, "", err
		}
	}
	if a.Data != nil {
		for k, v := range a.Data {
			bodyWriter.WriteField(k, v)
		}
	}
	contentType = bodyWriter.FormDataContentType()
	defer bodyWriter.Close()
	body = bodyBuffer
	return
}

func newJsonBody(a *Args) (body io.Reader, contentType string, err error) {
	b, err := json.Marshal(a.Json)
	if err != nil {
		return nil, "", err
	}
	return bytes.NewReader(b), defaultJsonType, err
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
	client := a.Client
	body, contentType, err := newBody(a)
	u := newURL(url, a.Params)
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}
	applyHeaders(a, req, contentType)
	applyCookies(a, req)

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

func Post(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("POST", url, a)
	return
}

func Put(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("PUT", url, a)
	return
}

func Patch(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("PATCH", url, a)
	return
}

func Delete(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("DELETE", url, a)
	return
}

func Options(url string, a *Args) (resp *Response, err error) {
	resp, err = newRequest("OPTIONS", url, a)
	return
}
