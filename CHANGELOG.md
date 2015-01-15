# Changelog


## 0.2.0 (2015-01-15)

### API Changes

* Add `Args.Proxy` for support HTTP/HTTPS/SOCKS5 proxy
* Add `Args.BasicAuth` for Support HTTP Basic Authentication
* Add `func (resp *Response) URL() (*url.URL, error)`
* Add `var DefaultRedirectLimit = 10`

### Bugfixes

* Fix "http.Client don't use original Header when it do redirect" [#6](https://github.com/mozillazg/request/issues/6)


## 0.1.0 (2015-01-08)

* Initial Release
