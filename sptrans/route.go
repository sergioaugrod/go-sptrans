package sptrans

import (
	"fmt"
)

const (
	defaultRoutePath          = "Linha/Buscar"
	defaultRouteDirectionPath = "Linha/BuscarLinhaSentido"
)

type RouteService service

// Route structure
// Endpoint Documentation: http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-linhas
type Route struct {
	Cl int64  `json:"cl"`
	Lc bool   `json:"lc"`
	Lt string `json:"lt"`
	Sl int64  `json:"sl"`
	Tl int64  `json:"tl"`
	Tp string `json:"tp"`
	Ts string `json:"ts"`
}

// Serach performs a search of the bus lines based on the parameter informed (description or line number)
// Example:
// client.Route.Search("Lapa")
// client.Route.Search("8000")
func (r *RouteService) Search(filter string) ([]*Route, error) {
	path := fmt.Sprintf("%s?termosBusca=%s", defaultRoutePath, filter)
	var routes []*Route
	_, err := r.client.Request("GET", path, nil, &routes)

	if err != nil {
		return nil, err
	}

	return routes, err
}

// SearchByDirection performs a search of the bus lines based on the parameter informed (description or direction)
// Direction must be:
// 1: Main Terminal for Secondary Terminal; 2: for Secondary Terminal for Main Terminal
// Example:
// client.Route.SearchByDirection("Lapa", 1)
// client.Route.SearchByDirection("Lapa", 2)
func (r *RouteService) SearchByDirection(filter string, direction int) ([]*Route, error) {
	path := fmt.Sprintf("%s?termosBusca=%s&sentido=%d", defaultRouteDirectionPath, filter, direction)
	var routes []*Route
	_, err := r.client.Request("GET", path, nil, &routes)

	if err != nil {
		return nil, err
	}

	return routes, err
}
