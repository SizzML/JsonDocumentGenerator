package document

type DocumentDto struct {
	Document Document `json:"document"`
}

type Document struct {
	Views    []View   `json:"views"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Rows      int                `json:"rows"`
	RangeDate Metadata_RangeDate `json:"range_date"`
	User      string             `json:"user"`
}

type Metadata_RangeDate struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
}

type View struct {
	Company   View_Company   `json:"company"`
	DateRange View_DateRange `json:"date_range"`
	Balance   View_Balance   `json:"balance"`
	Details   View_Details   `json:"details"`
}
