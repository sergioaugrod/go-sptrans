package sptrans

import (
	"fmt"
)

const (
	defaultStopPath         = "Parada/Buscar"
	defaultStopRoutePath    = "Parada/BuscarParadasPorLinha"
	defaultStopCorridorPath = "Parada/BuscarParadasPorCorredor"
)

// StopService provide functions to request stops endpoints
type StopService service

// Stop structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-paradas
type Stop struct {
	Cp int64   `json:"cp"`
	Np string  `json:"np"`
	Ed string  `json:"ed"`
	Py float64 `json:"py"`
	Px float64 `json:"px"`
}

// Search performs a search of the bus stops based on the parameter informed (stop name or address)
func (r *StopService) Search(filter string) ([]*Stop, error) {
	path := fmt.Sprintf("%s?termosBusca=%s", defaultStopPath, filter)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	if err != nil {
		return nil, err
	}

	return stops, err
}

// SearchByRoute performs a search of the bus stops based on the route code
func (r *StopService) SearchByRoute(routeCode int) ([]*Stop, error) {
	path := fmt.Sprintf("%s?codigoLinha=%d", defaultStopRoutePath, routeCode)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	if err != nil {
		return nil, err
	}

	return stops, err
}

// SearchByCorridor performs a search of the bus stops based on the corridor code
func (r *StopService) SearchByCorridor(corridorCode int) ([]*Stop, error) {
	path := fmt.Sprintf("%s?codigoCorredor=%d", defaultStopCorridorPath, corridorCode)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	if err != nil {
		return nil, err
	}

	return stops, err
}
