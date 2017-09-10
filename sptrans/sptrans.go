package sptrans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	defaultBaseURL  = "http://api.olhovivo.sptrans.com.br/v2.1/"
	defaultAuthPath = "Login/Autenticar"
)

type Client struct {
	BaseURL *url.URL
	Http    *http.Client
	Route   *RouteService
	Token   string
}

type service struct {
	client *Client
}

func (c *Client) Request(method, path string, bodyParams interface{}, decoder interface{}) (*http.Response, error) {
	var buffer io.ReadWriter
	url, _ := c.BaseURL.Parse(path)

	if bodyParams != nil {
		buffer = new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(bodyParams)
		if err != nil {
			return nil, err
		}
	}

	req, _ := http.NewRequest(method, url.String(), buffer)
	resp, err := c.Http.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &decoder)

	return resp, nil
}

func NewClient(token string) (*Client, error) {
	baseURL, _ := url.Parse(defaultBaseURL)
	cookieJar, err := getAuthCookieJar(token, baseURL)

	if err != nil {
		return nil, fmt.Errorf("Authentication failed with token %s.", token)
	}

	httpClient := &http.Client{Jar: cookieJar}

	client := &Client{BaseURL: baseURL, Http: httpClient, Token: token}
	client.Route = &RouteService{client: client}

	return client, nil
}

func getAuthCookieJar(token string, baseURL *url.URL) (*cookiejar.Jar, error) {
	authPath := fmt.Sprintf("%s?token=%s", defaultAuthPath, token)
	url, _ := baseURL.Parse(authPath)
	resp, err := http.Post(url.String(), "application/json", nil)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(baseURL, resp.Cookies())

	return jar, nil
}
