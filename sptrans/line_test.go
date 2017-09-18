package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchToReturnLines(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Linha/Buscar?termosBusca=Lapa" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cl":1234,"lc":true,"lt":"N1234","sl":2,"tl":11,"tp":"TERM. PINHEIROS","ts":"TERM. LAPA"}]`)
	})

	lines, _ := client.Line.Search("Lapa")
	line := lines[0]

	if len(lines) != 1 {
		t.Error("Lines length different than 1")
	}

	if line.Cl != 1234 {
		t.Error("Line cl different than 1234")
	}
}

func TestSearchByDirectionToReturnLines(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Linha/BuscarLinhaSentido?termosBusca=Lapa&sentido=1" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cl":1234,"lc":true,"lt":"N1234","sl":2,"tl":11,"tp":"TERM. PINHEIROS","ts":"TERM. LAPA"}]`)
	})

	lines, _ := client.Line.SearchByDirection("Lapa", 1)
	line := lines[0]

	if len(lines) != 1 {
		t.Error("Lines length different than 1")
	}

	if line.Cl != 1234 {
		t.Error("Line cl different than 1234")
	}
}
