package document

type View_DateRange struct {
	Columns DateRange_Columns `json:"columns"`
	Records DateRange_Records `json:"records"`
}

type DateRange_Columns struct {
	From           string `json:"from"`
	To             string `json:"to"`
	GenerationDate string `json:"generation_date"`
	IdUser         string `json:"id_user"`
}

type DateRange_Records struct {
	From           string `json:"from"`
	To             string `json:"to"`
	GenerationDate string `json:"generation_date"`
	IdUser         string `json:"id_user"`
}

// type DateRangeType struct {
// 	From           string `json:"from"`
// 	To             string `json:"to"`
// 	GenerationDate string `json:"generation_date"`
// 	IdUser         string `json:"id_user"`
// }
