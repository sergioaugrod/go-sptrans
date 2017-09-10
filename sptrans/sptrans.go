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

func NewClient(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	httpClient := http.DefaultClient

	client := &Client{BaseURL: baseURL, Http: httpClient, Token: token}
	client.Route = &RouteService{client: client}

	return client
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

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &decoder)

	return resp, nil
}

func (c *Client) Authenticate() (bool, error) {
	authPath := fmt.Sprintf("%s?token=%s", defaultAuthPath, c.Token)
	url, _ := c.BaseURL.Parse(authPath)
	resp, err := c.Http.Post(url.String(), "application/json", nil)

	if err != nil {
		return false, err
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(c.BaseURL, resp.Cookies())
	c.Http.Jar = jar

	return true, nil
}
