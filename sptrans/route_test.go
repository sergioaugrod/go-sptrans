package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchToReturnRoutes(t *testing.T) {
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

	routes, _ := client.Route.Search("Lapa")
	route := routes[0]

	if len(routes) != 1 {
		t.Error("Routes length different than 1")
	}

	if route.Cl != 1234 {
		t.Error("Route cl different than 1234")
	}
}

func TestSearchByDirectionToReturnRoutes(t *testing.T) {
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

	routes, _ := client.Route.SearchByDirection("Lapa", 1)
	route := routes[0]

	if len(routes) != 1 {
		t.Error("Routes length different than 1")
	}

	if route.Cl != 1234 {
		t.Error("Route cl different than 1234")
	}
}
