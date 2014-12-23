package main

import (
	"fmt"
	"github.com/mozillazg/request"
)

func main() {
	a := request.NewArgs()
	resq, err := request.Get("http://httpbin.org", a)
	if err == nil {
		fmt.Println(resq.OK())
		fmt.Println(resq.Reason())
	}
	resq, err = request.Head("http://httpbin.org", a)
	if err == nil {
		fmt.Println(resq.OK())
		fmt.Println(resq.Reason())
	}
}
