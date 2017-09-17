package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchToReturnStops(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Parada/Buscar?termosBusca=Afonso" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.Search("Afonso")
	stop := stops[0]

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}

	if stop.Cp != 1234 {
		t.Error("Stop cp different than 1234")
	}
}

func TestSearchByRoute(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Parada/BuscarParadasPorLinha?codigoLinha=123456" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.SearchByRoute(123456)
	stop := stops[0]

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}

	if stop.Cp != 1234 {
		t.Error("Stop cp different than 1234")
	}
}

func TestSearchByCorridor(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Parada/BuscarParadasPorCorredor?codigoCorredor=123456" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.SearchByCorridor(123456)
	stop := stops[0]

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}

	if stop.Cp != 1234 {
		t.Error("Stop cp different than 1234")
	}
}
