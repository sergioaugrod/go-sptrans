package sptrans

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient("123456")
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url
}

func tearDown() {
	server.Close()
}

func TestNewClientToReturnClient(t *testing.T) {
	setup()
	defer tearDown()

	if client.Token != "123456" {
		t.Error("Client token is not equal 123456")
	}
}

func TestAuthenticateToReturnTrue(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "POST" != r.Method {
			t.Error("Request method is not POST")
		}

		if r.URL.String() != "/Login/Autenticar?token=123456" {
			t.Error("Incorrect requested url")
		}

		fmt.Fprint(w, `true`)
	})

	auth, _ := client.Authenticate()

	if auth != true {
		t.Error("Authenticate not return true")
	}
}

func TestAuthenticateToReturnFalse(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "POST" != r.Method {
			t.Error("Request method is not POST")
		}

		if r.URL.String() != "/Login/Autenticar?token=123456" {
			t.Error("Incorrect requested url")
		}

		fmt.Fprint(w, `false`)
	})

	auth, err := client.Authenticate()

	if auth != false && err != nil {
		t.Error("Authenticate not return false")
	}
}

func TestAuthenticateToReturnError(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "POST" != r.Method {
			t.Error("Request method is not POST")
		}

		if r.URL.String() != "/Login/Autenticar?token=123456" {
			t.Error("Incorrect requested url")
		}

		fmt.Fprint(w, nil)
	})

	_, err := client.Authenticate()

	if err == nil {
		t.Error("Authenticate not return error")
	}
}
