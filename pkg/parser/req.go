package parser

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	req.Header = http.Header{
		"user-agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.174 YaBrowser/22.1.3.848 Yowser/2.5 Safari/537.36"},
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return response
}

func Login(url string, login string, password string) *http.Response {
	response, err := http.PostForm(url, map[string][]string{})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return response
}

func GetContent(resp *http.Response) (string, error) {
	htmlBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(htmlBody), nil
}
