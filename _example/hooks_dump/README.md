
```
$ go run _example/hooks_dump/dump.go
POST /post HTTP/1.1
Host: httpbin.org
User-Agent: go-request/0.7.0
Content-Length: 7
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Type: application/json; charset=utf-8

[1,2,3]

HTTP/1.1 200 OK
Content-Length: 413
Access-Control-Allow-Credentials: true
Access-Control-Allow-Origin: *
Connection: keep-alive
Content-Type: application/json
Date: Sun, 11 Dec 2016 13:48:19 GMT
Server: nginx

{
  "args": {},
  "data": "[1,2,3]",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip, deflate",
    "Content-Length": "7",
    "Content-Type": "application/json; charset=utf-8",
    "Host": "httpbin.org",
    "User-Agent": "go-request/0.7.0"
  },
  "json": [
    1,
    2,
    3
  ],
  "origin": "192.168.1.22",
  "url": "https://httpbin.org/post"
}

```
