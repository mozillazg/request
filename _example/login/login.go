package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/mozillazg/request"
)

const (
	loginRequiredPageURL = "http://login-test.3sd.me:10081/"
	loginPageURL         = "http://login-test.3sd.me:10081/login/"
)

func home(req *request.Request) (statusCode int) {
	resp, err := req.Get(loginRequiredPageURL)
	if err != nil {
		return 500
	}
	return resp.StatusCode
}

func getCSRFToken(req *request.Request) (string, error) {
	resp, err := req.Get(loginPageURL)
	if err != nil {
		return "", err
	}
	s, err := resp.Text()
	if err != nil {
		return "", err
	}

	reInput := regexp.MustCompile(
		`<input\s+[^>]*?name=['"]csrfmiddlewaretoken['"'][^>]*>`,
	)
	input := reInput.FindString(s)
	reValue := regexp.MustCompile(`value=['"]([^'"]+)['"]`)
	csrfToken := reValue.FindStringSubmatch(input)
	if len(csrfToken) < 2 {
		return "", err
	}
	return csrfToken[1], err
}

func login(req *request.Request) error {
	_, err := req.Post(loginPageURL)
	return err
}

func main() {
	c := new(http.Client)
	req := request.NewRequest(c)
	log.Println(home(req)) // 403

	// login
	csrfToken, err := getCSRFToken(req)
	if err != nil {
		log.Fatal(err)
	}
	req.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"name":                "go-request",
		"password":            "go-request-passwd",
	}
	log.Println(csrfToken)
	err = login(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(home(req)) // 200
}
