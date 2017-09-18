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
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cl":1234,"lc":true,"lt":"N1234","sl":2,"tl":11,"tp":"TERM. PINHEIROS","ts":"TERM. LAPA"}]`)
	})

	lines, _ := client.Line.Search("Lapa")
	line := lines[0]

	if len(lines) != 1 {
		t.Error("Lines length different than 1")
	}

	if line.Id != 1234 {
		t.Error("Line Id different than 1234")
	}

	if line.Circular != true {
		t.Error("Line Circular different than true")
	}

	if line.DisplaySign != "N1234" {
		t.Error("Line DisplaySign different than N1234")
	}

	if line.Type != 11 {
		t.Error("Line Type different than 11")
	}

	if line.Direction != 2 {
		t.Error("Line Direction different than 2")
	}

	if line.MainTerminal != "TERM. PINHEIROS" {
		t.Error("Line MainTerminal different than TERM. PINHEIROS")
	}

	if line.SecondaryTerminal != "TERM. LAPA" {
		t.Error("Line SecondaryTerminal different than TERM. LAPA")
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
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cl":1234,"lc":true,"lt":"N1234","sl":2,"tl":11,"tp":"TERM. PINHEIROS","ts":"TERM. LAPA"}]`)
	})

	lines, _ := client.Line.SearchByDirection("Lapa", 1)

	if len(lines) != 1 {
		t.Error("Lines length different than 1")
	}
}
