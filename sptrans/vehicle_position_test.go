package sptrans

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAllToReturnVehiclesPosition(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Posicao" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"11:30","l":[{"c":"5015-10","cl":33887,"sl":2,"lt0":"METRÔ JABAQUARA","lt1":"JD. SÃO JORGE","qv":1,"vs":[{"p":68021,"a":true,"ta":"2017-05-12T14:30:37Z","py":-23.678712500000003,"px":-46.65674}]}]}`)
	})

	vehiclesPosition, _ := client.VehiclePosition.All()
	vehiclePosition := vehiclesPosition[0]
	vehicles := vehiclePosition.Vehicles
	vehicle := vehicles[0]

	if vehiclePosition.FullSign != "5015-10" {
		t.Error("VehiclePosition FullSign different than 5015-10")
	}

	if vehiclePosition.Id != 33887 {
		t.Error("VehiclePosition Id different than 33887")
	}

	if vehiclePosition.Direction != 2 {
		t.Error("VehiclePosition Direction different than 2")
	}

	if vehiclePosition.DestinationSign != "METRÔ JABAQUARA" {
		t.Error("VehiclePosition DestinationSign different than METRÔ JABAQUARA")
	}

	if vehiclePosition.OriginSign != "JD. SÃO JORGE" {
		t.Error("VehiclePosition OriginSign different than JD. SÃO JORGE")
	}

	if vehiclePosition.QuantityVehicles != 1 {
		t.Error("VehiclePosition Quantity different than 1")
	}

	if vehicle.Prefix != 68021 {
		t.Error("vehicle Prefix different than 68021")
	}

	if vehicle.Accessible != true {
		t.Error("vehicle Accessible different than true")
	}

	if vehicle.Hour != "2017-05-12T14:30:37Z" {
		t.Error("vehicle Hour different than 2017-05-12T14:30:37Z")
	}

	if vehicle.Latitude != -23.678712500000003 {
		t.Error("vehicle Latitude different than -23.678712500000003")
	}

	if vehicle.Longitude != -46.65674 {
		t.Error("vehicle Longitude different than -46.65674")
	}
}

func TestSearchByLineToReturnVehiclesPosition(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Posicao/Linha?codigoLinha=123456" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"19:57","vs":[{"p":"11433","a":false,"ta":"2017-05-07T22:57:02Z","py":-23.540150375000003,"px":-46.64414075}]}`)
	})

	vehiclesPosition, _ := client.VehiclePosition.SearchByLine(123456)

	if len(vehiclesPosition) != 1 {
		t.Error("VehiclesPosition length different than 1")
	}
}

func TestSearchByCompanyToReturnVehiclesPosition(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if "GET" != r.Method {
			t.Error("Request method is not GET")
		}

		if r.URL.String() != "/Posicao/Garagem?codigoEmpresa=123456" {
			t.Errorf("Incorrect requested url: %s", r.URL.String())
		}

		fmt.Fprint(w, `{"hr":"11:30","l":[{"c":"5015-10","cl":33887,"sl":2,"lt0":"METRÔ JABAQUARA","lt1":"JD. SÃO JORGE","qv":1,"vs":[{"p":68021,"a":true,"ta":"2017-05-12T14:30:37Z","py":-23.678712500000003,"px":-46.65674}]}]}`)
	})

	vehiclesPosition, _ := client.VehiclePosition.SearchByCompany(123456)

	if len(vehiclesPosition) != 1 {
		t.Error("VehiclesPosition length different than 1")
	}
}
