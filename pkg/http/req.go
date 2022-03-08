package http

import (
	"io/ioutil"
	"log"
	"net/http"
	httpUrl "net/url"
	"strings"
)

type AbsHttp interface {
	Get(url string) *http.Response
	PostLogin(url string, login string, password string) *http.Response
	GetContent(resp *http.Response) (string, error)
}

type ExtendedHttp struct {
	Client *http.Client
}

func NewHttp() *ExtendedHttp {
	client := &http.Client{}

	return &ExtendedHttp{
		Client: client,
	}
}

var headers http.Header = http.Header{
	"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.174 YaBrowser/22.1.3.848 Yowser/2.5 Safari/537.36"},
}

func (h *ExtendedHttp) Get(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	req.Header = headers
	response, err := h.Client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return response
}

func (h *ExtendedHttp) LoginPost(url string, login string, password string) *http.Response {
	data := httpUrl.Values{
		"login":    {login},
		"password": {password},
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header = headers

	if err != nil {
		log.Fatal(err)
	}
	response, err := h.Client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func (h *ExtendedHttp) GetContent(resp *http.Response) (string, error) {
	htmlBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(htmlBody), nil
}
