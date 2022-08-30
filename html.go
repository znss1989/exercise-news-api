package main

import (
	"io/ioutil"
	"net/http"
)

func getRawHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func sanitizeHtml(htmlRaw string) string {
	return p.Sanitize(htmlRaw)
}
