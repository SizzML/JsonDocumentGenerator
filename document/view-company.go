package document

type View_Company struct {
	Columns Company_Columns `json:"columns"`
	Records Company_Records `json:"records"`
}

type Company_Columns struct {
	BusinessName string `json:"business_name"`
	Address      string `json:"address"`
	Name         string `json:"name"`
}

type Company_Records struct {
	BusinessName []string `json:"business_name"`
	Address      []string `json:"address"`
	Name         []string `json:"name"`
}
