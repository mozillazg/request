// Go HTTP Requests for Humans™.
//
// HTTP Request is so easy:
//
// GET Request
//
//
// 	c := &http.Client{}
// 	a := request.NewArgs(c)
// 	resp, err := request.Get("http://httpbin.org/get", a)
// 	j, err := resp.Json()
// 	defer resp.Body.Close()  // Don't forget close the response body
//
// POST Request
//
//	a.Data = map[string]string{
//		"key": "value",
//		"a":   "123",
//	}
//	resp, err := request.Post("http://httpbin.org/post", a)
//
// Custom Cookies
//
//	a.Cookies = map[string]string{
//		"key": "value",
//		"a":   "123",
//	}
//	resp, err := request.Get("http://httpbin.org/cookies", a)
//
//
// Custom Headers
//
//	a.Headers = map[string]string{
//		"Accept-Encoding": "gzip,deflate,sdch",
//		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
//	}
//	resp, err := request.Get("http://httpbin.org/get", a)
//
// Upload Files
//
//	f, err := os.Open("test.txt")
//	a.Files = []request.FileField{
//		request.FileField{"file", "test.txt", f},
//	}
//	resp, err := request.Post("http://httpbin.org/post", a)
//
// Json Body
//
//	a.Json = map[string]string{
//		"a": "A",
//		"b": "B",
//	}
//	resp, err := request.Post("http://httpbin.org/post", a)
//	a.Json = []int{1, 2, 3}
//	resp, err = request.Post("http://httpbin.org/post", a)
//
// Proxy
//
//	a.Proxy = "http://127.0.0.1:8080"
//	// a.Proxy = "https://127.0.0.1:8080"
//	// a.Proxy = "socks5://127.0.0.1:57341"
//	resp, err := request.Post("http://httpbin.org/get", a)
package request
