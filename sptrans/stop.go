package sptrans

import (
	"fmt"
)

const (
	defaultStopPath         = "Parada/Buscar"
	defaultStopLinePath     = "Parada/BuscarParadasPorLinha"
	defaultStopCorridorPath = "Parada/BuscarParadasPorCorredor"
)

// StopService provide functions to request stops endpoints
type StopService service

// Stop structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-paradas
type Stop struct {
	Id        int64   `json:"cp"`
	Name      string  `json:"np"`
	Address   string  `json:"ed"`
	Latitude  float64 `json:"py"`
	Longitude float64 `json:"px"`
}

// Search performs a search of the bus stops based on the parameter informed (stop name or address)
func (r *StopService) Search(filter string) ([]*Stop, error) {
	path := fmt.Sprintf("%s?termosBusca=%s", defaultStopPath, filter)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	return stops, err
}

// SearchByLine performs a search of the bus stops based on the line id
func (r *StopService) SearchByLine(lineId int) ([]*Stop, error) {
	path := fmt.Sprintf("%s?codigoLinha=%d", defaultStopLinePath, lineId)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	return stops, err
}

// SearchByCorridor performs a search of the bus stops based on the corridor id
func (r *StopService) SearchByCorridor(corridorId int) ([]*Stop, error) {
	path := fmt.Sprintf("%s?codigoCorredor=%d", defaultStopCorridorPath, corridorId)
	var stops []*Stop
	_, err := r.client.Request("GET", path, nil, &stops)

	return stops, err
}
