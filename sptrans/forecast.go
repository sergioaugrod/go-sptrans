package sptrans

import (
	"fmt"
)

const (
	defaultForecastPath     = "Previsao"
	defaultForecastLinePath = "Previsao/Linha"
	defaultForecastStopPath = "Previsao/Parada"
)

// ForecastService provide functions to request forecasts endpoints
type ForecastService service

// ForecastResponse structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-previsao
type ForecastResponse struct {
	Hour string       `json:"hr"`
	Stop ForecastStop `json:"p"`
}

// ForecastLineResponse is a line response structure with hour and a list of stops
type ForecastLineResponse struct {
	Hour  string              `json:"hr"`
	Stops []*ForecastLineStop `json:"ps"`
}

// ForecastLineStop is a line stop structure
type ForecastLineStop struct {
	Id        int64              `json:"cp"`
	Name      string             `json:"np"`
	Latitude  float64            `json:"py"`
	Longitude float64            `json:"px"`
	Vehicles  []*ForecastVehicle `json:"vs"`
}

// ForecastStop is a stop structure
type ForecastStop struct {
	Id        int64           `json:"cp"`
	Name      string          `json:"np"`
	Latitude  float64         `json:"py"`
	Longitude float64         `json:"px"`
	Lines     []*ForecastLine `json:"l"`
}

//ForecastLine is a line structure
type ForecastLine struct {
	FullSign         string             `json:"c"`
	Id               int64              `json:"cl"`
	Direction        int64              `json:"sl"`
	DestinationSign  string             `json:"lt0"`
	OriginSign       string             `json:"lt1"`
	QuantityVehicles int64              `json:"qv"`
	Vehicles         []*ForecastVehicle `json:"vs"`
}

//ForecastVehicle is a vehicle structure
type ForecastVehicle struct {
	Prefix     int64   `json:"p"`
	Hour       string  `json:"t"`
	Accessible bool    `json:"a"`
	UpdateTime string  `json:"ta"`
	Latitude   float64 `json:"py"`
	Longitude  float64 `json:"px"`
}

// Search performs a search of bus arrivals based on stop id or line id
func (r *ForecastService) Search(stopId int, lineId int) (ForecastStop, error) {
	path := fmt.Sprintf("%s?codigoParada=%d&codigoLinha=%d", defaultForecastPath, stopId, lineId)

	var forecastResponse *ForecastResponse
	_, err := r.client.Request("GET", path, nil, &forecastResponse)

	return forecastResponse.Stop, err
}

// SearchByLine performs a search of the bus arrivals based on the line id
func (r *ForecastService) SearchByLine(lineId int) ([]*ForecastLineStop, error) {
	path := fmt.Sprintf("%s?codigoLinha=%d", defaultForecastLinePath, lineId)

	var forecastLineResponse *ForecastLineResponse
	_, err := r.client.Request("GET", path, nil, &forecastLineResponse)

	return forecastLineResponse.Stops, err
}

// SearchByStop performs a search of the bus arrivals based on the stop id
func (r *ForecastService) SearchByStop(stopId int) (ForecastStop, error) {
	path := fmt.Sprintf("%s?codigoParada=%d", defaultForecastStopPath, stopId)

	var forecastResponse *ForecastResponse
	_, err := r.client.Request("GET", path, nil, &forecastResponse)

	return forecastResponse.Stop, err
}
