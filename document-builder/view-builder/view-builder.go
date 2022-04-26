package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type ViewBuilder interface {
	BuildView() View
	AppendView(View) ViewBuilder
	CreateCompany() CompanyBuilder
	CreateDateRange() DateRangeBuilder
	// CreateBalance()
}

type viewBuilder struct {
	companyBuilder   *companyBuilder
	dateRangeBuilder *dateRangeBuilder
	Company          View_Company
	DateRange        View_DateRange
}

func (vb *viewBuilder) BuildView() View {
	return View{
		Company:   vb.Company,
		DateRange: vb.DateRange,
	}
}

func (vb *viewBuilder) AppendView(view View) ViewBuilder {
	return vb
}

func (vb *viewBuilder) CreateCompany() CompanyBuilder {
	return vb.companyBuilder
}

func (vb *viewBuilder) CreateDateRange() DateRangeBuilder {
	return vb.dateRangeBuilder
}

func NewViewsBuilder() *viewBuilder {
	viewBuilderInstance := &viewBuilder{}
	companyBuilderInstance := &companyBuilder{}
	dateRangeBuilderInstance := &dateRangeBuilder{}

	viewBuilderInstance.dateRangeBuilder = dateRangeBuilderInstance
	viewBuilderInstance.companyBuilder = companyBuilderInstance

	companyBuilderInstance.injectViewInstance(viewBuilderInstance)
	dateRangeBuilderInstance.injectViewInstance(viewBuilderInstance)
	return viewBuilderInstance
}
