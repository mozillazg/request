package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"io/ioutil"

	"github.com/bmizerany/assert"
)

type hookNothing struct {
	callBeforeHook bool
	callAfterHook  bool
}

func (h *hookNothing) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	h.callBeforeHook = true
	return
}
func (h *hookNothing) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	h.callAfterHook = true
	return
}

func TestHookNothing(t *testing.T) {
	h := &hookNothing{}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	resp, _ := req.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, h.callBeforeHook, true)
	assert.Equal(t, h.callAfterHook, true)
}

func TestHookNothingError(t *testing.T) {
	h := &hookNothing{}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	_, err := req.Get("http://127.0.0.1:12345/get")
	assert.Equal(t, err != nil, true)
}

type beforeRequestHookError struct {
	err error
}

func (h *beforeRequestHookError) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	err = h.err
	return
}
func (h *beforeRequestHookError) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	return
}

func TestBeforeRequestHookError(t *testing.T) {
	e := errors.New("beforeRequestHookError")
	h := &beforeRequestHookError{e}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	_, err := req.Get("https://httpbin.org/get")
	assert.Equal(t, err, h.err)
}

type beforeRequestHookResp struct {
	resp *http.Response
}

func (h *beforeRequestHookResp) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	resp = h.resp
	return
}
func (h *beforeRequestHookResp) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	return
}

func TestBeforeRequestHookResp(t *testing.T) {
	j, _ := json.Marshal(map[string]string{
		"url": "http://test",
	})
	b := ioutil.NopCloser(bytes.NewReader(j))
	r := &http.Response{Body: b}
	h := &beforeRequestHookResp{r}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	resp, _ := req.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	assert.Equal(t, resp.Response, h.resp)
}

type afterRequestHookError struct {
	err error
}

func (h *afterRequestHookError) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	return
}
func (h *afterRequestHookError) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	err = h.err
	return
}

func TestAfterRequestHookError(t *testing.T) {
	e := errors.New("afterRequestHookError")
	h := &beforeRequestHookError{e}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	_, err := req.Get("https://httpbin.org/get")
	assert.Equal(t, err, h.err)
}

type afterRequestHookResp struct {
	resp *http.Response
}

func (h *afterRequestHookResp) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	return
}
func (h *afterRequestHookResp) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	newResp = h.resp
	return
}

func TestAfterRequestHookResp(t *testing.T) {
	j, _ := json.Marshal(map[string]string{
		"url": "http://test",
	})
	b := ioutil.NopCloser(bytes.NewReader(j))
	r := &http.Response{Body: b}
	h := &afterRequestHookResp{r}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	resp, _ := req.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	assert.Equal(t, resp.Response, h.resp)
}

type allHookResp struct {
	beforeResp *http.Response
	afterResp  *http.Response
}

func (h *allHookResp) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	resp = h.beforeResp
	return
}
func (h *allHookResp) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	newResp = h.afterResp
	return
}

func TestAllHookResp(t *testing.T) {
	j, _ := json.Marshal(map[string]string{
		"url": "http://test",
	})
	b1 := ioutil.NopCloser(bytes.NewReader(j))
	b2 := ioutil.NopCloser(bytes.NewBuffer([]byte("test")))
	beforeR := &http.Response{Body: b1}
	afterR := &http.Response{Body: b2}
	h := &allHookResp{beforeR, afterR}

	c := &http.Client{}
	req := NewRequest(c)
	req.Hooks = []Hook{h}
	resp, _ := req.Get("https://httpbin.org/get")
	defer resp.Body.Close()
	assert.Equal(t, resp.Response, h.beforeResp)
}
