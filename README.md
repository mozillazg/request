request
=======
[![Build Status](https://travis-ci.org/mozillazg/request.svg?branch=master)](https://travis-ci.org/mozillazg/request)
[![Coverage Status](https://coveralls.io/repos/mozillazg/request/badge.png?branch=master)](https://coveralls.io/r/mozillazg/request?branch=master)
[![GoDoc](https://godoc.org/github.com/mozillazg/request?status.svg)](https://godoc.org/github.com/mozillazg/request)

Go HTTP Requests for Humans™. Inspired by [Python-Requests](https://github.com/kennethreitz/requests).


Usage
-------

**GET**:

```go
c := &http.Client{}
a := request.NewArgs(c)
resp, err := request.Get("http://httpbin.org/get", a)
j, err := resp.Json()
```

**POST**:

```go
c := &http.Client{}
a := request.NewArgs(c)
a.Data = map[string]string{
	"key": "value",
	"a":   "123",
}
resp, err := request.Post("http://httpbin.org/post", a)
j, err := resp.Json()
```

**Cookies**:

```go
c := &http.Client{}
a := request.NewArgs(c)
a.Cookies = map[string]string{
	"key": "value",
	"a":   "123",
}
resp, err := request.Get("http://httpbin.org/cookies", a)
j, err := resp.Json()
```

**Headers**:

```go
c := &http.Client{}
a := request.NewArgs(c)
a.Headers = map[string]string{
  "Accept-Encoding": "gzip,deflate,sdch",
  "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8", 
}
resp, err := request.Get("http://httpbin.org/gzip", a)
j, err := resp.Json()
```
