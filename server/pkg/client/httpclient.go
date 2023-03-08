package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type HttpClient struct {
	Host string
	Port uint64
}

func (c HttpClient) NewClient() *httpRequestClient {
	baseURI := c.Host
	if c.Port > 0 {
		baseURI = fmt.Sprintf("%s:%d", baseURI, c.Port)
	}

	return &httpRequestClient{BaseURI: baseURI}
}

type httpRequestClient struct {
	BaseURI string
}

func (c *httpRequestClient) params2string(params map[string]string) string {
	var query string

	for key, value := range params {
		if query == "" {
			query = fmt.Sprintf("%s=%s", key, value)
		} else {
			query = fmt.Sprintf("%s&%s=%s", query, key, value)
		}
	}

	return query
}

type RequestOption struct {
	Method   string
	Path     string
	Params   *map[string]string
	Headers  *map[string]string
	Data     interface{}
	Duration uint
}

func (c *httpRequestClient) Request(option RequestOption) ([]byte, error) {
	var url string = option.Path
	var query string
	var body *strings.Reader

	if option.Params != nil {
		query = c.params2string(*option.Params)
	}

	if query != "" {
		if ok, _ := regexp.MatchString("\\?", url); ok {
			url = fmt.Sprintf("%s&%s", url, query)
		} else {
			url = fmt.Sprintf("%s?%s", url, query)
		}
	}

	var req *http.Request
	var err error

	if option.Data != nil {
		data, _ := json.Marshal(option.Data)
		body = strings.NewReader(string(data))
		req, err = http.NewRequest(option.Method, url, body)
	} else {
		req, err = http.NewRequest(option.Method, url, nil)
	}

	if err != nil {
		bytes, _ := json.Marshal(struct {
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		}{ErrCode: 500, ErrMsg: "service connenct err"})

		return nil, errors.New(string(bytes))
	}

	// 设置headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if option.Headers != nil {
		for k, v := range *option.Headers {
			req.Header.Set(k, v)
		}
	}

	duration := option.Duration
	if duration == 0 {
		duration = 5
	}
	client := &http.Client{Timeout: time.Second * time.Duration(duration)}
	resp, err := client.Do(req)
	if err != nil {
		bytes, _ := json.Marshal(struct {
			ErrCode int64  `json:"err_code"`
			ErrMsg  string `json:"err_msg"`
		}{ErrCode: 500, ErrMsg: err.Error()})

		return nil, errors.New(string(bytes))
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
