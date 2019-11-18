package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

func HttpGet(url, username, password string) (*http.Response, error) {
	return Http(HTTPGet, url, username, password, nil)
}

func HttpPost(url, username, password string, data []byte) (*http.Response, error) {
	return Http(HTTPPost, url, username, password, data)
}

func HttpDelete(url, username, password string, data []byte) (*http.Response, error) {
	return Http(HTTPDelete, url, username, password, data)
}

func Http(method, url, username, password string, data []byte) (*http.Response, error) {
	if !isValidHTTPMethod(method) {
		return nil, errors.New("HTTP method is illegal")
	}

	var b io.Reader
	if data != nil {
		b = bytes.NewBuffer(data)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func isValidHTTPMethod(method string) bool {
	return strings.EqualFold(method, HTTPGet) || strings.EqualFold(method, HTTPPost) || strings.EqualFold(method, HTTPPut) || strings.EqualFold(method, HTTPDelete)
}
