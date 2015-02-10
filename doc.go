// Go HTTP Requests for Humans™.
//
// HTTP Request is so easy:
//
// GET Request:
//
//
// 	c := &http.Client{}
// 	req := request.NewRequest(c)
// 	resp, err := req.Get("http://httpbin.org/get")
// 	j, err := resp.Json()
// 	defer resp.Body.Close()  // Don't forget close the response body
//
// POST Request:
//
//	req.Data = map[string]string{
//		"key": "value",
//		"a":   "123",
//	}
//	resp, err := req.Post("http://httpbin.org/post")
//
// Custom Cookies:
//
//	req.Cookies = map[string]string{
//		"key": "value",
//		"a":   "123",
//	}
//	resp, err := req.Get("http://httpbin.org/cookies")
//
//
// Custom Headers:
//
//	req.Headers = map[string]string{
//		"Accept-Encoding": "gzip,deflate,sdch",
//		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
//	}
//	resp, err := req.Get("http://httpbin.org/get")
//
// Upload Files:
//
//	f, err := os.Open("test.txt")
//	req.Files = []request.FileField{
//		request.FileField{"file", "test.txt", f},
//	}
//	resp, err := req.Post("http://httpbin.org/post")
//
// Json Body:
//
//	req.Json = map[string]string{
//		"a": "A",
//		"b": "B",
//	}
//	resp, err := req.Post("http://httpbin.org/post")
//	req.Json = []int{1, 2, 3}
//	resp, err = req.Post("http://httpbin.org/post")
//
// Proxy:
//
//	req.Proxy = "http://127.0.0.1:8080"
//	// req.Proxy = "https://127.0.0.1:8080"
//	// req.Proxy = "socks5://127.0.0.1:57341"
//	resp, err := req.Get("http://httpbin.org/get")
//
// HTTP Basic Authentication:
//
//	req.BasicAuth = request.BasicAuth{"user", "passwd"}
//	resp, err := req.Get("http://httpbin.org/basic-auth/user/passwd")
package request
