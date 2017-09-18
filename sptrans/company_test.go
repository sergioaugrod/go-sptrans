package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAllToReturnCompanies(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Empresa" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"00:00","e":[{"a":1,"e":[{"a": 1,"c": 999,"n": "NOME"}]}]}`)
	})

	companies, _ := client.Company.All()
	company := companies[0]

	if len(companies) != 1 {
		t.Error("Companies length different than 1")
	}

	if company.A != 1 {
		t.Error("Company a different than 1")
	}
}
