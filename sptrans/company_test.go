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
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"00:00","e":[{"a":1,"e":[{"a": 1,"c": 999,"n": "Company"}]}]}`)
	})

	companyOperations, _ := client.Company.All()
	companyOperation := companyOperations[0]
	companies := companyOperations[0].Companies
	company := companies[0]

	if len(companyOperations) != 1 {
		t.Error("CompanyOperations length different than 1")
	}

	if len(companies) != 1 {
		t.Error("Companies length different than 1")
	}

	if companyOperation.Id != 1 {
		t.Error("CompanyOperation Id different than 1")
	}

	if company.Id != 999 {
		t.Error("Company Id different than 999")
	}

	if company.OperationId != 1 {
		t.Error("Company OperationId different than 1")
	}

	if company.Name != "Company" {
		t.Error("Company Name different than Company")
	}
}
