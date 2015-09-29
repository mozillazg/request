package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mozillazg/request"
)

func diff(req *request.Request) {
	url := "http://example.com:12345"
	start := time.Now()

	req.Get(url)

	diff := time.Now().Sub(start)
	fmt.Println(diff.Seconds())
}

func main() {
	c := new(http.Client)
	req := request.NewRequest(c)

	fmt.Println("default timeout")
	diff(req)

	timeout := time.Duration(1 * time.Second)
	c.Timeout = timeout
	fmt.Printf("set timeout = %f seconds\n", timeout.Seconds())
	diff(req)

	// Or use req.Client
	c = new(http.Client)
	req = request.NewRequest(c)
	req.Client.Timeout = timeout
	fmt.Printf("set timeout = %f seconds\n", timeout.Seconds())
	diff(req)
}
