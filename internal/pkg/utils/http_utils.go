package utils

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request)(*http.Response, error)
}

func NewHTTPRequest(ctx context.Context, method string, URL string, body io.Reader)(*http.Request, error) {
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		return req, nil
	}
	return req.WithContext(ctx), nil
}

func ExecuteHttpRequest(client HTTPClient, URL string, queryMap map[string]string,
	method string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := NewHTTPRequest(context.Background(), method, URL, body)
	if err != nil {

	}

	q := req.URL.Query()
	for key, value := range queryMap {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {

	}
	defer CloseQuietly(res.Body)

	responseBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {

	}

	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	return res, nil
}

func CloseQuietly(r io.Closer){
	err := r.Close()
	if err != nil {
		panic(err)
	}
}