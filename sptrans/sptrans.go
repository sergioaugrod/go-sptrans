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
	"strconv"
)

const (
	defaultBaseURL  = "http://api.olhovivo.sptrans.com.br/v2.1/"
	defaultAuthPath = "Login/Autenticar"
)

// Client is a SPTrans client for making Olho Vivo API requests
type Client struct {
	BaseURL         *url.URL
	HTTP            *http.Client
	Line            *LineService
	Stop            *StopService
	Corridor        *CorridorService
	Company         *CompanyService
	VehiclePosition *VehiclePositionService
	Forecast        *ForecastService
	Token           string
}

type service struct {
	client *Client
}

// NewClient returns client with default configurations
func NewClient(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	httpClient := http.DefaultClient

	client := &Client{BaseURL: baseURL, HTTP: httpClient, Token: token}
	client.Line = &LineService{client: client}
	client.Stop = &StopService{client: client}
	client.Corridor = &CorridorService{client: client}
	client.Company = &CompanyService{client: client}
	client.VehiclePosition = &VehiclePositionService{client: client}
	client.Forecast = &ForecastService{client: client}

	return client
}

// Request enables requests for a given api path, and decodes the return according to a structure
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
	resp, err := c.HTTP.Do(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &decoder)

	return resp, nil
}

// Authenticate authenticates according to client token
func (c *Client) Authenticate() (bool, error) {
	authPath := fmt.Sprintf("%s?token=%s", defaultAuthPath, c.Token)
	url, _ := c.BaseURL.Parse(authPath)
	resp, err := c.HTTP.Post(url.String(), "application/json", nil)
	body, err := ioutil.ReadAll(resp.Body)
	auth, err := strconv.ParseBool(string(body))

	if auth != true || err != nil {
		return false, err
	}

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(c.BaseURL, resp.Cookies())
	c.HTTP.Jar = jar

	return true, nil
}
