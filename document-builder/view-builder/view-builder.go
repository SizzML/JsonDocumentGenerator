package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type ViewBuilder interface {
	BuildView() View
	CreateCompany() CompanyBuilder
	CreateDateRange() DateRangeBuilder
	CreateBalance() BalanceBuilder
	CreateDetails() DetailsBuilder
}

type viewBuilder struct {
	companyBuilder   *companyBuilder
	dateRangeBuilder *dateRangeBuilder
	balanceBuilder   *balanceBuilder
	detailsBuilder   *detailsBuilder
	Company          View_Company
	DateRange        View_DateRange
	Balance          View_Balance
	Details          View_Details
}

func (vb *viewBuilder) BuildView() View {
	return View{
		Company:   vb.Company,
		DateRange: vb.DateRange,
		Balance:   vb.Balance,
		Details:   vb.Details,
	}
}

func (vb *viewBuilder) CreateCompany() CompanyBuilder {
	return vb.companyBuilder
}

func (vb *viewBuilder) CreateDateRange() DateRangeBuilder {
	return vb.dateRangeBuilder
}

func (vb *viewBuilder) CreateBalance() BalanceBuilder {
	return vb.balanceBuilder
}

func (vb *viewBuilder) CreateDetails() DetailsBuilder {
	return vb.detailsBuilder
}

func NewViewsBuilder() *viewBuilder {
	viewBuilderInstance := &viewBuilder{}
	companyBuilderInstance := &companyBuilder{}
	dateRangeBuilderInstance := &dateRangeBuilder{}
	balanceBuilderInstance := &balanceBuilder{}
	detailsBuilderInstance := &detailsBuilder{}

	viewBuilderInstance.dateRangeBuilder = dateRangeBuilderInstance
	viewBuilderInstance.companyBuilder = companyBuilderInstance
	viewBuilderInstance.balanceBuilder = balanceBuilderInstance
	viewBuilderInstance.detailsBuilder = detailsBuilderInstance

	companyBuilderInstance.injectViewInstance(viewBuilderInstance)
	dateRangeBuilderInstance.injectViewInstance(viewBuilderInstance)
	balanceBuilderInstance.injectViewInstance(viewBuilderInstance)
	detailsBuilderInstance.injectViewInstance(viewBuilderInstance)
	return viewBuilderInstance
}
