package sptrans

const (
	defaultCompanyPath = "Empresa"
)

// CompanieService provide functions to request companies endpoints
type CompanyService service

// Company structure, see documentation on http://www.sptrans.com.br/desenvolvedores/APIOlhoVivo/Documentacao.aspx?1#docApi-empresas
type Company struct {
	Hr string              `json:"hr"`
	E  []*CompanyOperation `json:"e"`
}

// CompanyOperation describes operation
type CompanyOperation struct {
	A int64      `json:"a"`
	E []*Company `json:"e"`
}

//CompanyInfo describes company
type CompanyInfo struct {
	A int64  `json:"a"`
	C int64  `json:"c"`
	N string `json:"n"`
}

// All returns all companies
func (r *CompanyService) All() ([]*CompanyOperation, error) {
	var company *Company
	_, err := r.client.Request("GET", defaultCompanyPath, nil, &company)

	if err != nil {
		return nil, err
	}

	return company.E, err
}
