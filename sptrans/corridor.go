package sptrans

const (
	defaultCorridorPath = "Corredor"
)

// CorridorService provide functions to request corridors endpoints
type CorridorService service

// Corridor structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-corredores
type Corridor struct {
	Id   int64  `json:"cc"`
	Name string `json:"nc"`
}

// All returns all smart corridors
func (r *CorridorService) All() ([]*Corridor, error) {
	var corridors []*Corridor
	_, err := r.client.Request("GET", defaultCorridorPath, nil, &corridors)

	if err != nil {
		return nil, err
	}

	return corridors, err
}
