package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/mozillazg/request"
)

func home(a *request.Args) (statusCode int) {
	resp, err := request.Get("http://login-test.3sd.me:10081/", a)
	if err != nil {
		return 500
	}
	return resp.StatusCode
}

func getCSRFToken(a *request.Args) (string, error) {
	url := "http://login-test.3sd.me:10081/login/"
	resp, err := request.Get(url, a)
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

func login(a *request.Args) error {
	url := "http://login-test.3sd.me:10081/login/"
	_, err := request.Post(url, a)
	return err
}

func main() {
	c := new(http.Client)
	a := request.NewArgs(c)
	log.Println(home(a)) // 403

	// login
	csrfToken, err := getCSRFToken(a)
	if err != nil {
		log.Fatal(err)
	}
	a.Data = map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"name":                "go-request",
		"password":            "go-request-passwd",
	}
	log.Println(csrfToken)
	err = login(a)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(home(a)) // 200
}
