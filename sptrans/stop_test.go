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
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.Search("Afonso")
	stop := stops[0]

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}

	if stop.Id != 1234 {
		t.Error("Stop Id different than 1234")
	}

	if stop.Name != "AFONSO BRAZ" {
		t.Error("Stop Name different than AFONSO BRAZ")
	}

	if stop.Address != "R ARMINDA" {
		t.Error("Stop Address different than R ARMINDA")
	}

	if stop.Latitude != -23.592938 {
		t.Error("Stop Latitude different than -23.592938")
	}

	if stop.Longitude != -46.672727 {
		t.Error("Stop Longitude different than -46.672727")
	}
}

func TestSearchByLineToReturnStops(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Parada/BuscarParadasPorLinha?codigoLinha=123456" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.SearchByLine(123456)

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}
}

func TestSearchByCorridorToReturnStops(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Parada/BuscarParadasPorCorredor?codigoCorredor=123456" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cp": 1234, "np": "AFONSO BRAZ", "ed": "R ARMINDA", "py": -23.592938, "px": -46.672727}]`)
	})

	stops, _ := client.Stop.SearchByCorridor(123456)

	if len(stops) != 1 {
		t.Error("Stops length different than 1")
	}
}
