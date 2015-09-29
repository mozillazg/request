package request

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestHead(t *testing.T) {
	c := new(http.Client)
	req := NewRequest(c)
	url := "http://httpbin.org/get"
	resp, _ := req.Head(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}
