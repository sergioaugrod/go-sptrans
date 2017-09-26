package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSearchToReturnForecastStop(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Previsao?codigoParada=4200953&codigoLinha=2550" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"20:09","p":{"cp":4200953,"np":"PARADA ROBERTO SELMI DEI B/C","py":-23.675901,"px":-46.752812,"l":[{"c":"7021-10","cl":1989,"sl":1,"lt0":"TERM. JOÃO DIAS","lt1":"JD. MARACÁ","qv":1,"vs":[{"p":74558,"t":"23:11","a":true,"ta":"2017-05-07T23:09:05Z","py":-23.67603,"px":-46.75891166666667}]}]}}`)
	})

	forecastStop, _ := client.Forecast.Search(4200953, 2550)
	forecastLine := forecastStop.Lines[0]
	forecastVehicle := forecastLine.Vehicles[0]

	if forecastVehicle.Hour != "23:11" {
		t.Error("ForecastVehicle Hour different than 23:11")
	}

	if forecastVehicle.Prefix != 74558 {
		t.Error("ForecastVehicle Prefix different than 74558")
	}
}

func TestSearchByLineToReturnForecastLineStops(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Previsao/Linha?codigoLinha=1989" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"20:18","ps":[{"cp":700016623,"np":"ANA CINTRA B/C","py":-23.538763,"px":-46.646925,"vs":[{"p":"11436","t":"23:26","a":false,"ta":"2017-05-07T23:18:02Z","py":-23.528119999999998,"px":-46.670674999999996}]}]}`)
	})

	forecastLineStops, _ := client.Forecast.SearchByLine(1989)
	forecastLineStop := forecastLineStops[0]

	if forecastLineStop.Name != "ANA CINTRA B/C" {
		t.Error("ForecastLineStop Name different than ANA CINTRA B/C")
	}
}

func TestSearchByStopToReturnForecastStop(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Previsao/Parada?codigoParada=4200953" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"20:09","p":{"cp":4200953,"np":"PARADA ROBERTO SELMI DEI B/C","py":-23.675901,"px":-46.752812,"l":[{"c":"7021-10","cl":1989,"sl":1,"lt0":"TERM. JOÃO DIAS","lt1":"JD. MARACÁ","qv":1,"vs":[{"p":74558,"t":"23:11","a":true,"ta":"2017-05-07T23:09:05Z","py":-23.67603,"px":-46.75891166666667}]}]}}`)
	})

	forecastStop, _ := client.Forecast.SearchByStop(4200953)
	forecastLine := forecastStop.Lines[0]
	forecastVehicle := forecastLine.Vehicles[0]

	if forecastVehicle.Hour != "23:11" {
		t.Error("ForecastVehicle Hour different than 23:11")
	}

	if forecastVehicle.Prefix != 74558 {
		t.Error("ForecastVehicle Prefix different than 74558")
	}
}
