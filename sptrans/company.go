package sptrans

const (
	defaultCompanyPath = "Empresa"
)

// CompanyService provide functions to request companies endpoints
type CompanyService service

// CompanyResponse structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-empresas
type CompanyResponse struct {
	Hour       string              `json:"hr"`
	Operations []*CompanyOperation `json:"e"`
}

// CompanyOperation structure
type CompanyOperation struct {
	Id        int64      `json:"a"`
	Companies []*Company `json:"e"`
}

//Company structure
type Company struct {
	Id          int64  `json:"c"`
	OperationId int64  `json:"a"`
	Name        string `json:"n"`
}

// All returns all companies
func (r *CompanyService) All() ([]*CompanyOperation, error) {
	var companyResponse *CompanyResponse
	_, err := r.client.Request("GET", defaultCompanyPath, nil, &companyResponse)

	return companyResponse.Operations, err
}
