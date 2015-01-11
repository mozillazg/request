request
=======
[![Build Status](https://travis-ci.org/mozillazg/request.svg?branch=master)](https://travis-ci.org/mozillazg/request)
[![Coverage Status](https://coveralls.io/repos/mozillazg/request/badge.png?branch=master)](https://coveralls.io/r/mozillazg/request?branch=master)
[![GoDoc](https://godoc.org/github.com/mozillazg/request?status.svg)](https://godoc.org/github.com/mozillazg/request)

Go HTTP Requests for Humans™. Inspired by [Python-Requests](https://github.com/kennethreitz/requests).


Installation
------------

```
go get -u github.com/mozillazg/request
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/request


Usage
-------

```go
import (
	"github.com/mozillazg/request"
)
```

**GET**:

```go
c := &http.Client{}
a := request.NewArgs(c)
resp, err := request.Get("http://httpbin.org/get", a)
j, err := resp.Json()
defer resp.Body.Close()  // Don't forget close the response body
```

**POST**:

```go
a.Data = map[string]string{
	"key": "value",
	"a":   "123",
}
resp, err := request.Post("http://httpbin.org/post", a)
```

**Cookies**:

```go
a.Cookies = map[string]string{
	"key": "value",
	"a":   "123",
}
resp, err := request.Get("http://httpbin.org/cookies", a)
```

**Headers**:

```go
a.Headers = map[string]string{
	"Accept-Encoding": "gzip,deflate,sdch",
	"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
}
resp, err := request.Get("http://httpbin.org/get", a)
```

**Files**:

```go
f, err := os.Open("test.txt")
a.Files = []request.FileField{
	request.FileField{"file", "test.txt", f},
}
resp, err := request.Post("http://httpbin.org/post", a)
```

**Json**:

```go
a.Json = map[string]string{
	"a": "A",
	"b": "B",
}
resp, err := request.Post("http://httpbin.org/post", a)
a.Json = []int{1, 2, 3}
resp, err = request.Post("http://httpbin.org/post", a)
```

**Proxy**:
```go
a.Proxy = "http://127.0.0.1:8080"
// a.Proxy = "https://127.0.0.1:8080"
// a.Proxy = "socks5://127.0.0.1:57341"
resp, err := request.Get("http://httpbin.org/get", a)
```
or https://github.com/mozillazg/request/tree/develop/_example/proxy

**HTTP Basic Authentication**:
```go
a.BasicAuth = request.BasicAuth{"user", "passwd"}
resp, err := request.Get("http://httpbin.org/basic-auth/user/passwd", a)
```
