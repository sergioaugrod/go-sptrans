package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAllToReturnCorridors(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Corredor" {
			t.Error("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `[{"cc":1234,"lc":"Corredor"}]`)
	})

	corridors, _ := client.Corridor.All()
	corridor := corridors[0]

	if len(corridors) != 1 {
		t.Error("Corridors length different than 1")
	}

	if corridor.Cc != 1234 {
		t.Error("Corridor cc different than 1234")
	}
}
