package sptrans

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClientToReturnClient(t *testing.T) {
	token := "123456"
	client := NewClient(token)

	if client.Token != token {
		t.Errorf("Client token is not equal %s", token)
	}
}

func TestAuthenticateToSetHttpJar(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient("123456")
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "POST" != r.Method {
			t.Error("Request method is not POST")
		}

		if r.URL.String() != "/Login/Autenticar?token=123456" {
			t.Error("Incorrect requested url")
		}

		fmt.Fprint(w, `true`)
	})

	client.Authenticate()
}
