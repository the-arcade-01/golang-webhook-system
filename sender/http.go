package main

import (
	"net/http"
	"time"
)

type RestClient struct {
	client     *http.Client
	maxRetries int
	baseDelay  int
}

func NewRestClient(client *http.Client, maxRetries int, baseDelay int) *RestClient {
	return &RestClient{
		client:     client,
		maxRetries: maxRetries,
		baseDelay:  baseDelay,
	}
}

func (r *RestClient) Do(req *http.Request) (*http.Response, error) {
	var err error
	var res *http.Response
	for try := 0; try <= r.maxRetries; try++ {
		res, err = r.client.Do(req)
		if err == nil && res.StatusCode < 500 {
			return res, nil
		}

		if try < r.maxRetries {
			time.Sleep(time.Duration(r.baseDelay) * time.Duration(1<<try))
		}
	}
	return res, err
}
