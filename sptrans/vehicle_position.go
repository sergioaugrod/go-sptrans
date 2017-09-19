package sptrans

const (
	defaultVehiclePositionPath = "Posicao"
)

// VehiclePositionService provide functions to request vehicles position endpoints
type VehiclePositionService service

// VehiclePositionResponse structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-posicao
type VehiclePositionResponse struct {
	Hour  string                 `json:"hr"`
	Lines []*VehiclePositionLine `json:"l"`
}

// VehiclePositionLine relation of localized lines
type VehiclePositionLine struct {
	FullSign         string             `json:"c"`
	Id               int64              `json:"cl"`
	Direction        int64              `json:"sl"`
	DestinationSign  string             `json:"lt0"`
	OriginSign       string             `json:"lt1"`
	QuantityVehicles int64              `json:"qv"`
	Vehicles         []*VehiclePosition `json:"vs"`
}

// VehiclePosition relation of localized vehicles
type VehiclePosition struct {
	Prefix     int64   `json:"p"`
	Accessible bool    `json:"a"`
	Hour       string  `json:"ta"`
	Latitude   float64 `json:"py"`
	Longitude  float64 `json:"px"`
}

// All returns all vehicles position
func (r *VehiclePositionService) All() ([]*VehiclePositionLine, error) {
	var vehiclePositionResponse *VehiclePositionResponse
	_, err := r.client.Request("GET", defaultVehiclePositionPath, nil, &vehiclePositionResponse)

	return vehiclePositionResponse.Lines, err
}
