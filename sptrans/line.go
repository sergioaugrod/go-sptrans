package sptrans

import (
	"fmt"
)

const (
	defaultLinePath          = "Linha/Buscar"
	defaultLineDirectionPath = "Linha/BuscarLinhaSentido"
)

// LineService provide functions to request lines endpoints
type LineService service

// Line structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-linhas
type Line struct {
	Cl int64  `json:"cl"`
	Lc bool   `json:"lc"`
	Lt string `json:"lt"`
	Sl int64  `json:"sl"`
	Tl int64  `json:"tl"`
	Tp string `json:"tp"`
	Ts string `json:"ts"`
}

// Search performs a search of the bus lines based on the parameter informed (description or line number)
func (r *LineService) Search(filter string) ([]*Line, error) {
	path := fmt.Sprintf("%s?termosBusca=%s", defaultLinePath, filter)
	var lines []*Line
	_, err := r.client.Request("GET", path, nil, &lines)

	if err != nil {
		return nil, err
	}

	return lines, err
}

// SearchByDirection performs a search of the bus lines based on the parameter informed (description or direction)
func (r *LineService) SearchByDirection(filter string, direction int) ([]*Line, error) {
	path := fmt.Sprintf("%s?termosBusca=%s&sentido=%d", defaultLineDirectionPath, filter, direction)
	var lines []*Line
	_, err := r.client.Request("GET", path, nil, &lines)

	if err != nil {
		return nil, err
	}

	return lines, err
}
