package request

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
)

type Response struct {
	*http.Response
	content []byte
}

// Get Response Body as simplejson.Json
func (resp *Response) Json() (*simplejson.Json, error) {
	b, err := resp.Content()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(b)
}

// Get Response Body as []byte
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

// Get Response Body as string
func (resp *Response) Text() (string, error) {
	b, err := resp.Content()
	s := string(b)
	return s, err
}

// Does Response StatusCode < 400 ?
func (resp *Response) OK() bool {
	return resp.StatusCode < 400
}

// Does Response StatusCode < 400 ?
func (resp *Response) Ok() bool {
	return resp.OK()
}

// Get Response Status
func (resp *Response) Reason() string {
	return resp.Status
}

func (resp *Response) URL() (*url.URL, error) {
	u := resp.Request.URL
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound,
		http.StatusSeeOther, http.StatusTemporaryRedirect:
		location, err := resp.Location()
		if err != nil {
			return nil, err
		}
		u = u.ResolveReference(location)
	}
	return u, nil
}
