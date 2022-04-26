package document

type View_Balance struct {
	Columns Balance_Columns `json:"columns"`
	Records Balance_Records `json:"records"`
}
type Balance_Columns struct {
	InitialBalance string `json:"initial_balance"`
	Deposits       string `json:"deposits"`
	Debits         string `json:"debits"`
	Commisions     string `json:"commisions"`
	Retencion      string `json:"retencion"`
	AverageBalance string `json:"average_balance"`
	Balance        string `json:"balance"`
}
type Balance_Records struct {
	InitialBalance string `json:"initial_balance"`
	Deposits       string `json:"deposits"`
	Debits         string `json:"debits"`
	Commisions     string `json:"commisions"`
	Retencion      string `json:"retencion"`
	AverageBalance string `json:"average_balance"`
	Balance        string `json:"balance"`
}

// type BalanceType struct {
// 	InitialBalance string `json:"initial_balance"`
// 	Deposits       string `json:"deposits"`
// 	Debits         string `json:"debits"`
// 	Commisions     string `json:"commisions"`
// 	Retencion      string `json:"retencion"`
// 	AverageBalance string `json:"average_balance"`
// 	Balance        string `json:"balance"`
// }
