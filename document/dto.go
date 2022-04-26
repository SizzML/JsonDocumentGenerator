package document

type DocumentDto struct {
	Document Document
}

type Document struct {
	Views    []View
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Rows      string `json:"rows"`
	RangeDate struct {
		Begin string `json:"begin"`
		End   string `json:"end"`
	} `json:"range_date"`
	User string `json:"user"`
}

type View struct {
	Company struct {
		Columns struct {
			BusinessName string `json:"business_name"`
			Address      string `json:"address"`
			Name         string `json:"name"`
		} `json:"columns"`
		Records struct {
			BusinessName []string `json:"business_name"`
			Address      []string `json:"address"`
			Name         []string `json:"name"`
		} `json:"records"`
	} `json:"company"`
	DateRange struct {
		Columns DateRangeType `json:"columns"`
		Records DateRangeType `json:"records"`
	} `json:"date_range"`
	Balance struct {
		Columns BalanceType `json:"columns"`
		Records BalanceType `json:"records"`
	} `json:"balance"`
	Details struct {
		Columns struct {
			AccreditationDate string `json:"accreditation_date"`
			MovementType      string `json:"movement_type"`
			TransactionType   string `json:"transaction_type"`
			TransactionId     string `json:"transaction_id"`
			AccountCommission string `json:"account_commission"`
			BusinessName      string `json:"business_name"`
		} `json:"columns"`
		Records []string `json:"records"`
	} `json:"details"`
}

type BalanceType struct {
	InitialBalance string `json:"initial_balance"`
	Deposits       string `json:"deposits"`
	Debits         string `json:"debits"`
	Commisions     string `json:"commisions"`
	Retencion      string `json:"retencion"`
	AverageBalance string `json:"average_balance"`
	Balance        string `json:"balance"`
}

type DateRangeType struct {
	From           string `json:"from"`
	To             string `json:"to"`
	GenerationDate string `json:"generation_date"`
	IdUser         string `json:"id_user"`
}
