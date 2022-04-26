package document

type View_Details struct {
	Columns Details_Columns `json:"columns"`
	Records Details_Records `json:"records"`
}

type Details_Columns struct {
	AccreditationDate string `json:"accreditation_date"`
	MovementType      string `json:"movement_type"`
	TransactionType   string `json:"transaction_type"`
	TransactionId     string `json:"transaction_id"`
	Currency          string `json:"currency"`
	AccountCommission string `json:"account_commission"`
	BusinessName      string `json:"business_name"`
}
type Details_Records []string
