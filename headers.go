package request

import "net/http"

var defaultUserAgent = "go-request/" + Version
var defaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      defaultUserAgent,
}
var defaultContentType = "application/x-www-form-urlencoded; charset=utf-8"
var defaultJsonType = "application/json; charset=utf-8"

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
